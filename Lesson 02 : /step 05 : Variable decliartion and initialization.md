
# 📌  Variable Declaration and Initialization in Go

In Go, variables are used to store values in memory. You can declare and initialize them in different ways.

---

## 1. **Declaration with `var` keyword**

* Syntax:

```go
var variableName type
```

* Default **zero value** is assigned if not initialized.

```go
var age int
fmt.Println(age) // Output: 0
```

---

## 2. **Declaration with Initialization**

* Syntax:

```go
var variableName type = value
```

```go
var name string = "Mahfuz"
var score int = 95
fmt.Println(name, score)
```

---

## 3. **Type Inference (without explicit type)**

* If you provide a value, Go can infer the type automatically.

```go
var country = "Bangladesh" // Go infers type as string
var gpa = 3.85             // inferred as float64
fmt.Println(country, gpa)
```

---

## 4. **Short Variable Declaration (`:=`)**

* Only works **inside functions**.
* You don’t need to write `var` or specify type.

```go
name := "Hasan"
age := 21
fmt.Println(name, age)
```

---

## 5. **Multiple Variable Declaration**

* You can declare multiple variables in a single line.

```go
var x, y, z int = 10, 20, 30
fmt.Println(x, y, z)
```

Or with type inference:

```go
a, b, c := 1, 2, 3
fmt.Println(a, b, c)
```

---

## 6. **Block Declaration**

* Useful for grouping related variables.

```go
var (
    studentName string = "Mahfuz"
    studentAge  int    = 21
    isPassed    bool   = true
)
fmt.Println(studentName, studentAge, isPassed)
```

---

## 7. **Zero Values (Default Initialization)**

When no value is provided, Go assigns **zero values**:

| Type    | Zero Value        |
| ------- | ----------------- |
| int     | 0                 |
| float   | 0.0               |
| string  | "" (empty string) |
| bool    | false             |
| pointer | nil               |

```go
var a int
var b float64
var c string
var d bool
fmt.Println(a, b, c, d) // Output: 0 0 "" false
```

---

# 📘 Example Program

```go
package main
import "fmt"

func main() {
    // 1. var with type
    var age int = 21

    // 2. var without explicit type
    var name = "Mahfuz"

    // 3. short declaration
    gpa := 3.85

    // 4. multiple variables
    x, y := 10, 20

    // 5. block declaration
    var (
        country = "Bangladesh"
        passed  = true
    )

    fmt.Println("Name:", name, "| Age:", age, "| GPA:", gpa)
    fmt.Println("x:", x, "y:", y)
    fmt.Println("Country:", country, "| Passed:", passed)
}
```

**Output:**

```
Name: Mahfuz | Age: 21 | GPA: 3.85
x: 10 y: 20
Country: Bangladesh | Passed: true
```

---

# ✅ Quick Recap

1. `var x int` → declares with zero value
2. `var x int = 5` → declares and initializes
3. `var x = 5` → type inferred
4. `x := 5` → short declaration (inside functions)
5. Multiple & block declarations supported
6. Default values = zero values

