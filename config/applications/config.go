package applications

import "github.com/crossplane/upjet/pkg/config"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("authentik_application", func(r *config.Resource) {
		r.ShortGroup = "applications"
		r.Kind = "Application"
		r.References["protocol_provider"] = config.Reference{
			TerraformName: "authentik_provider_oauth2",
		}
	})
	p.AddResourceConfigurator("authentik_application_entitlement", func(r *config.Resource) {
		r.ShortGroup = "applications"
		r.Kind = "ApplicationEntitlement"
		r.References["application"] = config.Reference{
			TerraformName: "authentik_application",
		}
	})
	p.AddResourceConfigurator("authentik_outpost", func(r *config.Resource) {
		r.ShortGroup = "applications"
		r.Kind = "Outpost"
		r.References["protocol_providers"] = config.Reference{
			TerraformName: "authentik_provider_proxy",
		}
		r.References["service_connection"] = config.Reference{
			TerraformName: "authentik_service_connection_kubernetes",
		}
	})
	p.AddResourceConfigurator("authentik_provider_google_workspace", func(r *config.Resource) {
		r.ShortGroup = "applications"
		r.Kind = "ProviderGoogleWorkspace"
	})
	p.AddResourceConfigurator("authentik_provider_ldap", func(r *config.Resource) {
		r.ShortGroup = "applications"
		r.Kind = "ProviderLdap"
		r.References["bind_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["unbind_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
	})
	p.AddResourceConfigurator("authentik_provider_microsoft_entra", func(r *config.Resource) {
		r.ShortGroup = "applications"
		r.Kind = "ProviderMicrosoftEntra"
	})
	p.AddResourceConfigurator("authentik_provider_oauth2", func(r *config.Resource) {
		r.ShortGroup = "applications"
		r.Kind = "ProviderOauth2"
		r.References["authorization_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["authentication_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["invalidation_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
	})
	p.AddResourceConfigurator("authentik_provider_proxy", func(r *config.Resource) {
		r.ShortGroup = "applications"
		r.Kind = "ProviderProxy"
		r.References["authorization_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["authentication_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["invalidation_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
	})
	p.AddResourceConfigurator("authentik_provider_rac", func(r *config.Resource) {
		r.ShortGroup = "applications"
		r.Kind = "ProviderRac"
		r.References["authorization_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["authentication_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
	})
	p.AddResourceConfigurator("authentik_provider_radius", func(r *config.Resource) {
		r.ShortGroup = "applications"
		r.Kind = "ProviderRadius"
		r.References["authorization_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["invalidation_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
	})
	p.AddResourceConfigurator("authentik_provider_saml", func(r *config.Resource) {
		r.ShortGroup = "applications"
		r.Kind = "ProviderSaml"
		r.References["authorization_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["invalidation_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
		r.References["authentication_flow"] = config.Reference{
			TerraformName: "authentik_flow",
		}
	})
	p.AddResourceConfigurator("authentik_provider_scim", func(r *config.Resource) {
		r.ShortGroup = "applications"
		r.Kind = "ProviderScim"
	})
	p.AddResourceConfigurator("authentik_provider_ssf", func(r *config.Resource) {
		r.ShortGroup = "applications"
		r.Kind = "ProviderSsf"
	})
	p.AddResourceConfigurator("authentik_rac_endpoint", func(r *config.Resource) {
		r.ShortGroup = "applications"
		r.Kind = "RacEndpoint"
	})
	p.AddResourceConfigurator("authentik_service_connection_docker", func(r *config.Resource) {
		r.ShortGroup = "applications"
		r.Kind = "ServiceConnectionDocker"
	})
	p.AddResourceConfigurator("authentik_service_connection_kubernetes", func(r *config.Resource) {
		r.ShortGroup = "applications"
		r.Kind = "ServiceConnectionKubernetes"
	})
}
