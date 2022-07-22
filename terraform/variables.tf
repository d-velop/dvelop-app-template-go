variable "aws_region" {
  default = "eu-central-1"
}

variable "signature_secret" {
  description = "base64 encoded secret key for validation of tenant headers"
}

variable "build_version" {
  description = "human readable build metadata like git commit hash or timestamp to identify this build of the lambda function"
  default     = "unknown"
}

variable "asset_hash" {
  description = "hash over all assets"
  default     = "unknown"
}

variable "appname" {
  description = "appname without app suffix e.g. pdf, dms, inbound."
}

variable "domainsuffix" {
  description = "dns suffix for the service endpoint"
}

