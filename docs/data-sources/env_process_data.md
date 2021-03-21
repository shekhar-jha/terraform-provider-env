---
page_title: "env_process_data Data Source - terraform-provider-env"

subcategory: ""
description: |- The env_process_data data source allows you to retrieve details of process running the provider plugin
along with some host information.
---

# Data Source `env_process_data`

The `env_process_data` data source allows you to retrieve details of process running the provider plugin along with some
host information.

## Example Usage

```terraform
data "env_process_data" "process_details" {}
```

## Attributes Reference

The following attributes are exported.

- **call_context** (String, Compute) - Context passed to function call.
- **executable_file** (Map, Compute) - Details about the provider plugin executable file.
  See [file_details](#file_details) below for additional details.
- **host_name** (String, Compute) - Host name of machine on which provider plugin is running.
- **temp_dir** (Map, Compute) - Details about the Temporary directory associated with the process.
  See [file_details](#file_details) below for additiona details.
- **work_dir** (Map, Compute) - Details about the present working directory associated with the process.
  See [file_details](#file_details) below for additiona details.
- **user_cache_dir** (String, Compute) - Path of the cache directory associated with user.
- **user_config_dir** (String, Compute) - Path of the configuration directory associated with user.
- **user_home_dir** (String, Compute) - Path of the home directory associated with user.
- **process_id** (Integer, Compute) - Process ID of the provider plugin process running.
- **process_parent_id** (Integer, Compute) - ID of parent process of provider plugin process running.
- **process_group_id_effective** (Integer, Compute) - Effective group ID of provider plugin process running.
- **process_user_id_effective** (Integer, Compute) - Effective user ID of provider plugin process running.
- **process_user_id** (Integer, Compute) - User ID of provider plugin process running.
- **process_group_id** (Integer, Compute) - Primary group ID of provider plugin process running.
- **process_secondary_group_ids** ([]Integer, Compute) - Additional groups associated with provider plugin process
  running.

### file_details

- **path** (String, Compute) - Location of the file including file name
- **name** (String, Compute) - Name of file
- **isDirectory** (Bool, Compute) - Whether file is directory
- **size** (String, Compute) - Size of file/directory
- **mode** (String, Compute) - File mode including sticky bit. e.g. `-rwxr-x---`
- **last_modified** (String, Compute) - Last modified time in `RFC 3339` format e.g. `2021-03-21T12:54:59-05:00`
