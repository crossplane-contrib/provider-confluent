package confluent_flink_artifact

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("confluent_flink_artifact", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = "flink"
		r.UseAsync = true
		r.Kind = "Artifact"
	})
}
