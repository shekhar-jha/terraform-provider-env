package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func dataSourceEnvironmentSchema() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceEnvironmentRead,
		Schema: map[string]*schema.Schema{
			"env_vars": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"env_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"env_value": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceEnvironmentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var envValues []string = os.Environ()
	var envVars = make([]map[string]interface{}, len(envValues), len(envValues))
	sort.Strings(envValues)
	for counter, value := range envValues {
		pair := strings.SplitN(value, "=", 2)
		var envVar map[string]interface{} = make(map[string]interface{})
		envVar["id"] = counter
		envVar["env_name"] = pair[0]
		envVar["env_value"] = pair[1]
		envVars[counter] = envVar
	}
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10)) //very important since without it the data is not passed back.
	if err := d.Set("env_vars", envVars); err != nil {
		return diag.FromErr(err)
	}
	return diags
}
