variable "ubuntu_docker_image" {
   description = "Name of the ubuntu docker image"
   type = string
   default = "tektutor/ubuntu-ansible-node:latest"
}

variable "rocky_docker_image" {
   description = "Name of the rocky docker image"
   type = string
   default = "tektutor/rocky-ansible-node:latest"
}

variable "ubuntu_container_name" {
   description = "Prefix of container name"
   type = string
   default = "ubuntu"
}

variable "rocky_container_name" {
   description = "Prefix of container name"
   type = string
   default = "rocky"
}

variable "container_count" {
   description = "Number of containers you wish to create"
   type = number

   validation {
      condition = var.container_count > 1 && var.container_count < 10
      error_message = "The container count should be in the range 2 to 9"
   }

}
