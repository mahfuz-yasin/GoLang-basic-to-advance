Excellent 👍 — let’s break down **Formatting Output** and **Constants in Go** step by step with examples.

---

# 📌 1. Formatting Output in Go

Go uses the **`fmt` package** for input/output operations.

### ✅ Common Output Functions

| Function        | Description            | Example                                 |
| --------------- | ---------------------- | --------------------------------------- |
| `fmt.Print()`   | Prints without newline | `fmt.Print("Hello")` → `Hello`          |
| `fmt.Println()` | Prints with newline    | `fmt.Println("Hello")` → `Hello\n`      |
| `fmt.Printf()`  | Formatted printing     | `fmt.Printf("Age: %d", 21)` → `Age: 21` |

---

### ✅ Format Specifiers in `fmt.Printf`

| Specifier | Meaning                   | Example                      | Output       |
| --------- | ------------------------- | ---------------------------- | ------------ |
| `%d`      | Decimal integer           | `fmt.Printf("%d", 25)`       | 25           |
| `%f`      | Floating number           | `fmt.Printf("%.2f", 3.1416)` | 3.14         |
| `%s`      | String                    | `fmt.Printf("%s", "Go")`     | Go           |
| `%t`      | Boolean                   | `fmt.Printf("%t", true)`     | true         |
| `%c`      | Character                 | `fmt.Printf("%c", 65)`       | A            |
| `%v`      | Default format (any type) | `fmt.Printf("%v", 42)`       | 42           |
| `%T`      | Type of the variable      | `fmt.Printf("%T", 42)`       | int          |
| `%p`      | Memory address (pointer)  | `fmt.Printf("%p", &x)`       | 0xc000014090 |
| `%%`      | Prints % symbol           | `fmt.Printf("100%%")`        | 100%         |

---

### 📘 Example: Formatting Output

```go
package main
import "fmt"

func main() {
    name := "Mahfuz"
    age := 21
    gpa := 3.75
    pass := true

    fmt.Print("Hello, ", name)
    fmt.Println(" is a Go programmer.")
    fmt.Printf("Name: %s | Age: %d | GPA: %.2f | Pass: %t\n", name, age, gpa, pass)
}
```

**Output:**

```
Hello, Mahfuz is a Go programmer.
Name: Mahfuz | Age: 21 | GPA: 3.75 | Pass: true
```

---

# 📌 2. Constants in Go

A **constant** is a variable whose value **cannot be changed** once assigned.
Declared using the `const` keyword.

### ✅ Declaring Constants

```go
const Pi = 3.1416
const Greeting string = "Assalamu Alaikum"
```

---

### ✅ Multiple Constants

```go
const (
    A = 1
    B = 2
    C = 3
)
```

---

### ✅ Typed vs Untyped Constants

* **Typed Constant** → type specified explicitly
* **Untyped Constant** → type inferred when used

```go
const typedConst float64 = 3.14 // typed
const untypedConst = 100        // untyped
```

---

### ✅ Constant Expressions

Constants can be created from expressions (evaluated at compile time).

```go
const X = 5
const Y = X + 10
```

---

### ✅ `iota` (Special Constant Generator)

`iota` is a special identifier used for generating sequential values.
Each new constant in a block increments `iota` by 1.

```go
const (
    A = iota // 0
    B        // 1
    C        // 2
)
```

**Output:** `A=0, B=1, C=2`

---

# 📘 Example: Constants in Action

```go
package main
import "fmt"

const Pi = 3.1416
const (
    Sunday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)

func main() {
    fmt.Println("Pi:", Pi)
    fmt.Println("Tuesday:", Tuesday)
}
```

**Output:**

```
Pi: 3.1416
Tuesday: 2
```

---

# ✅ Quick Recap

### **Formatting Output**

* Use `fmt.Print`, `fmt.Println`, and `fmt.Printf`.
* Common format specifiers: `%d`, `%f`, `%s`, `%t`, `%v`, `%T`.

### **Constants**

* Declared with `const`.
* Immutable after declaration.
* `iota` helps auto-generate sequences.
