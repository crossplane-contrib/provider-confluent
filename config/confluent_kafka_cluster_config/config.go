package confluent_kafka_cluster_config

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("confluent_kafka_cluster_config", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = "confluent"
		r.UseAsync = true

		// Allows us to reference the Kafka cluster ID via spec.forProvider.kafkaCluster.idRef
		r.References["kafka_cluster.id"] = config.Reference{
			Type: "Cluster",
		}
	})
}
