# 🧰 go-toolbox

**go-toolbox** is a personal Go library that collects utility types, helpers, and patterns I’ve built to make my Go development faster, cleaner, and more consistent across projects.

This project serves as a central place for reusable Go code — small, reliable, and well-tested tools that I can import anywhere.

---

## 🚀 Features

- **Reusable utilities** for everyday Go development  
- **Strongly typed helpers** (like version parsing and comparison)  
- **Well-tested and documented** code for easy maintenance  
- **Modular structure**, so you can import only what you need  

---

## 📦 Installation

```bash
go get github.com/Diaszano/go-toolbox@latest
````

Then import it in your code:

```go
import "github.com/Diaszano/go-toolbox/types/version"
```

---

## 🧩 Example

Here’s an example using the `version` package:

```go
package main

import (
	"fmt"
	"github.com/Diaszano/go-toolbox/types/version"
)

func main() {
	v, err := version.TryParse("1.2.3")
	if err != nil {
		panic(err)
	}

	fmt.Println("Version:", v.String())   // Output: 1.2.3
	fmt.Println("Prefixed:", v.Version()) // Output: v1.2.3

	// Compare versions
	v2, _ := version.Parse("1.3.0")
	fmt.Println("Is v2 newer?", v2.Compare(v) > 0) // true
}
```

---

## 🧠 Project Structure

```
go-toolbox/
├── go.mod
├── go.sum
├── README.md
└── types/
    └── version/
        ├── version.go
        └── version_test.go
```

Each folder under `types` or `utils` represents a self-contained package that can be imported independently.

---

## 🧪 Testing

Run all tests with:

```bash
go test ./...
```

You can also use `make test` if you have a Makefile.

---

## 🧹 Linting

This project uses [`golangci-lint`](https://golangci-lint.run/) for static analysis.
To run it locally:

```bash
golangci-lint run
```

---

## 📄 License

This project is licensed under the **MIT License** — feel free to use and modify it.
