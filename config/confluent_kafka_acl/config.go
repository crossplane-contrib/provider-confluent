package confluent_kafka_acl

import (
	"fmt"

	xpref "github.com/crossplane/crossplane-runtime/pkg/reference"
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/crossplane/upjet/pkg/config"
	"github.com/crossplane/upjet/pkg/resource"
)

// Constants for custom Extractor function
var (
	selfPackagePath     = "github.com/stakater/provider-confluent/config/confluent_kafka_acl"
	extractResourceIDFn = selfPackagePath + ".ExtractResourceID()"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("confluent_kafka_acl", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = "confluent"
		r.UseAsync = true
		r.Kind = "KafkaACL"

		// Allows us to reference managedResource ID via spec.forProvider.managedResource.id
		r.References["kafka_cluster.id"] = config.Reference{
			Type: "Cluster",
		}

		// Allows us to reference managedResource ID via spec.forProvider.principal.idSelector
		r.References["principal"] = config.Reference{
			Type:      "ServiceAccount",
			Extractor: extractResourceIDFn,
		}
	})
}

func ExtractResourceID() xpref.ExtractValueFn {
	return func(mr xpresource.Managed) string {
		tr, ok := mr.(resource.Terraformed)
		if !ok {
			return ""
		}

		return fmt.Sprintf("User:%v", tr.GetID()) // append 'User:' infront of the service account ID when resolving 'principal' field.
	}
}
