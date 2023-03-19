# Description: This file contains the Terraform code to enable the required GCP APIs for the project

# List of GCP APIs to enable in this project
locals {
  services = [
    "cloudresourcemanager.googleapis.com",
    "iam.googleapis.com",
    "iamcredentials.googleapis.com",
    "servicecontrol.googleapis.com",
    "servicemanagement.googleapis.com",
    "monitoring.googleapis.com",
  ]
}

# Data source to access GCP project metadata 
data "google_project" "project" {}


# Enable the required GCP APIs
resource "google_project_service" "default" {
  for_each = toset(local.services)

  project = var.project_id
  service = each.value

  disable_on_destroy = false
}