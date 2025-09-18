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
