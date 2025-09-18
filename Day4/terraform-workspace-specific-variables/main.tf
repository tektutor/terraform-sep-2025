terraform {
  required_providers {
     local = {
       source = "hashicorp/local"
     }
  }
}

provider "local" {}

locals {
   env_name = terraform.workspace
   
   settings = {
      dev = {
         message = "Welcome to Dev workspace !"
      }
      staging = {
         message = "Welcome to Staging workspace !"
      }
      prod = {
         message = "Welcome to Production workspace !"
      }
   }
}

resource "local_file" "myfile" {
   filename = "${path.module}/message-${terraform.workspace}.txt"
   content  = "${terraform.workspace} - ${local.settings[local.env_name].message} - ${var.common_variable}"
}
