data "docker_image" "tektutor_ansible_ubuntu_image" {
  name = "tektutor/ubuntu-ansible-node:latest"
}

data "docker_image" "tektutor_ansible_rocky_image" {
  name = "tektutor/rocky-ansible-node:latest"
}

resource "docker_container" "my_ubuntu_container1" {
    image = data.docker_image.tektutor_ansible_ubuntu_image.name
    name  = var.container_name1

    ports {
	internal = "22"
	external = "2001"
    }

    ports {
	internal = "80"
	external = "8001"
    }
}

resource "docker_container" "my_rocky_container1" {
    image = data.docker_image.tektutor_ansible_rocky_image.name
    name  = var.container_name2

    ports {
	internal = "22"
	external = "2002"
    }

    ports {
	internal = "80"
	external = "8002"
    }
}

resource "local_file" "my_local_file" {
   filename = "./test.txt"
   content  = "this is a test file"

   provisioner "local-exec" {
	command = "ansible-playbook -i ./inventory.py install-nginx-playbook.yml"
   }
}
