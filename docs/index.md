---
page_title: "Provider: env"
subcategory: ""
description: |-
  Terraform provider for interacting with machine on which terraform is running.
---

# Env Provider

Provides access to environment details of the server on which terraform is being run.

## Example Usage

```terraform
terraform {
  required_providers {
    env = {
      version = "~> 0.1"
      source = "shekhar-jha/env"
    }
  }
}

data "env_os" "all_envs" {
}
```
