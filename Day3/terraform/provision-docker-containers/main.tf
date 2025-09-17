terraform {
  required_providers {
    docker = {
      source = "kreuzwerker/docker"
      version = "3.6.2"
    }
  }
}

provider "docker" {
  # Configuration options
}

# Terraform will consider the docker image as a read-only resource as we are using data block
data "docker_image" "tektutor_ansible_ubuntu_image" {
   name = "tektutor/ubuntu-ansible-node:latest"
}

# Terraform will consider the docker image as a read-only resource as we are using data block
data "docker_image" "tektutor_ansible_rocky_image" {
   name = "tektutor/rocky-ansible-node:latest"
}

# Terraform manages this resource, hence Terraform can Create, Replace, Update, Delete this resource (CRUD operations)
resource "docker_container" "my_ubuntu_container1" {
   image = data.docker_image.tektutor_ansible_ubuntu_image.name
   name  = "ubuntu_container_1"
}

# Terraform manages this resource, hence Terraform can Create, Replace, Update, Delete this resource (CRUD operations)
resource "docker_container" "my_rocky_container1" {
   image = data.docker_image.tektutor_ansible_rocky_image.name
   name  = "rocky_container_1"
}
