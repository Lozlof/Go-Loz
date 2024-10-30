In Go, `fmt.Print`, `fmt.Printf`, and `fmt.Println` each have different purposes and behaviors for formatting and printing output:

### 1. `fmt.Print`

- Prints the arguments as they are, with no additional formatting.
- It **does not add a newline** at the end.
  
Example:
```go
fmt.Print("Hello")
fmt.Print(" World")
```
Output:
```
Hello World
```

### 2. `fmt.Println`

- Similar to `fmt.Print`, but it **adds a newline** at the end.
- Automatically inserts a space between arguments if they are separated by commas.

Example:
```go
fmt.Println("Hello")
fmt.Println("World")
```
Output:
```
Hello
World
```

### 3. `fmt.Printf`

- Allows for formatted printing, similar to `printf` in C.
- You can use **format specifiers** like `%s` for strings, `%d` for integers, `%f` for floating-point numbers, etc.
- Does not add a newline at the end by default, so you need to include `\n` if you want one.

Example:
```go
name := "Alice"
age := 25
fmt.Printf("Name: %s, Age: %d\n", name, age)
```
Output:
```
Name: Alice, Age: 25
```

### Summary

| Function    | Adds Newline | Inserts Space Between Args | Formatting Options |
|-------------|--------------|----------------------------|--------------------|
| `fmt.Print` | No           | No                         | No                 |
| `fmt.Println` | Yes         | Yes                        | No                 |
| `fmt.Printf` | No (unless `\n` included) | No            | Yes                |

Choose `fmt.Print` for basic output, `fmt.Println` for automatic newlines and spacing, and `fmt.Printf` when you need formatted output.
