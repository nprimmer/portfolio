terraform {
  backend "gcs" {
    bucket = "neilprimmer-terraform-state"
    prefix = "gc24/state"
  }
}
