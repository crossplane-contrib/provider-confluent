package confluent_connector

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("confluent_connector", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = "confluent"
		r.UseAsync = true
		r.Kind = "Connector"

		// Allows us to reference Environment ID via spec.forProvider.environment.idRef
		r.References["environment.id"] = config.Reference{
			Type: "Environment",
		}

		// Allows us to reference Cluster ID via spec.forProvider.kafkaCluster.idRef
		r.References["kafka_cluster.id"] = config.Reference{
			Type: "Cluster",
		}
	})
}
