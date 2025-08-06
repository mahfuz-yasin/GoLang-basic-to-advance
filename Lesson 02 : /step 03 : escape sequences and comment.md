
# 📌 Escape Sequences in Go

Escape sequences are **special character combinations** used inside strings.
They usually begin with a **backslash (`\`)**.

## 🔑 Common Escape Sequences

| Escape Sequence | Meaning                | Example                     | Output                  |
| --------------- | ---------------------- | --------------------------- | ----------------------- |
| `\n`            | New line               | `"Hello\nWorld"`            |                         |
| Hello           |                        |                             |                         |
| World           |                        |                             |                         |
| `\t`            | Horizontal tab         | `"Hello\tWorld"`            | Hello World             |
| `\\`            | Backslash              | `"This is a backslash: \\"` | This is a backslash: \\ |
| `\"`            | Double quote           | `"She said, \"Hello\""`     | She said, "Hello"       |
| `\'`            | Single quote (in rune) | `'\''`                      | '                       |
| `\r`            | Carriage return        | `"Hello\rWorld"`            | World                   |
| `\b`            | Backspace              | `"Hello\bWorld"`            | HellWorld               |
| `\f`            | Form feed              | Rarely used                 | —                       |
| `\v`            | Vertical tab           | Rarely used                 | —                       |
| `\uXXXX`        | Unicode code point     | `"\u03A9"`                  | Ω                       |
| `\UXXXXXXXX`    | Unicode (32-bit)       | `"\U0001F600"`              | 😀                      |

---

### 📘 Example: Escape Sequences in Action

```go
package main
import "fmt"

func main() {
    fmt.Println("Line1\nLine2")       // New line
    fmt.Println("Column1\tColumn2")   // Tab
    fmt.Println("She said: \"Go is fun!\"") // Quotes
    fmt.Println("Unicode: \u03A9")    // Greek Omega symbol
}
```

**Output:**

```
Line1
Line2
Column1    Column2
She said: "Go is fun!"
Unicode: Ω
```

---

# 📌 Comments in Go

Comments are used to **explain code**.
They are ignored by the compiler.

## 1. **Single-line Comments**

Use `//`

```go
// This is a single-line comment
fmt.Println("Hello, Go!") // Prints a message
```

---

## 2. **Multi-line Comments**

Use `/* ... */`

```go
/*
This is a multi-line comment
It can span multiple lines
*/
fmt.Println("Welcome to Go")
```

---

## 3. **Documentation Comments**

Go uses `//` style comments above packages, functions, or structs for **godoc** (documentation generator).

```go
// add returns the sum of two integers.
func add(x int, y int) int {
    return x + y
}
```

---

# ✅ Summary

### **Escape Sequences**

* Help you format strings (`\n`, `\t`, `\"`, etc.)
* Unicode supported with `\u` and `\U`.

### **Comments**

* **Single-line**: `//`
* **Multi-line**: `/* ... */`
* **Doc comments**: Start with `//` before packages, functions, etc.

---

⚡Tip: Always use comments for clarity, but don’t over-comment obvious code.
