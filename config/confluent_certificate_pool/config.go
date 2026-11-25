package confluent_certificate_pool

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("confluent_certificate_pool", func(r *config.Resource) {
		// We need to override the default group that upjet generated for
		r.ShortGroup = "confluent"
		r.UseAsync = true
		r.Kind = "CertificatePool"

		// Allows us to reference the Certificate Authority ID via spec.forProvider.certificateAuthority[].idRef
		r.References["certificate_authority.id"] = config.Reference{
			Type: "CertificateAuthority",
		}
	})
}
