/*
Copyright 2023 The Crossplane Authors.
*/

package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	ujconfig "github.com/crossplane/upjet/pkg/config"

	confluentapikey "github.com/stakater/provider-confluent/config/confluent_api_key"
	confluentenvironment "github.com/stakater/provider-confluent/config/confluent_environment"
	confluentidentityprovider "github.com/stakater/provider-confluent/config/confluent_identity_provider"
	confluentkafkaacl "github.com/stakater/provider-confluent/config/confluent_kafka_acl"
	confluentkafkacluster "github.com/stakater/provider-confluent/config/confluent_kafka_cluster"
	confluentkafkaclusterconfig "github.com/stakater/provider-confluent/config/confluent_kafka_cluster_config"
	confluentrolebinding "github.com/stakater/provider-confluent/config/confluent_role_binding"
	confluentserviceaccount "github.com/stakater/provider-confluent/config/confluent_service_account"
)

const (
	resourcePrefix = "confluent"
	modulePath     = "github.com/stakater/provider-confluent"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("crossplane.io"),
		ujconfig.WithIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		confluentenvironment.Configure,
		confluentkafkaclusterconfig.Configure,
		confluentkafkacluster.Configure,
		confluentserviceaccount.Configure,
		confluentapikey.Configure,
		confluentkafkaacl.Configure,
		confluentrolebinding.Configure,
		confluentidentityprovider.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
