package provider

import (
   "context"
   "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
   "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
  schema.DescriptionKind = schema.StringMarkdown
}

func New(version string) func() *schema.Provider {
    return func() *schema.Provider {
         p := &schema.Provider {
		DataSourcesMap: map[string]*schema.Resource{},
		ResourcesMap: map[string]*schema.Resource{
			"docker_container": resourceDockerContainer(),
			"docker_image": resourceDockerImage(),
		},
	 }

	 p.ConfigureContextFunc = configure(version, p)

	 return p
    }
}

type DockerConfig struct {
  dockerHost string
  username string
  password string
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) ( any, diag.Diagnostics ) {
	return func( context.Context, *schema.ResourceData ) ( any, diag.Diagnostics ) {
		dockerConfig := DockerConfig{}
		return &dockerConfig,nil
	}
}
