# :=        
#### In Go, := is a short variable declaration syntax.       
It is used to declare and initialize one or more variables in a concise way.        
The := operator is typically used within functions and can only be used when Go can infer the type of the variable from the right-hand side expression.        
**message := "Hello, Go!"**     
**var message string = "Hello, Go!"**        
## Key Points       
- **Type inference:** Go automatically determines the type of the variable based on the value on the right-hand side.          
- **Local scope:** It can only be used inside functions, not for global variable declarations.      
- **Must be used with a new variable:** You can't use := with an already declared variable unless at least one of the variables on the left-hand side is new.          