# Day 2

## Lab - Installing AWX 

#### Let's install minikube
```

curl -LO https://github.com/kubernetes/minikube/releases/latest/download/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube && rm minikube-linux-amd64

minikube config cpus 4
minikube config memory 12288
minikube start

# Download kubectl
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
chmod +x ./kubectl
sudo mv kubectl /usr/local/bin
```

#### Let's install 
```
# Clone the awx operator to install Ansible Tower within minikube
git clone https://github.com/ansible/awx-operator.git
cd awx-operator
git checkout tags/2.7.2
export VERSION=2.7.2

# Install make
sudo apt install make -y
make deploy
```
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/d4786b14-9d1c-4d39-bf2b-a0cfa7c50abc" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/24dbf636-80bd-4e39-9936-7224de6e91c8" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/a8268e57-5688-461d-8d8c-9fd5c429b95f" />

#### Check if the AWX required pods are running
```
oc get pods -n awx -w
```

#### Let's create a nodeport service for AWX
Create a file awx.yml with the below code
<pre>
---
apiVersion: awx.ansible.com/v1beta1
kind: AWX
metadata:
  name: awx-tower
spec:
  service_type: nodeport  
</pre>

Let's create the ansible tower instance
``` 
kubectl apply -f awx.ym
```

