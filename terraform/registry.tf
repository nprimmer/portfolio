resource "google_artifact_registry_repository" "challenge_repository" {
  provider      = google
  location      = "us-central1"
  repository_id = "challenge-series"
  description   = "Challenge Series Docker Repository"
  format        = "DOCKER"
}

