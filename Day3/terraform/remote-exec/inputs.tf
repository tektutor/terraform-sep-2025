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
}
