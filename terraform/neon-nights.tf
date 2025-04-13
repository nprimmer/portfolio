module "neon-nights-frontend" {
  source               = "./application"
  image_name           = "neon-nights-frontend"
  port                 = "3000"
  fqdn                 = "miami-first-federal-of-miami.toolchest.app"
  cloudflare_api_token = var.cloudflare_api_token
  cloudflare_zone_id   = var.cloudflare_zone_id
  remove_record        = true
  replica_count        = 0
}

module "neon-nights-backend" {
  source               = "./application"
  image_name           = "neon-nights-backend"
  port                 = "8080"
  fqdn                 = "first-miami-backend.toolchest.app"
  cloudflare_api_token = var.cloudflare_api_token
  cloudflare_zone_id   = var.cloudflare_zone_id
  liveness_endpoint    = "/health"
  remove_record        = true
  replica_count        = 0
}
