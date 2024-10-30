# Create a GO Module and Call a GO module Notes
In bash, home directory:  
mkdir ~/greetings, cd greetings  
``go mod init example.com/greetings``, go.mod file is created in ~/greetings  
#### The go mod init command creates a go.mod file to track your code's dependencies. So far, the file includes only the name of your module and the Go version your code supports. But as you add dependencies, the go.mod file will list the versions your code depends on. This keeps builds reproducible and gives you direct control over which module versions to use.  
nano greetings.go, pasted the code into the file and saved.  
mkdir ~/hello, cd hello  
``go mod init example.com/hello``  
nano hello.go, paste the code into the file and save.  
#### For production use, you’d publish the example.com/greetings module from its repository (with a module path that reflected its published location), where Go tools could find it to download it.   
For now, because you haven't published the module yet, you need to adapt the example.com/hello module so it can find the example.com/greetings code on your local file system.  
From ~/hello directory: ``go mod edit -replace example.com/greetings=../greetings``  
The command specifies that example.com/greetings should be replaced with ../greetings for the purpose of locating the dependency. After you run the command, the go.mod file in the hello directory should include a replace directive.  
#### From the command prompt in the hello directory, run the go mod tidy command to synchronize the example.com/hello module's dependencies, adding those required by the code, but not yet tracked in the module.  
## greetings.go
```
package greetings

import "fmt"

// Hello returns a greeting for the named person.
func Hello(name string) string {
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf("Hi, %v. Welcome!", name)
    return message
}
```
This function takes a name parameter whose type is string. The function also returns a string.  
### In Go, a function whose name starts with a capital letter can be called by a function not in the same package. This is known in Go as an exported name.  
Function name: "Hello"  
Parameter type: "string"  
Return type: "string"  
#### In Go, the := operator is a shortcut for declaring and initializing a variable in one line (Go uses the value on the right to determine the variable's type).  
Taking the long way, you might have written this as:  
```
var message string
message = fmt.Sprintf("Hi, %v. Welcome!", name)
```
## hello.go
```
package main

import (
    "fmt"

    "example.com/greetings"
)

func main() {
    // Get a greeting message and print it.
    message := greetings.Hello("Gladys")
    fmt.Println(message)
}
```
Declare a main package. In Go, code executed as an application must be in a main package.  
Import two packages: example.com/greetings and the fmt package. This gives your code access to functions in those packages.  
Importing example.com/greetings (the package contained in the module you created earlier) gives you access to the Hello function.  
Get a greeting by calling the greetings package’s Hello function.
