terraform {
  required_providers {
    cloudflare = {
      source  = "cloudflare/cloudflare"
      version = "~> 3.0"
    }
    google = {
      source  = "hashicorp/google"
      version = "~> 5.0"
    }
  }
}

provider "google" {
  project = "neilprimmer"
  region  = "us-central1"
}

provider "cloudflare" {
  api_token = var.cloudflare_api_token
}

data "google_client_config" "provider" {}


provider "kubernetes" {
  host  = "https://${google_container_cluster.challenge_cluster.endpoint}"
  token = data.google_client_config.provider.access_token
  cluster_ca_certificate = base64decode(
    google_container_cluster.challenge_cluster.master_auth[0].cluster_ca_certificate,
  )
}

variable "cloudflare_api_token" {}
variable "cloudflare_zone_id" {}
variable "cloudflare_account_id" {}