
# 📌 Tokens in Go

In Go, tokens are classified into **6 main categories**:

---

## 1. **Keywords**

Reserved words that have special meaning in Go.
👉 You cannot use them as identifiers (variable, function names).

### Example Keywords

```
break    default      func     interface    select
case     defer        go       map          struct
chan     else         goto     package      switch
const    fallthrough  if       range        type
continue for          import   return       var
```

### Example in Code:

```go
package main
import "fmt"

func main() {
    var age int = 21
    if age >= 18 {
        fmt.Println("Adult")
    } else {
        fmt.Println("Minor")
    }
}
```

---

## 2. **Identifiers**

Names given to variables, functions, types, etc.

✅ Rules:

* Must start with a letter or underscore `_`.
* Cannot start with a digit.
* Case-sensitive.
* Cannot be a Go keyword.

```go
var studentName string = "Mahfuz"
func calculateSum(x int, y int) int {
    return x + y
}
```

---

## 3. **Literals**

Fixed values that appear directly in the code.

### Types of Literals:

* **Integer literals** → `10`, `255`, `-42`
* **Floating literals** → `3.14`, `-0.99`
* **Boolean literals** → `true`, `false`
* **String literals** → `"Hello"`, `"GoLang"`
* **Rune literals** → `'A'`, `'Ω'`

```go
var x = 100      // integer literal
var pi = 3.1416  // float literal
var flag = true  // boolean literal
var name = "Go"  // string literal
```

---

## 4. **Operators**

Symbols that perform operations on values.

### Categories:

* **Arithmetic**: `+ - * / %`
* **Relational**: `== != > < >= <=`
* **Logical**: `&& || !`
* **Assignment**: `= += -= *= /= %=`
* **Increment/Decrement**: `++ --`
* **Bitwise**: `& | ^ << >>`

```go
x, y := 10, 5
fmt.Println(x + y) // 15
fmt.Println(x > y) // true
```

---

## 5. **Separators (Delimiters)**

Used to structure the program.

| Separator | Meaning                                       |
| --------- | --------------------------------------------- |
| `;`       | Statement terminator (usually optional in Go) |
| `{ }`     | Block of code                                 |
| `( )`     | Function calls / expressions                  |
| `[ ]`     | Array/Slice index                             |
| `,`       | Separator in function arguments               |
| `.`       | Access struct fields or package functions     |

```go
package main
import "fmt"

func main() {
    arr := [3]int{1, 2, 3}
    fmt.Println(arr[0]) // Access first element
}
```

---

## 6. **Comments**

Ignored by the compiler but used for documentation.

* **Single-line** → `// comment`
* **Multi-line** → `/* comment */`

```go
// Single-line comment
/*
 Multi-line
 comment
*/
```

---

# 📘 Complete Example with All Tokens

```go
package main // keyword + identifier

import "fmt" // keyword + string literal

// add function returns sum of two numbers
func add(a int, b int) int { // identifiers, keywords, operators, separators
    return a + b // return keyword, operator
}

func main() {
    var x, y int = 10, 20 // keyword, identifiers, literals, operator
    result := add(x, y)   // function call
    fmt.Println("Result:", result) // output
}
```

**Breakdown of tokens:**

* Keywords: `package`, `import`, `func`, `var`, `return`
* Identifiers: `main`, `fmt`, `add`, `x`, `y`, `result`
* Literals: `"fmt"`, `10`, `20`, `"Result:"`
* Operators: `=`, `+`, `:=`
* Separators: `(`, `)`, `{`, `}`, `,`, `.`
* Comments: `// add function...`

---

## ✅ Summary

**Tokens in Go** are divided into:

1. **Keywords**
2. **Identifiers**
3. **Literals**
4. **Operators**
5. **Separators**
6. **Comments**

Together, they form the **syntax backbone** of a Go program.
