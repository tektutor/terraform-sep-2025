# Day 1

## Info - Configuration Management Overview
<pre>
- is one of the DevOps tools
- helps in automating system administration tasks
  - given a machine with OS pre-installed, using configuration management tools we install/uninstall/upgrade software tools in those machine
  - using configuration management tools, we could do user management, manage service, configure the software installed, manage security policy, add/remove a machine for AD domain, managing Network switches,routers, etc.,
- examples
  - Puppet
  - Chef
  - Ansible
</pre>  

## Info - Puppet Overview
<pre>
- is one of the oldest configuration management tool  
- it uses Puppet Proprietary language as the Domain Specific Language(DSL) to write the automation code
- it follows client/servcer architecture
- the machines/servers/workstations that are managed by Puppet - we need to install a Puppet agent
- Puppet nodes are the servers managed by Puppet
- Puppet nodes can be
  - Linux machine
  - a virtual machine, container, ec2 instances, azure vm, or a on-prem server, cisco routers, switches etc.,
- It is the Puppet agent's responsibility to poll the Puppet Server to download the automation scripts from the server to the Puppet nodes
</pre>

## Ansible Overview
<pre>
- implemented in Python by Michael Deehan
- Michael Deehan is a former Red Hat employee
- Michael Deehan incorporated a company called Ansible Inc
- Michael Deehan developed Ansible Core as an open source configuration management tool
- Ansible comes in 3 flavours
  - Ansible Core ( opensource and command-line )
  - AWX ( supports webconsole & user management - developed on top of Ansible core and it is opensource )
  - Ansible Automation Platform ( a.k.a Ansible Tower in the past )
    - supports web console
    - developed on top of opensource AWX
    - it is an enterprise product that requires license for commercial use
    - can be installed as a stand-alone product or can be installed within Red Hat Openshift
    - one can get world-wide support from Red Hat ( an IBM company )
- the DSL used by Ansible is YAML
- YAML is a superset of JSON ( Javascript object notation )
- installation of Ansible is pretty simple compared to Puppet/Chef
- Mastering Ansible will also take less time compared to Puppet/Chef
- Ansible is agent-less
- the servers managed by ansible is called ansible node
  - ansible node can be an Unix/Linux/Mac/Windows machine
- software requirements of Ansible nodes
  - Windows Servers
    - Powershell must be installed
    - WinRM must be configured
    - the above tools normally comes out of the box in any server-grade windows OS
  - Unix/Linux/Mac Servers
    - Python must be installed
    - SSH Server should be installed
    - the above tools normally comes out of the box in any Unix/Linux/Mac Servers
</pre>

## Lab - Cloning the Training repository
```
cd ~/
git clone https://github.com/tektutor/terraform-sep-2025.git
cd terraform-sep-2025
git pull
tree
```

## Lab - Building custom ubuntu docker image to create ansible node containers

Let's generate key pair on the ubuntu lab machine, accept all defaults by hitting enter ( 3 times :) )
```
ssh-keygen 
```
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/1249cc00-aed9-4fd5-a56e-f8aa3fd2d474" />

```
cd ~/terraform-sep-2025/Day1/CustomDockerImage
cp ~/.ssh/id_ed25519.pub authorized_keys
ls -l

docker build -t tektutor/ubuntu-ansible-node:latest .
docker images
```
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/a2416762-22bf-499b-b676-dc2f54794955" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/9cccaba8-9817-4fc5-b8d1-7d06d0a98e34" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/d835c570-ca3a-4c47-808b-311aefc153ed" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/00ee2b97-3c11-4167-adc1-b04d4fc574bf" />

## Lab - Let's create couple of ubuntu ansible containers using our custom ubunut ansible node image
```
docker run -d --name ubuntu1 --hostname ubuntu1 -p 2001:22 -p 8001:80 tektutor/ubuntu-ansible-node:latest
docker run -d --name ubuntu2 --hostname ubuntu2 -p 2002:22 -p 8002:80 tektutor/ubuntu-ansible-node:latest
```

List and see if the ubuntu1 and ubuntu2 ansible node containers are running
```
docker ps
```

Let's ssh into the ubuntu1 container, it shouldn't prompt for password
```
ssh -p 2001 root@localhost
exit

ssh -p 2002 root@localhost
exit
```
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/c8fcea11-0bb4-4241-b524-c0c2d06b8a87" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/25618fa5-8079-49f8-afa4-af9fb36637a7" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/5006ea2b-3e40-4f53-84e9-7f48b443f8d3" />

## Lab - Run ansible ad-hoc command using ping module
```
cd ~/terraform-sep-2025
git pull
cd Day1/ansible
cat inventory
ansible -i inventory all -m ping
```
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/34265c2e-7ac7-4ccd-aebd-5ecaf2196a46" />

