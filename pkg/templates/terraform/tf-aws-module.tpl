# This is the terraform template
# Customize this file as needed
# 

{{define "provider.tf"}}
{{if .IsTerraformV0_13OrLater}}
# Define the required providers here
# This section is for Terraform v0.13 or later
terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

provider "aws" {
  region = " "
}
{{else}}
# This section is for Terraform v0.12 or earlier
provider "aws" {
  version = "~> 5.0"
  region  = " "
}
{{end}}
{{end}}

{{define "main.tf"}}
# Main entry point for resources
# Add resource definitions here
{{end}}

{{define "variables.tf"}}
# Declare input variables here
{{end}}

{{define "outputs.tf"}}
# Declare output values here
{{end}}

{{define "Readme.md"}}
# {{ .ProjectName }}
/# Author: {{ .Author }}
/# Created At: {{ .CreatedAt }}
{{end}}

# Module Structure

{{define ".pre-commit-config.yaml"}}
repos:
  - repo: https://github.com/antonbabenko/pre-commit-terraform
    rev: v1.97.0
    hooks:
      - id: terraform_fmt
      - id: terraform_wrapper_module_for_each
      - id: terraform_docs
        args:
          - "--args=--lockfile=false"
      - id: terraform_validate
  - repo: https://github.com/pre-commit/pre-commit-hooks
    rev: v5.0.0
    hooks:
      - id: check-merge-conflict
      - id: end-of-file-fixer
      - id: trailing-whitespace
{{end}}
