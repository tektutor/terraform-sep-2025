output "ubuntu_container_names" {
   value = docker_container.myubuntu_container[*].name
}

output "ubuntu_container_ids" {
   value = docker_container.myubuntu_container[*].id
}

output "rocky_container_names" {
   value = docker_container.myrocky_container[*].name
}

output "rocky_container_ids" {
   value = docker_container.myrocky_container[*].id
}

output "ubuntu_container_ips" {
   value = docker_container.myubuntu_container[*].network_data[0].ip_address

}

output "rocky_container_ips" {
   value = docker_container.myrocky_container[*].network_data[0].ip_address
}

