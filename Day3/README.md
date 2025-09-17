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



Let's verify what terraform will do if we run it
```
cd ~/terraform-projects/ex1
terraform plan
```

<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/9380afeb-8979-4cc6-9863-a97201235f80" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/85b1bda0-5e12-45ac-a7e9-3e1172d1d55d" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/ff213a63-c250-462b-9a32-41ce9503cacd" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/8090ea60-e433-4823-937e-0f07ea135ba1" />

Let's provision the containers using our Terraform manifest script
```
cd ~/terraform-projects/ex1
terraform apply
ls -lha
cat terraform.tfstate
```
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/50157606-f42b-44de-b433-6b0b7ab0e67d" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/25abb5ce-4ca0-467d-9160-c04455574ea5" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/8d61dc26-c107-4e69-a4fc-a5abf780c8ed" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/44560ac6-08fa-41f3-af33-a4a1dd652775" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/9485e8a2-e1ec-48ab-8205-301958b01a3f" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/81d0b87b-995e-4088-813e-d0843f2ab771" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/9302c1b9-a41f-4c5d-ae72-ecd9c9c310d4" />

Cleaning up the resources created by Terraform once you are done with using the resources
```
terraform destroy
docker ps -a
docker images | grep ansible
```

<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/fa23eaf4-4c9f-4b6d-9624-803dbd3b5067" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/ca6a741f-7213-4bec-9a88-0c8c0bf9eed5" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/563f4cc0-0f80-4f15-8bd7-5da6744845c6" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/ab3eb630-be8b-4b24-8452-dcf9cc584e94" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/e4670fad-4d4f-433a-bff1-e69c3339b340" />

## Lab - Terraform Input and Output variables
```
cd ~/terraform-sep-2025
git pull
cd Day3/terraform/provision-containers-with-input-and-output-variables
terraform init
terraform plan
terraform apply --auto-approve
```

<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/1a393abc-7202-40a9-91f5-3f2027dacb7c" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/b0a07e94-1e12-4e2e-9d33-a996824907f7" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/eb411c3d-06bd-4b31-affd-76bfb8058630" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/584e5581-6a7f-4afd-9336-e28ff8d8f293" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/d90af2a7-2c89-4611-b7b2-34fbd65c5914" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/2c460ca7-5811-4807-ac54-a2bdbbbfef32" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/a84bd256-0944-41c8-9760-4927bb0021b1" />

You could change the container names in the inputs.tf file and perform terraform apply.  Terraform will replace the old containers with new containers with the updated container names.

Once you are done, you may dispose the resources created by terraform
```
terraform destroy --auto-approve
```
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/0ca90aa6-80bf-4560-9345-43ef7e87e12b" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/3b2c1b5a-b7d3-4864-ad58-f391f83ca174" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/17dbaa36-c0b6-4054-8494-4f68ce16f038" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/05482fc0-d9d8-4d5a-96be-e974aacabc18" />

## Lab - Invoking ansible playbook from Terraform using local-exec block
```
cd ~/terraform-sep-2025
git pull
cd Day3/terraform/local-exec
ls -l
docker rm -f $(docker ps -aq)
echo "" > ~/.ssh/known_hosts

terraform init
terraform plan
terraform apply --auto-approve
docker ps
curl http://localhost:8001
curl http://localhost:8002
```

<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/0007db8f-711b-4960-b387-c892e5ee8e63" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/c0cfe5f6-6758-4a1a-a047-1efbf6d9af9f" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/6ac03136-efee-416e-9c93-fc05818870e2" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/3bd5500c-a5eb-4eb0-8353-6ebc6a684643" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/dc4f281f-f66d-4b8b-8a57-ece94827cacb" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/621f1cc2-3494-4609-b246-dae0b189fbf0" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/3c2f3e2e-0efc-4949-8636-20160e91bce7" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/4f74f0b5-04de-499a-8c78-aba126dba763" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/5be0f18d-df4f-43e5-85cb-2f64560e15a4" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/095348ad-3562-4b2f-9be5-3264076a9851" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/1cd8e6e6-b64c-4738-89aa-ddfc55c32027" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/8336085f-a9f9-4c79-8314-ea38187b53e3" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/c80318f5-b3ca-4f71-8443-bf4fdcbd1adc" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/045d1261-3bd9-4e2b-a42f-711ccb14e216" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/bdc34756-138a-4735-b63d-465cda2cc3f4" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/4d459cf8-3389-471e-8471-1e6c2914e284" />

Once you are done, you may dispose the resources
```
terraform destroy --auto-approve
```
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/b292afbd-14ab-46d9-ad6b-bc25d4acbf1c" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/f04457f9-6b03-4376-a764-9de486b6e69f" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/1f75f849-93ee-43c0-814d-413f0c589842" />

## Lab - Terraform - Running commands on the remote machine using remote-exec provisioner block
```
cd ~/terraform-sep-2025
git pull
cd Day3/terraform/remote-exec
terraform init
terraform apply --auto-approve
```
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/6901c010-b9b6-488b-b80c-bef39844a1aa" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/ecf5b391-79ff-4890-8f2a-e72b5ec1b58b" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/783c95b7-6d71-423d-9cbc-d0e2092d3416" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/8370ca0f-c50a-4a50-9014-69aa368cb8e2" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/0552d89d-1b92-45db-82bb-11e208be7744" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/b4c637b7-a775-4b27-8d48-42cf931eab65" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/bc6c0f25-1630-4cbe-9513-13fb18706b27" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/38a0879e-6623-4392-bdfe-ba82ebfc42d7" />
