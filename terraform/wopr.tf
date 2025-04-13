module "wopr" {
  source               = "./application"
  image_name           = "wopr"
  port                 = "8081"
  liveness_endpoint    = "/healthcheck"
  fqdn                 = "wopr.toolchest.app"
  cloudflare_api_token = var.cloudflare_api_token
  cloudflare_zone_id   = var.cloudflare_zone_id
  is_http              = false
  remove_record        = true
  replica_count        = 0
}

