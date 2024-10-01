/*
Copyright 2023 The Crossplane Authors.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane-contrib/provider-confluent/apis/v1beta1"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/terraform"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// ProviderConfig secret keys
	cloudAPIKey       = "cloud_api_key"
	cloudAPISecret    = "cloud_api_secret"
	kafkaAPIKey       = "kafka_api_key"
	kafkaAPISecret    = "kafka_api_secret"
	kafkaRESTEndpoint = "kafka_rest_endpoint"

	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal confluent credentials as JSON"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
// NOTE(hasheddan): this function is slightly over our cyclomatic complexity
// goal. Consider refactoring before adding new branches.
func TerraformSetupBuilder(version, providerSource, providerVersion string, scheduler terraform.ProviderScheduler) terraform.SetupFn { //nolint:gocyclo
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
			Scheduler: scheduler,
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}
		creds := map[string]string{}
		if err := json.Unmarshal(data, &creds); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		// Set credentials in Terraform provider configuration.
		ps.Configuration = map[string]any{}
		if cloudAPIKey, ok := creds[cloudAPIKey]; ok {
			ps.Configuration[cloudAPIKey] = creds[cloudAPIKey]
		}
		if cloudAPISecret, ok := creds[cloudAPISecret]; ok {
			ps.Configuration[cloudAPISecret] = creds[cloudAPISecret]
		}
		if kafkaAPIKey, ok := creds[kafkaAPIKey]; ok {
			ps.Configuration[kafkaAPIKey] = creds[kafkaAPIKey]
		}
		if kafkaAPISecret, ok := creds[kafkaAPISecret]; ok {
			ps.Configuration[kafkaAPISecret] = creds[kafkaAPISecret]
		}
		if kafkaRESTEndpoint, ok := creds[kafkaRESTEndpoint]; ok {
			ps.Configuration[kafkaRESTEndpoint] = creds[kafkaRESTEndpoint]
		}
		return ps, nil
	}
}
