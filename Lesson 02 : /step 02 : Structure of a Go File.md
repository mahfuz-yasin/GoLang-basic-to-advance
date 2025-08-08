
# 📌 Structure of a Go File

A typical Go file consists of **4 main parts**:

---

## 1. **Package Declaration**

* Every Go file **must** start with a `package` declaration.
* The `main` package is special → it’s the **entry point** for executable programs.
* Other packages can be libraries (reusable code).

```go
package main
```

---

## 2. **Import Packages**

* Use `import` to include libraries.
* You can import single or multiple packages.

```go
import "fmt"               // single import
import (                   // multiple imports
    "math"
    "time"
)
```

---

## 3. **Functions**

* Go code is organized into **functions**.
* The `main()` function is the entry point.

```go
func main() {
    fmt.Println("Hello, Go!")
}
```

You can also define your own functions:

```go
func add(x int, y int) int {
    return x + y
}
```

---

## 4. **Statements and Expressions**

* **Statements** → complete instructions (e.g., variable declaration, loops).
* **Expressions** → parts that evaluate to a value (e.g., `x + y`).

### Example:

```go
func main() {
    var a, b int = 10, 20          // Statement
    sum := a + b                   // Expression inside a statement
    fmt.Println("Sum:", sum)       // Statement
}
```

---

# 📘 Complete Example of a Go File

```go
package main   // 1. Package Declaration

import (       // 2. Import Packages
    "fmt"
    "math"
)

func add(x int, y int) int {  // 3. Function
    return x + y
}

func main() {  // Main function → program starts here
    a, b := 15, 25
    result := add(a, b)       // 4. Statement + Expression
    fmt.Println("Result:", result)
    fmt.Println("Square Root of 16:", math.Sqrt(16))
}
```

**Output:**

```
Result: 40
Square Root of 16: 4
```

---

## ✅ Summary Table

| Part                         | Purpose                                    | Example        |
| ---------------------------- | ------------------------------------------ | -------------- |
| **Package declaration**      | Defines the package (main for executables) | `package main` |
| **Import packages**          | Bring in external or standard libraries    | `import "fmt"` |
| **Functions**                | Organize code into reusable blocks         | `func main()`  |
| **Statements & Expressions** | Actual logic of the program                | `x := 5 + 3`   |

---

⚡ In short:
A Go file is like a **recipe** →
1️⃣ Name the recipe (package) →
2️⃣ Bring ingredients (imports) →
3️⃣ Write instructions (functions) →
4️⃣ Mix steps (statements & expressions).


## how create go file and run ?
-   create a file : test.go
-   whrite this command in terminal : go mod init projectName
-   go run fileName
