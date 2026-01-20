package confluent_flink_statement

import (
	"github.com/crossplane/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("confluent_flink_statement", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = "flink"
		r.UseAsync = true
		r.Kind = "Statement"

		// Allow references to Flink compute pools
		r.References["compute_pool.id"] = config.Reference{
			Type: "ComputePool",
		}
	})
}
