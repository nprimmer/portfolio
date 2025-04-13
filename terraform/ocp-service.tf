module "ocp-service" {
  source               = "./application"
  image_name           = "ocp-service"
  port                 = "8080"
  liveness_endpoint    = "/healthcheck"
  fqdn                 = "ocp-service.toolchest.app"
  cloudflare_api_token = var.cloudflare_api_token
  cloudflare_zone_id   = var.cloudflare_zone_id
  is_http              = false
  remove_record        = true
  replica_count        = 0
}

