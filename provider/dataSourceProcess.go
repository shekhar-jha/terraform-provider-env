package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"os"
	"strconv"
	"time"
)

func dataSourceProcessSchema() *schema.Resource {
	log.Printf("Entering dataSourceProcessSchema()")
	defer log.Printf("Exiting dataSourceProcessSchema()")
	return &schema.Resource{
		ReadContext: dataSourceProcessRead,
		Schema: map[string]*schema.Schema{
			"call_context": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"executable_file": GetFileDetailSchema(),
			"host_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"temp_dir": GetFileDetailSchema(),
			"work_dir": GetFileDetailSchema(),
			"user_cache_dir": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_config_dir": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_home_dir": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"process_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"process_parent_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"process_group_id_effective": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"process_user_id_effective": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"process_user_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"process_group_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"process_secondary_group_ids": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeInt,
				},
			},
		},
	}
}

func dataSourceProcessRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Println("Entering dataSourceProcessRead()")
	defer log.Println("Exiting dataSourceProcessRead()")
	var diags diag.Diagnostics
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10)) //very important since without it the data is not passed back.
	diags = AddAttributeValueToResourceData(func() interface{} { return fmt.Sprintf("Context %v", ctx) }, d, "call_context", diags)
	diags = AddAttributeValueToResourceDataAndProcessValueWithHandleError(func() (interface{}, error) { return os.Executable() }, func(value interface{}) (interface{}, error) { return GetFileDetail(ctx, value) }, d, "executable_file", "executable file", diags)
	diags = AddAttributeValueToResourceDataAndProcessValueWithHandleError(func() (interface{}, error) { return os.Getwd() }, func(value interface{}) (interface{}, error) { return GetFileDetail(ctx, value) }, d, "work_dir", "working directory", diags)
	diags = AddAttributeValueToResourceDataAndHandleError(func() (interface{}, error) { return os.UserCacheDir() }, d, "user_cache_dir", "user cache directory", diags)
	diags = AddAttributeValueToResourceDataAndHandleError(func() (interface{}, error) { return os.UserConfigDir() }, d, "user_config_dir", "user config directory", diags)
	diags = AddAttributeValueToResourceDataAndHandleError(func() (interface{}, error) { return os.UserHomeDir() }, d, "user_home_dir", "home directory", diags)
	diags = AddAttributeValueToResourceDataAndHandleError(func() (interface{}, error) { return os.Hostname() }, d, "host_name", "hostname", diags)
	diags = AddAttributeValueToResourceDataAndHandleError(func() (interface{}, error) { return GetFileDetail(ctx, os.TempDir()) }, d, "temp_dir", "temporary directory", diags)
	diags = AddAttributeValueToResourceData(func() interface{} { return os.Getpid() }, d, "process_id", diags)
	diags = AddAttributeValueToResourceData(func() interface{} { return os.Getppid() }, d, "process_parent_id", diags)
	diags = AddAttributeValueToResourceData(func() interface{} { return os.Getegid() }, d, "process_group_id_effective", diags)
	diags = AddAttributeValueToResourceData(func() interface{} { return os.Geteuid() }, d, "process_user_id_effective", diags)
	diags = AddAttributeValueToResourceData(func() interface{} { return os.Getuid() }, d, "process_user_id", diags)
	diags = AddAttributeValueToResourceData(func() interface{} { return os.Getgid() }, d, "process_group_id", diags)
	diags = AddAttributeValueToResourceDataAndHandleError(func() (interface{}, error) { return os.Getgroups() }, d, "process_secondary_group_ids", "group ids of process", diags)
	return diags
}

