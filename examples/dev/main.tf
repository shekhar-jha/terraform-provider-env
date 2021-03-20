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

output "all_env_vars" {
  value = data.env_os.all_envs.env_vars
  description = "All environment variables"
}