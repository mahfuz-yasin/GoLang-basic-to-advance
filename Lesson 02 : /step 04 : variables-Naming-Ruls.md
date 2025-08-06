
# 📌 Go Variable Naming Rules

### ✅ 1. Must Start with a Letter or Underscore

A variable name must begin with a **letter (a–z / A–Z)** or an **underscore `_`**.

```go
var name string = "Mahfuz" // valid
var _score int = 100       // valid
```

---

### ❌ 2. Cannot Start with a Digit

```go
var 1number int = 5  // ❌ invalid
var number1 int = 5  // ✅ valid
```

---

### ✅ 3. Can Only Contain Letters, Digits, and Underscores

```go
var student_name string = "Hasan" // valid
var student1 int = 21             // valid
var student-name string = "Ali"   // ❌ invalid (hyphen not allowed)
```

---

### ✅ 4. Case-Sensitive

`age`, `Age`, and `AGE` are different variables.

```go
var age = 20
var Age = 30
var AGE = 40
fmt.Println(age, Age, AGE) // Output: 20 30 40
```

---

### ✅ 5. No Limit on Length

You can make variable names as long as you want,
but **short & meaningful** names are recommended.

```go
var myVeryLongVariableNameThatDescribesThePurpose int = 100
```

---

### ❌ 6. Cannot Contain Spaces

```go
var first name string = "Mahfuz" // ❌ invalid
var first_name string = "Mahfuz" // ✅ valid
```

---

### ❌ 7. Cannot Be a Go Keyword

Go has reserved **keywords** that you cannot use as variable names.
Examples of reserved keywords:
`break`, `case`, `chan`, `const`, `continue`, `default`, `defer`, `else`, `for`, `func`, `go`, `if`, `import`, `interface`, `map`, `package`, `return`, `select`, `struct`, `switch`, `type`, `var`

```go
var func = 10    // ❌ invalid
var myFunc = 10  // ✅ valid
```

---

## 📘 Summary Table

| Rule                         | Example ✅        | Example ❌    |
| ---------------------------- | ---------------- | ------------ |
| Start with letter/underscore | `_name`, `user1` | `1user`      |
| No starting digit            | `score99`        | `9score`     |
| Only a-z, A-Z, 0-9, \_       | `first_name`     | `first-name` |
| Case-sensitive               | `age`, `Age`     | N/A          |
| No spaces                    | `user_name`      | `user name`  |
| Not a keyword                | `user`           | `package`    |

----


# 📌 Multi-Word Variable Names in Go

## 1. **Camel Case** (Preferred in Go for local variables)

* The first word starts with lowercase, and each subsequent word starts with uppercase.
* Very common in Go.

```go
var studentName string = "Mahfuz"
var totalMarks int = 95
```

---

## 2. **Pascal Case** (Used for Exported Identifiers)

* Every word starts with uppercase.
* Used when you want the variable/function/struct to be **exported** (public) outside the package.

```go
var StudentName string = "Hasan"
var TotalMarks int = 90
```

---

## 3. **Snake Case** (Rare in Go, but allowed)

* Words separated by underscores.
* Usually seen in constants or legacy code, not common for variables.

```go
var student_name string = "Ali"
var total_marks int = 85
```

---

## 4. **All Caps with Underscores** (For Constants)

* Recommended for global constants.

```go
const MAX_SCORE = 100
const PI_VALUE = 3.1416
```

---

# ✅ Best Practices (Go Style Guide)

* **Local variables** → use **camelCase**
* **Exported (public) variables, functions, structs** → use **PascalCase**
* **Constants** → use **ALL\_CAPS with underscores**
* Avoid **snake\_case** unless absolutely necessary

---

## 📌 Example Program

```go
package main
import "fmt"

const MAX_SCORE = 100 // constant in ALL_CAPS

func main() {
    var studentName string = "Mahfuz" // camelCase
    var StudentAge int = 21           // PascalCase (exported)
    var totalMarks int = 90           // camelCase
    
    fmt.Println("Name:", studentName)
    fmt.Println("Age:", StudentAge)
    fmt.Println("Marks:", totalMarks, "/", MAX_SCORE)
}
```

**Output:**

```
Name: Mahfuz
Age: 21
Marks: 90 / 100
```

---

⚡ **Tip:**

* Use **camelCase** inside functions (local scope).
* Use **PascalCase** for identifiers you want to share with other packages.
