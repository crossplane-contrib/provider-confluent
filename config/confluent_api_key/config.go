package confluent_api_key

import (
	"github.com/upbound/upjet/pkg/config"
)

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("confluent_api_key", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = "confluent"
		r.UseAsync = true
		r.Kind = "APIKey"

		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}

			if apiKeyID, ok := attr["id"].(string); ok {
				conn["api_key_id"] = []byte(apiKeyID)
			}
			if apiKeySecret, ok := attr["secret"].(string); ok {
				conn["api_key_secret"] = []byte(apiKeySecret)
			}
			return conn, nil
		}

		// Allows us to reference managedResource ID via spec.forProvider.managedResource.id
		r.References["managed_resource.id"] = config.Reference{
			Type: "Cluster",
		}

		// Allows us to reference managedResource.environment ID via spec.forProvider.managedResource.environment.id
		r.References["managed_resource.environment.id"] = config.Reference{
			Type: "Environment",
		}

		// Allows us to reference ServiceAccount ID via spec.forProvider.owner.id
		r.References["owner.id"] = config.Reference{
			Type: "ServiceAccount",
		}

	})
}
