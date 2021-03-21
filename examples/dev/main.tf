terraform {
  required_providers {
    env = {
      version = "~> 0.1.0"
      source = "github.com/shekhar-jha/env"
    }
  }
}

data "env_os" "all_envs" {
}

data "env_process_data" "process_details" {}

output "all_env_vars" {
  value = data.env_os.all_envs.env_vars
  description = "All environment variables"
}
output "process_info" {
  value = data.env_process_data.process_details
}