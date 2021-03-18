package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func Provider() *schema.Provider {
	log.Printf("Entering Provider()")
	defer log.Printf("Exiting Provider()")
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			// terraform date source name: data source schema
			"env_os": dataSourceEnvironmentSchema(),
		},
	}
}
