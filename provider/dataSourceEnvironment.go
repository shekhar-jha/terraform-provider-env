package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func dataSourceEnvironmentSchema() *schema.Resource {
	log.Printf("Entering dataSourceEnvironmentSchema()")
	defer log.Printf("Exiting dataSourceEnvironmentSchema()")
	return &schema.Resource{
		ReadContext: dataSourceEnvironmentRead,
		Schema: map[string]*schema.Schema{
			"call_context": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
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
	log.Println("Entering dataSourceEnvironmentRead()")
	defer log.Println("Exiting dataSourceEnvironmentRead()")
	var diags diag.Diagnostics
	var envValues []string = os.Environ()
	log.Printf("Read environment %d", len(envValues))
	var envVars = make([]map[string]interface{}, len(envValues), len(envValues))
	sort.Strings(envValues)
	for counter, value := range envValues {
		pair := strings.SplitN(value, "=", 2)
		var envVar map[string]interface{} = make(map[string]interface{})
		log.Printf("Adding %d item with env name %s", counter, pair[0])
		envVar["id"] = counter
		envVar["env_name"] = pair[0]
		envVar["env_value"] = pair[1]
		envVars[counter] = envVar
	}
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10)) //very important since without it the data is not passed back.
	if err := d.Set("env_vars", envVars); err != nil {
		log.Printf("Setting value 'env_vars' failed with error %v", err)
		return diag.FromErr(err)
	}
	contextValue := fmt.Sprintf("%v", ctx)
	if err := d.Set("call_context", contextValue); err != nil {
		log.Printf("Setting value 'call_context' failed with error %v", err)
		return diag.FromErr(err)
	}
	return diags
}
