resource "kubernetes_service" "default" {
  metadata {
    name      = var.image_name
    namespace = var.namespace
  }
  spec {
    selector = {
      app = var.image_name
    }
    type = var.is_http ? "NodePort" : "LoadBalancer"
    port {
      port        = var.is_http ? "80" : var.port
      target_port = var.port
      protocol    = "TCP"
    }
  }
}
