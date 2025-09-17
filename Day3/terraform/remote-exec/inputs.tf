variable "ubuntu_container_name" {
   description = "Name of the container"
   type = string
   default = "ubuntu_container"
}

variable "rocky_container_name" {
   description = "Name of the container"
   type = string
   default = "rocky_container"
}

variable "container_count" {
  description = "Number of containers"
  type = number

  validation {
     condition = var.container_count >= 5 && var.container_count <= 10
     error_message = "The container count must be between 5 and 10 inclussive."
  }
}
