# Day 3

## Info - Terraform Overview
<pre>
- is one of the Infrastructure as a code tool (Iac)
- it is cloud newtral, hence this Iac tool can be used in any public cloud environment like AWS, Azure, GCE, etc.,
- this also called used to provision infrastructure on your on-prem data-centres
- it helps you provision containers, manage images, provision virtual machines locally or on public cloud, etc.,
- it can be used provision storage cluster, etc.,
- it can be used to provision eks, aks, ROSA, ARO on public cloud
- unlike the AWS cloudformation it works on any environment and any cloud
- it comes in 2 flavours
  1. Terraform core ( command-line only - opensource and free )
  2. Terraform Enterprise ( Web console and it is a paid tool )
</pre>

## Info - Terraform High Level Architecture
![Terraform](terraform-architecture-diagram.png)

## Lab - Checking the Terraform version
```
terraform --version
```
<img width="1760" height="422" alt="image" src="https://github.com/user-attachments/assets/7c9dc5b9-074c-4803-93f8-d2a07d00b2e6" />

## Info - Terraform Providers
<pre>
- Terraform depends on Providers to provision resources
- For example
  - In order to provision an ec2 instance in AWS, Terraform depends on a provider called AWS ( registry.terraform.io )
  - IN order to provision an azure VM in Azure portal, Terraform depends on a provider called Azure
  - as long as there is a provider, Terraform can provision resources on that environment
  - In case, to provision a particular type of resource within your organization and there is no read-made provider, you can
    develop your own provider in Golang using Terraform Provider SDK
  - Providers supports two types of objects/resources
    1. Resources
       - If you wish to Provision ec2 instances using Terraform, then you will define a resource block expressing your expected state
       - Terrafrom can Create, Replace, Update and Delete the resources managed by Terraform
    2. DataSources ( already existing resources - these objects will be treated by Terraform as a read-only resource )
       - these resources are not managed by Terraform
       - they are managed outside Terraform
       - Terraform can refer and use it the HCL (Hashicorp Configuration Language - Terraform's proprietary language )
       - IN case to provision certain resource you declarative terraform script(manifest) file depends on already existing resource
         then, we call them as DataSources or Data block
</pre>

## Info - Terraform Resources
<pre>
- Each Terraform Provider supports one to many Resources and one to many Datasources
- For instance, the docker provider supports the following resources
  - docker_image
  - docker_container
</pre>

## Lab - Using existing Docker Image to provision containers locally

First of all, you need a create folder
```
cd ~
mkdir -p terraform-projects/ex1
cd terraform-projects/ex1
touch main.tf
```

Create a file named main.tf
```
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
```
Let's ensure the required terraform providers are downloaded by Terraform before proceeding further
```
cd ~/terraform-projects/ex1
terraform init
```

<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/77549396-feb7-4254-ba15-ce0b08eec782" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/724608e2-1b27-4b97-a151-bb9c185e0824" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/1e273a55-a3dc-4c4d-b843-4a7e113a6902" />


