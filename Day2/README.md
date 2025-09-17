# Day 2 

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
kubectl get secret awx-tower-admin-password -n awx -o jsonpath='{.data.password}' | base64 -d; echo
```

Login credentials
<pre>
username - admin
password
</pre>

#### Troubleshooting pods crash and AWX Dashboard login failure

Create a file named kustomization.yaml with below code
<pre>
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
namespace: awx
resources:
  - github.com/ansible/awx-operator/config/default?ref=2.19.1

images:
  - name: quay.io/ansible/awx-operator
    newTag: 2.19.1
</pre>

Apply
```
kubectl apply -k .
```
<img width="1944" height="1193" alt="image" src="https://github.com/user-attachments/assets/6254f755-5e90-4c28-bd5b-3786c583a69e" />
<img width="1944" height="1193" alt="image" src="https://github.com/user-attachments/assets/17d20c4c-fd43-4834-8b3e-b0fd9e65952e" />

Delete your awx-tower-task-xxyyzz and awx-tower-web-xxyyzz pod
```
kubectl delete pod awx-tower-task-xxyyzz awx-tower-web-xxyyzz
```
<img width="1944" height="1193" alt="image" src="https://github.com/user-attachments/assets/28596b2a-3377-47c3-847b-0fa89db962c5" />

Wait until postgresdb upgrades to 15.0 and awx-tower-task and awx-tower-web pod turns to running state

Once you see
<pre>
kubectl get pods -n awx -w

NAME                                               READY   STATUS      RESTARTS   AGE
awx-operator-controller-manager-58b7c97f4b-p58r2   2/2     Running     0          9m5s
awx-tower-migration-24.6.1-k7tnl                   0/1     Completed   0          5m15s
awx-tower-postgres-15-0                            1/1     Running     0          7m35s
awx-tower-task-5f55685d77-crlp5                    4/4     Running     0          6m13s
awx-tower-web-75dbfddbf-h5gmg                      3/3     Running     0          6m15s
</pre>
	
Retrieve your admin password
```
kubectl get secret awx-tower-admin-password -n awx -o jsonpath='{.data.password}' | base64 -d; echo
```
<img width="1944" height="1193" alt="image" src="https://github.com/user-attachments/assets/c7883b06-1dec-4720-9291-4c793be90afd" />

<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/54e7f522-e6b4-4fe4-ac93-d8e2fa6a861f" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/72db3ac9-cad6-4c5a-bb20-74c407a60306" />
<img width="1944" height="1193" alt="image" src="https://github.com/user-attachments/assets/3db70544-7db7-409f-a5e3-94bfc9e44e69" />

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

## Lab - Golang loops 
The only type of loop supported in golang is for loop, but it supports many variations.  Hence, with for loop we can achieve whatever we do with regular for loop, while-do, do-while, etc.,

Create a file named loops.go with the below code
<pre>
package main

import "fmt"

func main() {
	count := 5 //Declares a count variable of type int and assigns a value 5

	//similar to while loop
	for count > 0 {
		fmt.Println("Before decrementing ", count)

		count-- //equivalent to count = count - 1
		//golang doesn't support pre-decrement or pre-increment unlinke C and C++

		fmt.Println("After decrementing ", count )
	}
	fmt.Println("Value of count is ", count, " after for loop")

	count = 0 //variable is already declared in line number 7, we are just trying to reset the value to 0 here

	fmt.Println()
	//Regular for loop
	for count=1; count<10; count++ {
		fmt.Printf( "%d\t", count )
	}
	fmt.Println()

	//similar to while do
	count = 0

	for {  //Infinite loop

		fmt.Printf ("Inside for loop %d\n", count ) 	
		count++

		if count > 3 {
			break
		}
	}
	fmt.Println("Control reached outside for infinite for loop")
}

</pre>

Run your application
```
go run ./loops.go
```

<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/cbfe58a2-fd87-4573-b706-f61e58b8a144" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/2aac4fcd-add3-4d23-9bd7-a71aa9c5336b" />

## Lab - Golang map
Create a file named maps.go with the below code
<pre>
package main

import "fmt"

func main() {

	toolsPath := map[string]string {
		"java_home": "/usr/lib/jdk11",
		"mvn_home": "/usr/share/maven",
	}

	fmt.Println("Java home directory ", toolsPath["java_home"])

	//add a key,value pair into an existing map
	toolsPath["go_home"] = "/usr/go"

	//iterating a map and printing its values
	for key,value := range toolsPath {
		fmt.Println(key,value)
	}

	//delete a key-value pair from an existing map
	delete(toolsPath, "go_home")
	fmt.Println(toolsPath)
}	
</pre>

Run it
```
go run ./maps.go
```

<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/d7ed88bc-8b62-4b11-89eb-74d7565f480a" />

## Lab - Pointers in golang

Create a file named pointers.go with the below code
<pre>
package main

import "fmt"

func sayHello( msgPtr *string ) {

	//Dereferencing - the values stored at address pointed by msgPtr will be printed here
	fmt.Println("Inside sayHello function ", *msgPtr )

	//Here the address pointed by msgPtr pointer will be printed
	fmt.Println("Address pointed by msgPtr is", msgPtr )

	//Prints the address of msgPtr
	fmt.Println("Address of msgPtr is", &msgPtr ) 

	//The value stored at the address pointed by msgPtr is assigned to tmp string
	tmp := *msgPtr

	*msgPtr = tmp + " Golang" + " !"

	fmt.Println("Inside sayHello before return ", *msgPtr) 
}

func main() {
	//declares a string variable named str with value "Hello"
	msg := "Hello"

	fmt.Println("Message before calling sayHello function is ", msg )
	fmt.Println("Address of msg string is ", &msg)

	sayHello( &msg )

	fmt.Println("Message after calling sayHello function is ", msg )
}	
</pre>

Run it
```
go run ./pointers.go
```

<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/b1a96d5f-bc5c-41dd-a57f-fa682df90b4a" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/d2cc2fc7-cc8b-465d-8d65-c51f514ea688" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/c6addf7d-a776-4924-b390-f3ca4e458c40" />

## Lab - Golang struct with Methods

Create a file named struct.go with the below code
<pre>
package main

import "fmt"

type Rectangle struct {
  length int
  width int
}

func ( rect Rectangle ) Area() int {
   area := rect.length * rect.width
   return area
}

func (rect Rectangle) GetLength() int {
  return rect.length
}

func (rect Rectangle) GetWidth() int {
  return rect.width
}

func main() {
   rectangle := Rectangle {
      length: 100,
      width : 200, 
   }

   fmt.Printf("Length of rectangle : %d\n", rectangle.GetLength() )
   fmt.Printf("Width of rectangle  : %d\n", rectangle.GetWidth() )
   fmt.Printf("Area of rectangle   : %d\n", rectangle.Area() )

}
</pre>

Run it
```
go run ./struct.go
```


<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/99d099dd-8b3a-4ce9-87df-b97ac68bc25f" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/80b08b1f-815a-4b98-b029-7b1e93b1fe2d" />

## Lab - Golang switch case

Create a file named switch-case.go with the below code
<pre>
package main

import "fmt"

func main() {

	var direction string

	fmt.Println("Possible values are east,west,south,north")

	fmt.Print("Enter some direction :")
	fmt.Scanln(&direction)

	switch direction {
	case "east":
		fmt.Println("You enterred direction ", direction)
	case "west":
		fmt.Println("You enterred direction ", direction)
	case "south":
		fmt.Println("You enterred direction ", direction)
	case "north":
		fmt.Println("You enterred direction ", direction)
	default:
		fmt.Println("Invalid direction", "possible values are east, west, north, south")
	}
}	
</pre>

Run it
```
go run ./switch-case.go
```

<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/46dac711-b48a-44c1-9469-fd5a265f453c" />

## Lab - Golang slice

Create a file named slice.go with the below code
<pre>
package main

import "fmt"

func main() {
        //index             0  1  2  3  4  5
	intArray := [6]int{10,20,30,40,50,60}

	fmt.Println ( "Array elements are ...")
	fmt.Println(intArray)

	var mySlice []int = intArray[2:5] //2 is the lower bound index, while 5 is hte upper bound index, index 5 is not included

	fmt.Println("Slice elements are ...")
	fmt.Println(mySlice)

	//Let's modify the slice at certain indices
	//when the slice is modified, it will affect the original array that is pointed by the slice
	mySlice[0] = 100 //mySlice[0] is nothing but intArray[2]
	mySlice[1] = 200 //mySlice[1] is nothing but intArray[3]
	mySlice[2] = 300 //mySlice[2] is nothing but intArray[4]

	fmt.Println("Slice elements after modifying slice are ...")
	fmt.Println(mySlice)

	fmt.Println("Array elements after modifying slice are ...")
	fmt.Println(intArray)

	mySlice = append(mySlice, 400)
	fmt.Println("Array elements after appending 400 in slice are ...")
	fmt.Println(intArray)

	mySlice = append(mySlice, 500) //after this append, mySlice is no more associated/pointing to intArray
	fmt.Println("Array elements after appending 500 in slice are ...")
	fmt.Println(intArray)

	fmt.Println("Slice elements after appending 500 into slice are ...")
	fmt.Println(mySlice)

}	
</pre>

Run it 
```
go run ./slice.go
```

<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/55d3add8-cd41-42ac-a77b-cdaaa0e23451" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/51d694c5-7dd3-440e-9f07-7b888b7d0802" />

## Lab - Creating custom Golang modules

Let's create three modules namely main, addition and subtraction

```
cd ~
mkdir custom-go-module
cd custom-go-module
mkdir addition subtraction
```

Let's create a module named addition
```
cd ~/custom-go-module/addition
go mod init addition //Creates a file named go.mod with name of the module and golang version supported
```

Under the addition folder, let's create afile called add.go with the below code
<pre>
package addition

func Add( x float32, y float32 ) float64 {
   return float64( x + y )
}
</pre>

Under the subtraction folder, let's create
```
cd ~/custom-go-module/subtraction
go mod init subtraction //Creates a file named go.mod with the name of the module and golang version supported
```

Under the subtraction folder, let's create a file named subtract.go with the below code
<pre>
package subtraction

func Subtract( x float32, y float32 ) float64 {
	return float64( x - y )
}
</pre>

Let's create the main module
```
cd ~/custom-go-module
go mod init main
```

Let's create a file named main.go with the below code
<pre>
package main

import (
   "fmt"
   "addition"
   "subtraction"
)

func main() {
	//x := 100.123 By default, golang will assume 100.123 as float64

	x := float32(100.123) // We are casting/converting float64 into float32
	y := float32(200.456) // We are casting/converting float64 into float32

	fmt.Println ( "The sum of ", x, " and ", y , " is ", addition.Add( x, y ) )
	fmt.Println ( "The difference of ", x, " and ", y , " is ", subtraction.Subtract( x, y ) )
}
</pre>

Run it
```
cd ~/custom-go-module
go mod edit --replace addition=./addition //This helps golang to locate the addition module locally
go mod edit --replace subtraction=./subtraction //This helps golang to locate the subtraction module locally
go mod tidy //This will download all the dependent modules of main.go and any other *.go files in main module

go run ./main.go
```

## Lab - Golang concurrency

Create a file named concurrency.go with the below code
<pre>
package main

import (
   "fmt"
   "time"
)

func firstFunction( count int ) {
   for i := 0; i < count; i++ {
	fmt.Println("First function ", i )
	time.Sleep(time.Millisecond * 5 )
   }
}

func secondFunction( count int ) {
   for i := 0; i < count; i++ {
	fmt.Println("Second function ", i )
	time.Sleep(time.Millisecond * 5 )
   }
}

func main() {
	fmt.Println ( "Press any key to exit ...")

	//Invoking firstFunction and secondFunction in sequence one after the other
	firstFunction(10)
	secondFunction(10)

	//We wish to run both firstFunction and secondFunction in parallel
	go firstFunction( 1000 )
	go secondFunction( 1000 )

	var tmp string
	fmt.Scanln(&tmp) //this will make sure the program waits until some key is pressed to exit
}	
</pre>

Run it
```
go run ./concurrency.go
```

<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/9d8e8839-91b6-4c69-95eb-232e385a244c" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/3fb2421a-c9a4-4043-bf83-34578a34e830" />
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/2cbb2249-1bf8-48ff-b242-033b1bcf7164" />

## Lab - Building go application into a binary ( useful while shipping your application to clients )
```
cd ~/terraform-sep-2025
git pull
cd Day2/golang/concurrency
ls
go build
ls
./concurrency
```
<img width="1920" height="1168" alt="image" src="https://github.com/user-attachments/assets/ad9c7758-0470-441a-b2e4-772025df90ed" />
