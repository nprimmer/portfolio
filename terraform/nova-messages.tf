module "nova-messages" {
  source               = "./application"
  image_name           = "nova-messages"
  port                 = "8080"
  fqdn                 = "nova-messages.toolchest.app"
  cloudflare_api_token = var.cloudflare_api_token
  cloudflare_zone_id   = var.cloudflare_zone_id
  liveness_endpoint    = "/health"
  remove_record        = true
}
