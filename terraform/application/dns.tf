resource "cloudflare_record" "record" {
  count   = var.remove_record ? 0 : 1
  zone_id = var.cloudflare_zone_id
  type    = "A"
  value   = var.is_http ? kubernetes_ingress_v1.default.status[0].load_balancer[0].ingress[0].ip : kubernetes_service.default.status[0].load_balancer[0].ingress[0].ip
  ttl     = 300
  name    = split(".", var.fqdn)[0]
}
