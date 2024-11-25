package confluent_kafka_acl

import (
	"github.com/crossplane-contrib/provider-confluent/config/common"
	"github.com/crossplane/upjet/pkg/config"
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
			Extractor: common.ExtractPrincipalIDFuncPath,
		}
	})
}
