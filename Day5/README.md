# Day 5

## Info - Ansible Automation Platform
<pre>
- was also called as Ansible Tower
- it is developed on top the AWX (open source)
- AWX is developed on top Ansible core (open source)
- unlike Ansible core, Ansible automation platform supports webconsole, user management, etc.,
- this is a Red Hat Enterprise product with world-wide support
- functionally AWX and Ansible Automation Platform are one and same
- it is not possible develop Ansible Playbook within AWX or Ansible Automation Platform
- hence, we still need Ansible core to develop playbooks and test them before we push it to GitHub or any version control
- the existing Ansible Playbooks we can execute via Ansible Automation Platform
- can be installed as a stand-alone application on any Linux Distributions or we can install inside Kubernetes or Openshift
</pre>

## Lab - Starting the minikube ( Launch Ansible Tower )
```
docker ps -a
minikube start
minikube status
kubectl get nodes
```

Expected output
![image](img1.png)
![image](img2.png)

Accessing Ansible Tower Dashboard from chrome web browser on your RPS Cloud Ubuntu lab machine
```
minikube service awx-demo-service --url -n ansible-awx
```
You can launch the AWX Webconsole using the url shown in the above command
```
http://192.168.49.2:31225
```
![image](img3.png)

Ansible Tower Login Credentials, save the login credentials in a text file to avoid typing complex command each time
<pre>
username - admin
password - 
</pre>

To get the password, you to type the below command
```
kubectl get secret awx-demo-admin-password -o jsonpath="{.data.password}" -n ansible-awx | base64 --decode; echo
```

Once you login, you will get the below page
![image](img4.png)

## Lab - Creating a Project in Ansible AWX

Navigate to Ansible Automation Platform on your RPS Lab machine - chrome web browser
![image](img5.png)

On the menu that appears on the left side, Navigate to Resources --> Projects
![image](img6.png)

Click "Add"
![image](img7.png)
<pre>
Under the Name, type "TekTutor Training Repository"
Under the Source Code Type, select "Git"
Under the Source Control url, type "https://github.com/tektutor/terraform-may-2025.git"
Under the Source Control Branch/Tag/Commit, type "main"
Under the Options, enable the check box that says "Update Revision on Launch"
</pre>
![image](img8.png)
Click "Save"
![image](img9.png)
![image](img10.png)
Click "Successful"
![image](img11.png)
![image](img12.png)
![image](img13.png)

## Lab - Creating Inventory in Ansible Automation Platform(AWX)

Navigate to Ansible Automation Platform(AWX)
![image](img14.png)

Click Resources --> Inventories
![image](img15.png)
Click Add
Select the first option "Add Inventory
![image](img16.png)
![image](img17.png)
![image](img18.png)
Click "Save"
![image](img19.png)
Click the Tab named "Hosts" within the Inventory you saved just now
![image](img20.png)
Click "Add"
![image](img21.png)
![image](img22.png)
![image](img23.png)
Click "Save"
![image](img24.png)
![image](img25.png)
click Add to create other ansible nodes on the similar fashion
![image](img26.png)
![image](img27.png)
Click "Save"
![image](img28.png)
click Add to create other ansible nodes on the similar fashion
![image](img29.png)
Click "Add"
![image](img30.png)
![image](img31.png)
Click "Save"
![image](img32.png)

Repeat the procedure to add "Rocky2"
![image](img33.png)
![image](img34.png)
![image](img35.png)

To verify if all the hosts(ansible nodes) added to the inventory are reachable to Ansible Tower, Click on your inventory and move to the Hosts tab
![image](img36.png)
Click "Run command"
![image](img37.png)
Under the Module, choose "ping"
![image](img38.png)
![image](img39.png)
Click "Next"
![image](img40.png)
Click "Next"
![image](img41.png)
Select "RPS Private Key" we saved
Click "Next"
![image](img42.png)
Click "Launch"
![image](img43.png)
![image](img44.png)
![image](img45.png)


## Lab - Creating Credentials to store the Private key 
Navigate to Ansible Tower Dashboard
![image](img46.png)

Click Resources --> Credentials
![image](img47.png)
Click "Add"
![image](img48.png)
![image](img49.png)
Select "Machine" Credential Type
![image](img50.png)
Open your RPS Cloud Machine Terminal, type "cat /home/rps/.ssh/id_ed25519"
![image](img51.png)
Copy the private key including the Begin and End as shown below
![image](img52.png)
Paste the private key you copied under the "SSH Private Key" field (Remove extra space)
![image](img53.png)
Scroll down to save
![image](img54.png)
Click Save
![image](img55.png)

## Lab - Creating Job Template to invoke a playbook from Ansible Tower
Navigate to Ansible Tower Dashboard
![image](img56.png)

Click "Resources->Templates"
![image](img57.png)
Click "Add"
![image](img58.png)
Select "Add Job Template"
![image](img59.png)
<pre>
Under the Name, type "Install nginx playbook"
Click Search in Inventory and select "Docker Inventory" that we created
</pre>
![image](img60.png)
![image](img61.png)

Click Search in Project and Select "TekTutor Training Repository"
![image](img62.png)
![image](img63.png)
![image](img64.png)
Under the Playbook, select "Day2/ansible/after-refactoring/install-nginx-playbook.yml"
![image](img65.png)
Under Credential, click search and select "RPS private key file"
![image](img66.png)
![image](img67.png)
![image](img68.png)
Scroll down and click "Save"
![image](img69.png)


To run the playbook, click "Launch" Button
![image](img70.png)
![image](img71.png)
![image](img72.png)
![image](img73.png)
![image](img74.png)
![image](img75.png)

## Info - Jenkins
<pre>
- is a CI/CD Build Server
- developed by Josuke Kavaguchi, a former employee of Sun Microsystems
- initially this tool was called as Hudson, when originally it was developed in Java by Josuke Kavaguchi
- Josuke started a company called Cloudbees, forked the Hudson branch as Jenkins
- Jenkins opensource project today has got more than 10000 active opensource contributors
</pre>

## Lab - Setup Jenkins
Download jenkins from command-line
<pre>
cd ~
wget https://get.jenkins.io/war-stable/2.516.3/jenkins.war  
</pre>

Install Java
```
sudo apt install -y openjdk-21-jre-headless
```


Launching Jenkins
<pre>
cd ~
java -jar ./jenkins.war  
</pre>
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/783c7472-6d36-402e-bcdd-d84ee81bb544" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/39377a72-10c6-4f86-9804-dd96ccd2cf8e" />

You can access your jenkins from chrome browse on the lab machine 
```
http://localhost:8080
```
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/7944e0a8-c004-4c8c-864f-a7019897f83d" />
