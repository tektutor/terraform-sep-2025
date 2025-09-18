terraform {
  required_providers {
    docker = {
       source = "tektutor/docker"
    }
  }
}

resource "docker_image" "nginx" {
  image_name = "bitnami/nginx:latest"
}

resource "docker_container" "container1" {
  container_name = "c1"
  host_name      = "container_1"
  image_name     = "tektutor/ubuntu-ansible-node:latest"
}

resource "docker_container" "container2" {
  container_name = "c2"
  host_name      = "container_2"
  image_name     = "tektutor/ubuntu-ansible-node:latest"
}
