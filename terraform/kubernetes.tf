resource "google_service_account" "challenge_cluster_sa" {
  account_id   = "challenge-cluster"
  display_name = "Challenge Cluster Service Account"
}

resource "google_container_cluster" "challenge_cluster" {
  name                     = "challenge-cluster"
  location                 = "us-central1"
  initial_node_count       = 1
  remove_default_node_pool = true
}

resource "google_container_node_pool" "challenge_node_pool" {
  name       = "challenge-node-pool"
  location   = "us-central1"
  cluster    = google_container_cluster.challenge_cluster.name
  node_count = 2
  node_config {
    preemptible     = false
    machine_type    = "e2-standard-2"
    service_account = google_service_account.challenge_cluster_sa.email
    oauth_scopes = [
      "https://www.googleapis.com/auth/cloud-platform",
      "https://www.googleapis.com/auth/devstorage.read_only"
    ]
  }
}

resource "kubernetes_namespace" "challenges" {
  metadata {
    name = "challenges"
  }
}

resource "kubernetes_storage_class" "standard" {
  metadata {
    name = "gce-standard"
  }

  storage_provisioner = "kubernetes.io/gce-pd"
  parameters = {
    type             = "pd-standard"
    replication-type = "none"
  }
  reclaim_policy      = "Retain"
  volume_binding_mode = "WaitForFirstConsumer"
}