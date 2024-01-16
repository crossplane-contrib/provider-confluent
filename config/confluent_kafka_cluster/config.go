package confluent_kafka_cluster

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("confluent_kafka_cluster", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = "confluent"
		r.UseAsync = true

		// Allows us to reference environment by the metadata.name instead of the externally generated (random) name by Confluent.
		r.References["environment.id"] = config.Reference{
			Type: "github.com/stakater/provider-confluent/apis/confluent/v1alpha1.Environment",
		}

		// This is workaround for error related to terraform state containing spec.network[0].id = "", which gets late-initialized by upjet by default.
		// This overrides the default late initialization to ignore the network field in the terraform state file.
		// see https://github.com/upbound/upjet/blob/main/docs/add-new-resource-long.md#late-initialization-configuration
		r.LateInitializer = config.LateInitializer{
			IgnoredFields: []string{"network"},
		}
	})
}
