variable "namespace" {
  default = "challenges"
  type    = string
}

variable "image_name" {
  type = string
}

variable "port" {
  type = string
}

variable "fqdn" {
  type = string
}

variable "is_http" {
  type    = bool
  default = true
}

variable "region" {
  type    = string
  default = "us-central1"
}

variable "cloudflare_zone_id" {
  type = string
}

variable "cloudflare_api_token" {
  type = string
}

variable "replica_count" {
  type    = number
  default = 1
}
variable "liveness_endpoint" {
  type    = string
  default = "/"
}

variable "use_pv" {
  type    = bool
  default = false
}

variable "pv_size_gb" {
  type    = string
  default = "10"
}

variable "pv_mount_point" {
  description = "Mount point for the PersistentVolume"
  type        = string
  default     = "/mnt/data"
}

variable "image_root" {
  type    = string
  default = "us-central1-docker.pkg.dev/neilprimmer/challenge-series"
}

variable "image_version" {
  type    = string
  default = "latest"
}

variable "image_pull_policy" {
  type    = string
  default = "Always"
}

variable "cpu_limit" {
  type    = string
  default = "0.5"
}

variable "cpu_request" {
  type    = string
  default = "250m"
}

variable "memory_limit" {
  type    = string
  default = "256Mi"
}

variable "memory_request" {
  type    = string
  default = "64Mi"
}

variable "remove_record" {
  type    = bool
  default = false
}