
resource "google_compute_ssl_policy" "challenge_ssl_policy" {
  name            = "challenge-ssl-policy"
  profile         = "MODERN"
  min_tls_version = "TLS_1_2"
}
