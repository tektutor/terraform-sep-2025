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
```

