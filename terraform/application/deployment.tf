resource "kubernetes_persistent_volume_claim" "pvc" {
  count = var.use_pv ? 1 : 0

  metadata {
    name      = "${var.image_name}-pvc"
    namespace = var.namespace
  }

  spec {
    access_modes = ["ReadWriteOnce"]
    resources {
      requests = {
        storage = "${var.pv_size_gb}Gi"
      }
    }
    storage_class_name = "gce-standard"
  }
}

resource "kubernetes_deployment" "deployment" {
  metadata {
    namespace = var.namespace
    name      = var.image_name
    labels = {
      app = var.image_name
    }
  }

  spec {
    replicas = var.replica_count

    selector {
      match_labels = {
        app = var.image_name
      }
    }

    template {
      metadata {
        labels = {
          app = var.image_name
        }
      }

      spec {
        container {
          image_pull_policy = var.image_pull_policy
          image             = "${var.image_root}/${var.image_name}:${var.image_version}"
          name              = var.image_name

          resources {
            limits = {
              cpu    = var.cpu_limit
              memory = var.memory_limit
            }
            requests = {
              cpu    = var.cpu_request
              memory = var.memory_request
            }
          }

          liveness_probe {
            http_get {
              path = var.liveness_endpoint
              port = var.port
            }
            initial_delay_seconds = 5
            period_seconds        = 5
          }

          readiness_probe {
            http_get {
              path = var.liveness_endpoint
              port = var.port
            }
            initial_delay_seconds = 5
            period_seconds        = 5
          }

          dynamic "volume_mount" {
            for_each = var.use_pv ? [1] : []
            content {
              mount_path = var.pv_mount_point
              name       = "data-volume"
            }
          }
        }

        dynamic "volume" {
          for_each = var.use_pv ? [1] : []
          content {
            name = "data-volume"
            persistent_volume_claim {
              claim_name = kubernetes_persistent_volume_claim.pvc[0].metadata[0].name
            }
          }
        }
      }
    }
  }
}