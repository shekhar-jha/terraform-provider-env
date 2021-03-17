---
page_title: "env_os Data Source - terraform-provider-env"

subcategory: ""
description: |-
  The env_os data source allows you to retrieve all environment variables of the terraform process running the provider plugin.
---

# Data Source `env_os`

The env_os data source allows you to retrieve all environment variables of the terraform process running the provider plugin.

## Example Usage

```terraform
data "env_os" "all_envs" {
}
```

## Attributes Reference

The following attributes are exported.

- **env_vars** (List, Compute) - List of environment variables. See [env_var](#env_var) below for details

### env_var

- **id** (Int, Compute) - Identifier
- **env_name** (String, Compute) - Name of environment variable
- **env_value** (String, Compute) - Value of environment variable
