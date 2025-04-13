resource "kubernetes_manifest" "gke_managed_certificate" {
  provider = kubernetes

  manifest = {
    apiVersion = "networking.gke.io/v1"
    kind       = "ManagedCertificate"
    metadata = {
      name      = "gke-ssl-cert-${var.image_name}"
      namespace = var.namespace
    }
    spec = {
      domains = [var.fqdn]
    }
  }
}

resource "kubernetes_ingress_v1" "default" {
  depends_on = [
    kubernetes_manifest.gke_managed_certificate
  ]
  metadata {
    name      = var.image_name
    namespace = var.namespace
    annotations = {
      "kubernetes.io/ingress.class"            = "gce"
      "networking.gke.io/managed-certificates" = "gke-ssl-cert-${var.image_name}"
      "kubernetes.io/ingress.allow-http"       = "true"
    }
  }

  spec {
    default_backend {
      service {
        name = var.image_name
        port {
          number = var.is_http ? "80" : var.port
        }
      }
    }
    rule {
      host = var.fqdn
      http {
        path {
          path      = "/"
          path_type = "ImplementationSpecific"
          backend {
            service {
              name = var.image_name
              port {
                number = var.is_http ? "80" : var.port
              }
            }
          }
        }
      }
    }
  }
}
