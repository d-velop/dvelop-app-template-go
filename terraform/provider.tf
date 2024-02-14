terraform {
  required_providers {
    aws = {
      source = "hashicorp/aws"
      version = "5.35.0"
    }
  }
}

provider "aws" {
  #version = "~> 2.7"
  region  = var.aws_region
}
