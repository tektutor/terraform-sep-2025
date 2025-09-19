# Day 4

## Lab - Terraform Workspace
<pre>
cd ~/terraform-sep-2025
git pull
cd Day4/terraform-workspace
terraform init
ls -lha
tree .terraform

# Creates a new workspace named dev  
terraform workspace new dev
ls -lha
tree terraform.tfstate.d

# Creates a new workspace named stage  
terraform workspace new stage
ls -lha
tree terraform.tfstate.d

# Creates a new workspace named prod  
terraform workspace new prod
ls -lha
tree terraform.tfstate.d

terraform workspace show
terraform workspace list
terraform workspace select dev
terraform apply --auto-approve
tree terraform.tfstate.d

terraform workspace show
terraform workspace list
terraform workspace select stage
terraform apply --auto-approve
tree terraform.tfstate.d

terraform workspace show
terraform workspace list
terraform workspace select prod
terraform apply --auto-approve
tree terraform.tfstate.d

ls -1
</pre>

<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/76f67b4d-877f-45e1-af44-73aa1363134e" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/b3e281f6-38ae-4e5f-97e7-7fab0e0a5d42" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/3aa96dc4-9a70-4287-8aff-ca661b421958" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/46b1ddb3-d5d0-49bf-97c9-261b2c150208" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/47950f38-aee8-45a9-a50d-70606c6b56b3" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/0a6afdcf-1e89-4b95-a982-979fe2692fa9" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/d1285a34-825f-429d-9451-afd2f67fd1a5" />


## Lab - Using workspace specific variables
<pre>
cd ~/terraform-sep-2025
git pull
cd Day4/terraform-workspace-specific-variables
terraform init
ls -lha
tree .terraform

# Creates a new workspace named dev  
terraform workspace new dev
ls -lha
tree terraform.tfstate.d

# Creates a new workspace named stage  
terraform workspace new stage
ls -lha
tree terraform.tfstate.d

# Creates a new workspace named prod  
terraform workspace new prod
ls -lha
tree terraform.tfstate.d

terraform workspace show
terraform workspace list
terraform workspace select dev
terraform apply --auto-approve
tree terraform.tfstate.d

terraform workspace show
terraform workspace list
terraform workspace select staging
terraform apply --auto-approve
tree terraform.tfstate.d

terraform workspace show
terraform workspace list
terraform workspace select prod
terraform apply --auto-approve
tree terraform.tfstate.d

ls -1  
</pre>
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/c1673928-7e97-48dd-ac17-b96d5b6a73f7" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/6b2a9d01-8509-4a09-b167-e292d81515fe" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/89de6300-3770-4e02-8fc3-55a9cb7a9a77" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/53744f89-941f-47bf-9d03-0696eb9c695e" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/287d0774-c174-4367-9d2f-16f655a91b86" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/94ef9bf9-305c-410e-b9da-6d55ff9e0f45" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/b9edf707-f8d2-4e16-9ab7-87cc78305c97" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/2951e11d-4e30-456a-addd-9240cc6de3a7" />

## Lab - Developing a custom terraform provider plugin
You need to crate folder
```
mkdir -p ~/go/bin
touch ~/.terraformrc
```
Paste the below code in your ~/.terraformrc file
<pre>
provider_installation {
  dev_overrides {
      "registry.terraform.io/tektutor/file" = "/home/rps/go/bin",
  }
  direct {}
}
</pre>

Then you may proceed with the below instruction
```
cd ~/terraform-sep-2025
git pull
cd Day4/custom-terraform-providers/file
tree

go mod init github.com/tektutor/terraform-provider-file
go mod tidy
ls -l
go build
ls -l
go install
ls -l ~/go/bin
```

<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/bc905044-9664-4a51-9392-c0ec30ace199" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/ac7b59a2-5c70-47eb-82d7-5190188364a7" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/a99e8a43-2172-4011-a099-84391e533d9f" />

## Lab - Using our custom terraform file provider in our Terraform project
```
cd ~/terraform-sep-2025
git pull
cd Day4/custom-terraform-providers/test-file-custom-terraform-provider
ls -l
terraform plan
terraform apply --auto-approve
ls -l
cat 
```

<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/a33c321d-13d9-4d1f-baad-39f74014abcb" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/4c6373a8-1ef1-491e-8126-e4fa1c1088f4" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/fb0b3a99-63c1-431c-87e5-a44a5c351016" />

Once you are done, you may dispose the resources created by Terraform
```
terraform destroy --auto-approve
```

<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/057462fd-1d69-4c99-9698-67a0fee9e9ea" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/712fd7fa-15cb-4359-bdc4-ac5d3475b33d" />

## Lab - Implementing a Custom Docker Terraform Provider
```
cd ~/terraform-sep-2025
git pull
cd Day4/custom-terraform-providers
tree

go mod init github.com/tektutor/terraform-provider-docker
go mod tidy
ls -l

go build
ls -l
go install

ls -l ~/go/bin
```

<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/48e92371-cdd4-40aa-9327-5879bcbc4c3f" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/646cdc1a-a021-4b7c-830e-17d738182ed6" />

## Lab - Using our custom Docker Terraform Provider in Terraform manifest script
```
cd ~/terraform-sep-2025
git pull
cd Day4/custom-terraform-providers/test-docker-custom-terraform-provider
cat main.tf

docker images | grep nginx
docker ps
```

<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/828906d6-a50f-4b9c-b4ae-397f99953a76" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/caea654a-d939-4237-8a8a-21dba80823b9" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/286e5eee-7816-4e90-ab0a-cc0654b49858" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/4d782586-4f91-490a-93c6-58a9e47dda03" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/a5c4e8d4-0049-4f28-9093-07fc1dbbf94f" />

## Info - Sentinel Overview
<pre>
- Sentinel is a policy-as-code framework developed by HashiCorp.
- It lets you define and enforce fine-grained policies across infrastructure and application workflows.
- Think of it as a guardrail system:
  - Terraform manages your infrastructure.
  - Sentinel enforces rules to ensure deployments follow your organization’s standards.  
</pre>

#### Motivation to use Policy-as-Code
<pre>
- Traditional governance relies on documents, manual reviews, or checklists. These are slow and error-prone.
- With Sentinel, you write rules in code that are:
  - Version-controlled (like Terraform files)
  - Testable
  - Automatically enforced during provisioning  
</pre>

#### Sentinel Policy Examples
Sentinel policies are written in a custom language (inspired by HCL and JSON).

Restricting Cloud Regions
<pre>
import "tfplan/v2" as tfplan

main = rule {
  all tfplan.resources.aws_instance as instances {
    all instances as r {
      r.applied.provider_config.region is "us-east-1"
    }
  }
}
</pre>

Ensures all AWS instances are launched only in us-east-1.

Enforce Minimum Instance Size
<pre>
import "tfplan/v2" as tfplan

main = rule {
  all tfplan.resources.aws_instance as instances {
    all instances as r {
      r.applied.instance_type in ["t2.medium", "t2.large", "m5.large"]
    }
  }
}  
</pre>

Blocks provisioning of small/unsupported instance types.
