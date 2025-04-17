resource "google_compute_disk" "bbs_root_disk" {
  name  = "bbs-root-disk"
  type  = "pd-standard"
  zone  = "us-central1-a"
  size  = 20
  image = "ubuntu-2004-lts"
}

resource "google_compute_instance" "bbs_instance" {
  name                      = "bbs-instance"
  machine_type              = "e2-medium"
  zone                      = "us-central1-a"
  allow_stopping_for_update = true
  boot_disk {
    source = google_compute_disk.bbs_root_disk.id
  }
  tags = ["bbs"]

  network_interface {
    network = "default"
    access_config {}
  }

  service_account {
    scopes = ["userinfo-email", "compute-ro", "storage-ro"]
  }
}

resource "cloudflare_record" "bbs_record" {
  zone_id = var.cloudflare_zone_id
  type    = "A"
  value   = google_compute_instance.bbs_instance.network_interface.0.access_config.0.nat_ip
  ttl     = 300
  name    = "bbs"
}

resource "google_compute_firewall" "allow_23_inbound" {
  name    = "allow-23-inbound"
  network = "default"

  allow {
    protocol = "tcp"
    ports    = ["23"]
  }

  source_ranges = ["0.0.0.0/0"]
  target_tags   = ["bbs"]
}