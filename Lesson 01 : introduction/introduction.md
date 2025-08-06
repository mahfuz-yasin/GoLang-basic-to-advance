
# 📌 What is Go (Golang)?

**Go**, also called **Golang**, is an **open-source
programming language** created by **Google** in **2
and released publicly in **2009**.

It was designed by **Robert Griesemer, Rob Pike, an
Thompson** (three legendary software engineers from
Google).

---

## 🔑 Key Features of Go

1. **Compiled Language**

   * Go is compiled into machine code, so it runs v
fast (comparable to C/C++).

2. **Statically Typed**

   * Variables must have types, but Go also support
inference (`:=`).

3. **Simple Syntax**

   * Easier to learn compared to C++ or Java, yet v
powerful.

4. **Built-in Concurrency**

   * Go has **goroutines** and **channels** for run
multiple tasks at once.
   * This makes it excellent for network servers, c
apps, and distributed systems.

5. **Garbage Collection**

   * Memory management is automatic, so you don’t n
manually free memory.

6. **Cross-Platform**

   * Write once, run anywhere (Linux, Windows, macO

7. **Standard Library**

   * Provides a rich set of packages for web, 
cryptography, file handling, concurrency, etc.

---

## 🖥️ Example: A Simple Go Program

```go
package main
import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

🔹 **Explanation:**

* `package main` → defines the main package (entry 
* `import "fmt"` → imports the format package for i
output.
* `func main()` → execution starts here.
* `fmt.Println()` → prints text to the screen.

---

## ⚡ Why Use Go?

* **Speed**: Runs close to C/C++ performance.
* **Concurrency**: Best choice for modern applicati
like servers, APIs, cloud services.
* **Simplicity**: Clean, readable, and beginner-fri
* **Scalability**: Used in high-traffic systems lik
YouTube, Dropbox, Docker, Kubernetes.
* **Community & Support**: Backed by Google, widely
in industry.

---

## 📌 Who Uses Go?

* **Google** (for internal services)
* **Uber** (for real-time services)
* **Netflix** (microservices)
* **Dropbox** (backend systems)
* **Docker & Kubernetes** (written in Go)


# ✅ **Step 02: Environment Setup**

To start programming in Go, you need to set up a proper development environment. This includes an editor, compiler, and essential extensions.

---

## 1. 🖊️ IDE / Text Editor

Use any editor you’re comfortable with, but the most recommended one is:

### ➤ **Visual Studio Code (VSCode)**

* Lightweight and fast
* Rich extension support
* Cross-platform (Windows, Linux, macOS)

✅ **Download**: [https://code.visualstudio.com](https://code.visualstudio.com)

---

## 2. ⚙️ Go Compiler

### ➤ **Go Toolchain / Compiler**

* The Go compiler translates your source code (`.go` files) into **machine code**.
* It includes:

  * `go run`: compile and run the program
  * `go build`: compile into an executable
  * `go fmt`: format your code
  * `go mod`: manage dependencies

✅ **Download Go**: [https://go.dev/dl](https://go.dev/dl)

After installing, check version:

```bash
go version
```

---

## 3. 🧩 Go Extension for VSCode

### ➤ **Extension Name:** `Go`

* Developed by **Go Team at Google**
* Adds features like:

  * Syntax highlighting
  * Auto-completion
  * Debugging
  * Linting and formatting
  * IntelliSense
  * Code navigation and documentation

✅ Install it from VSCode Marketplace:
**Extensions → Search: “Go” → Install**

---

## ✅ Optional (Recommended):

| Tool    | Purpose                               |
| ------- | ------------------------------------- |
| `gofmt` | Format code automatically             |
| `gopls` | Go Language Server (for IntelliSense) |
| `Delve` | Debugging tool for Go                 |
| `Git`   | Version control                       |

---

## 🔍 Terminal Setup (Optional)

If you use **Linux/Mac**, it’s recommended to use:

```bash
sudo apt install golang-go
```

For **Windows**, Go installer handles everything.

---

## 📌 Summary

| Component      | Tool                     |
| -------------- | ------------------------ |
| Text Editor    | VSCode                   |
| Compiler       | Go (from go.dev)         |
| Extension      | Go for VSCode            |
| Optional Tools | gofmt, gopls, Delve, Git |
