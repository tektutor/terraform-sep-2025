terraform {
   required_providers {
      docker = {
         source = "kreuzwerker/docker"
         version = "3.6.2"
      }
      local = {
        source = "hashicorp/local"
	version = "2.5.3"
      }
   }
}

provider "docker" {
}
