<img width="1662" height="758" alt="image" src="https://github.com/user-attachments/assets/2d22c204-7074-4e4d-a71c-88987ce8a6ac" /><img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/9e86f9ba-75bf-4bcf-9b3d-aa5c599a5fc0" /># Day 2

## Lab - Installing AWX 

#### Let's install minikube
```

curl -LO https://github.com/kubernetes/minikube/releases/latest/download/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube && rm minikube-linux-amd64

minikube config set cpus 4
minikube config set memory 12288
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
kubectl get pods -n awx -w
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
kubectl config set-context --current --namespace=awx
kubectl apply -f awx.yml
kubectl logs -f deployments/awx-operator-controller-manager -c awx-manager
kubectl get svc -l "app.kubernetes.io/managed-by=awx-operator"
```

<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/779810b0-ce2d-4491-822f-e893bbbfd7cb" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/65fd6cc5-887b-4724-8d19-2848e440c8b1" />

You may access the ansible tower webconsole
```
http://192.168.49.2:30181
```

Retrieve the password
```
kubectl get secret awx-tower-admin-password -o jsonpath="{.data.password}" | base64 --decode 
```

Login credentials
<pre>
username - admin
password
</pre>

## Info - SOLID Design Principles
<pre>
S - Single Responsibility Principle (SRP)
O - Open Closed Principle (OCP)
    - Your design is good if it satisfies the following desing principles
    - A design/code is considered good if it is open of extension
    - Extending new functionality should be done without modifying existing code
    - To add new functionality you can write new code 
L - Liskov Substitution Principle (LSP)
I - Interface Seggregation
D - Dependency Injection or Dependency Inversion or Inversion of Control(IOC)
</pre>


## Info - Golang Overview
<pre>
- golang is developed by Google using C programming language
- golang is a compiled programming language
- golang syntax resembles very close to C programming language
- golang has 25 keywords
- golang only supports for loop
- just like C/C++/C# main function is the entry-point function ( the very function that will invoked when you run a go application )
- golang supports pointers but memory is managed by garbage collector unlike C/C++
- golang doesn't support class
- golang only supports functions
- using golang one can develop a new compiler/interpreter, a game, console based application, graphical application that runs on your local machine, can develop mobile applications, can develop AI/ML applications, web applications, etc.,
- using golang one can develop REST API, Microservices, etc.,
- is case-sensitive
- statically typed programming language
- performance wise, it is faster than most compiled languages, definitely more faster than interpretted and scripting languages
- even compilation is done faster in go lang for bulky applications
- Some of the popular tools developed in Golang
  - Terraform
  - Docker
  - Kubernetes
  - Openshift  
</pre>

## Info - Golang keywords
<pre>
- break
- default
- func
- interface
- select
- case
- defer
- go
- map
- struct
- chan
- else
- goto
- package
- switch
- const
- fallthrough
- if
- range
- type
- continue
- for
- import
- return
- var
</pre>

## Lab - Installing Golang in Ubuntu
```
cd ~
wget https://go.dev/dl/go1.25.1.linux-amd64.tar.gz
tar xvf go1.25.1.linux-amd64.tar.gz
```

Edit your /home/student/.bashrc and append the below at the end of the file
```
export PATH=$PATH:/home/student/go/bin
export GOROOT=/home/student/go
export GOPATH=/home/student/go/bin
```

To apply the exported variables 
```
source ~/.bashrc
```

Now you may verify the golang version
```
go version
```
It will report
<pre>
go version go1.25.1 linux/amd64  
</pre>


## Lab - Write your first Hello World in golang

Create a file named hello.go with the below content
<pre>
package main

import "fmt"

func main() {
  fmt.Println ("Hello World !")
}
</pre>

Run your application
```
go run ./hello.go
```
<img width="1536" height="506" alt="image" src="https://github.com/user-attachments/assets/39ca5b93-e2e5-4b17-a65b-9a0ca898f12b" />

## Lab - Accepting user inputs

Create a file named userinputs.go with the below content
<pre>
package main

import "fmt"

func main() {
	//The below line declares variables named x and y, later it is initialized with value 0
	x := 0
	y := 0

	fmt.Print("Enter your first integer :" )
	fmt.Scanf("%d", &x )

	fmt.Print("Enter your second integer :" )
	fmt.Scanf("%d", &y )

	fmt.Println("Value of x :", x )
	fmt.Println("Value of y :", y )

	//Declares a variable named tmp of type string
	var temp string
	fmt.Println("Press any key to exit ...")
	fmt.Scanln(&temp)
}  
</pre>

Run your application
```
go run ./userinputs.go
```

<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/c5457c8f-aee7-4ed2-904b-8bbc2f477fdf" />

## Lab - if else

Create file named if-else.go with the below code
<pre>
package main

import "fmt"

func main() {
	x := 100

	if x % 2 == 0 {
		fmt.Println(x, "is an even number")
	} else {
		fmt.Println(x, "is an odd number")
	}
}	
</pre>

Run your application
```
go run ./if-else.go
```
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/d87bde30-250e-46aa-a88e-cf1149801569" />

## Lab - Golang array

Create a file named arrays.go with the below code
<pre>
package main

import "fmt"

func main() {

	//We have declared an array of integers of size 5
	//So we can store upto 5 integer values into this array
	//go lang array size is fixed
	//array index starts from 0
 	//valid array index range is 0 to 4, total 5 values
	var arr [5]int

	//let's assign some values into the array
	arr[0]= 10
	arr[1]= 20
	arr[2]= 30
	arr[3]= 40
	arr[4]= 50

	//arr[5] = 60 This will report array index out of bounds error

	fmt.Println("Array elements are ...")
	fmt.Println(arr)

	count := len(arr)
	fmt.Println("Length of array :", count)
	
	//Modifying values stored in an array
	arr[3] = 25

	fmt.Println("Array elements are ...")
	for i := 0; i < count; i++ {
		fmt.Printf("%d\t", arr[i])
	}
	fmt.Println()
}	
</pre>

Run your application
```
go run ./arrays.go
```

<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/2a671bb5-9916-4e74-8adc-09ab97cec606" />

<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/9e5eb544-cb8b-4f0b-9fbc-d10ee14994f9" />

## Lab - Error Handling in golang
Create a file named error-handling.go with the below code
<pre>
package main

import (
	"fmt"
	"os"
	"os/user"
	"github.com/sttk/stringcase"
)

func main() {
	u, err := user.Current()

	if err != nil {
		fmt.Println("Cannot get current user:", err)
		os.Exit(1)
	}

	fmt.Printf("Hello %s, welcome !\n", stringcase.PascalCase(u.Username) )
}
</pre>

Run your application
```
go mod init main
cat go.mod
go mod tidy
go run ./error-handling.go
```

<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/2df64a25-8c8d-47d7-8756-c3259897407a" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/fe05cd5e-8af4-42ad-9f45-d6a1c51389ae" />

## Lab - User defined functions

Create a file named functions.go with the below code
<pre>
package main

import "fmt"

func yetAnotherFunction() {
	fmt.Println("Yet Another Function invoked")
}

func main() {
   fmt.Println( sayHello("Golang") )
   fmt.Println( sayHello("World") )
   yetAnotherFunction()
}

//This function accepts a string input argument and returns a string output
func sayHello( msg string ) string {
  return "Hello, " + msg + " !"
}

/* function overloading is not supported by golang
func sayHello() string {
  return "Hello World !"
}
*/	
</pre>

Run your application
```
go run ./functions.go
```

<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/85c65d54-d2ec-4be3-83ce-61c4ce4d9bfe" />

## Lab - Functions with multiple returns

Create a file named funciton-with-multiple-returns.go with below code
<pre>
package main

import "fmt"

func myFunction() (int,int) {
  return 10, 20
}

func main() {
  x, y := myFunction() // := is a short form of declaring a new variable and initialized with some value

  fmt.Println( "Value of x is ", x )
  fmt.Println( "Value of y is ", y )
}	
</pre>
Run your application
```
go run ./function-with-multiple-returns.go
```
<img width="1662" height="758" alt="image" src="https://github.com/user-attachments/assets/fefcc3eb-fb14-4c1f-92c2-dff062ba2133" />

