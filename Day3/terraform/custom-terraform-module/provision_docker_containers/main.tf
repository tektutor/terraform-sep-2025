data "docker_image" "tektutor_ansible_ubuntu_image" {
  name = var.ubuntu_docker_image 
}

data "docker_image" "tektutor_ansible_rocky_image" {
  name = var.rocky_docker_image 
}

resource "docker_container" "myubuntu_container" {
    count = var.container_count

    image = data.docker_image.tektutor_ansible_ubuntu_image.name
    name  = "${var.ubuntu_container_name}-${count.index+1}"
    hostname  = "${var.ubuntu_container_name}-${count.index+1}"

    ports {
	internal = "22"
	external = "2${format("%03d", count.index+1)}"
    }

    ports {
	internal = "80"
	external = "8${format("%03d", count.index+1)}"
    }

    depends_on = [
       data.docker_image.tektutor_ansible_ubuntu_image
    ]
}

resource "docker_container" "myrocky_container" {
    count = var.container_count

    image = data.docker_image.tektutor_ansible_rocky_image.name
    name  = "${var.rocky_container_name}-${count.index+1}"
    hostname  = "${var.rocky_container_name}-${count.index+1}"

    ports {
	internal = "22"
	external = "3${format("%03d", count.index+1)}"
    }

    ports {
	internal = "80"
	external = "9${format("%03d", count.index+1)}"
    }

    depends_on = [
       data.docker_image.tektutor_ansible_rocky_image
    ]
}

resource "null_resource" "invoke_ansible_playbook" {
   # This will make sure the playbook is executed everytime
   triggers = {
      always_run = timestamp()
   }

   provisioner "local-exec" {
        environment = {
           ANSIBLE_CONFIG ="${path.module}/ansible.cfg"
	}
	command = "ansible-playbook ${path.module}/install-nginx-playbook.yml"
   }

   depends_on = [
       docker_container.myubuntu_container,
       docker_container.myrocky_container
   ]
}
