package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			// terraform date source name: data source schema
			"env_os": dataSourceEnvironmentSchema(),
		},
	}
}
