# Golang Basics Learning Repository

A comprehensive, beginner-friendly Go programming language learning resource covering fundamental concepts through advanced concurrency patterns.

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Version](https://img.shields.io/badge/Go-1.x-blue.svg)](https://golang.org/)
![Language](https://img.shields.io/badge/Language-Go-00ADD8.svg)

## 📚 About This Repository

This repository is a structured learning guide for **Go (Golang)**, an open-source programming language created by Google. Go is used for a variety of purposes, including:
- 🌐 **Web Development** - Building REST APIs and web services
- ☁️ **Cloud Infrastructure** - Kubernetes, Docker, and cloud-native applications
- 🖥️ **Server Management** - High-performance backend services
- ⚡ **Systems Programming** - Concurrent and parallel computing

Whether you're a complete beginner or have programming experience in other languages, this repository provides practical examples and explanations for every Go concept.

---

## 📖 Curriculum Overview

The repository is organized into **24 progressive modules**, each focusing on a specific Go concept:

### Foundations
| Module | Topic | Description |
|--------|-------|-------------|
| **01** | [Hello World](1_hello_world) | Getting started with your first Go program |
| **02** | [Simple Values](2_simple_values) | Working with integers, floats, strings, and booleans |
| **03** | [Variables](3_variables) | Variable declaration, initialization, and scope |
| **04** | [Constants](4_constants) | Working with immutable constant values |

### Control Flow
| Module | Topic | Description |
|--------|-------|-------------|
| **05** | [For Loop](5_for_loop) | Understanding Go's iteration construct |
| **06** | [If/Else](6_if_else) | Conditional statements and branching |
| **07** | [Switch](7_switch) | Switch statements for multi-way branching |

### Data Structures
| Module | Topic | Description |
|--------|-------|-------------|
| **08** | [Arrays](8_arrays) | Fixed-size collections of elements |
| **09** | [Slices](9_slices) | Dynamic, flexible arrays in Go |
| **10** | [Maps](10_maps) | Key-value pairs and hash tables |
| **11** | [Range](11_Range) | Iterating over collections efficiently |

### Functions & Advanced Concepts
| Module | Topic | Description |
|--------|-------|-------------|
| **12** | [Functions](12_functions) | Defining and calling functions, multiple returns |
| **13** | [Variadic Functions](13_variadic_function) | Functions with variable number of arguments |
| **14** | [Closures](14_closures) | Functions that capture variables from outer scope |
| **15** | [Pointers](15_pointers) | Working with memory addresses and references |

### Object-Oriented Programming
| Module | Topic | Description |
|--------|-------|-------------|
| **16** | [Structs](16_structs) | Grouping related data into composite types |
| **17** | [Interfaces](17_interface) | Defining contracts and polymorphism |
| **18** | [Enums](18_enums) | Implementing enumeration patterns |

### Advanced Features
| Module | Topic | Description |
|--------|-------|-------------|
| **19** | [Generics](19_generics) | Type-safe, reusable code with generics (Go 1.18+) |
| **20** | [Goroutines](20_goroutines) | Lightweight concurrency primitives |
| **21** | [Channels](21_channels) | Communication between goroutines |
| **22** | [Mutex](22_mutex) | Synchronization and mutual exclusion |
| **23** | [File Operations](23_files) | Reading and writing files |
| **24** | [Packages](24_packages) | Organizing code into reusable modules |

---

## 🚀 Getting Started

### Prerequisites
- Go 1.x installed on your system ([Download Go](https://golang.org/dl/))
- Basic understanding of programming concepts (helpful but not required)
- A text editor or IDE (VS Code, GoLand, etc.)

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/isinghabhishek/golang_basic.git
   cd golang_basic
   ```

2. **Verify Go installation:**
   ```bash
   go version
   ```

3. **Navigate to any module and run examples:**
   ```bash
   cd 1_hello_world
   go run main.go
   ```

---

## 💡 How to Use This Repository

### For Beginners
1. Start with module **01 - Hello World** and progress sequentially
2. Read the comments in each code file carefully
3. Modify the examples and experiment with different inputs
4. Try to solve the challenges before looking at solutions

### For Experienced Programmers
- Jump directly to topics you want to learn
- Compare Go's approach with languages you know
- Focus on modules like Goroutines, Channels, and Generics
- Use this as a quick reference guide

### Best Practices
- ✅ Read and understand each example thoroughly
- ✅ Run every example and observe the output
- ✅ Modify code to see how it behaves
- ✅ Practice writing your own variations
- ✅ Move to the next module only after understanding the current one

---

## 📁 Repository Structure

```
golang_basic/
├── 1_hello_world/          # Your first Go program
├── 2_simple_values/        # Basic data types
├── 3_variables/            # Variable declarations
├── ...
├── 24_packages/            # Organizing code into packages
├── go.mod                  # Go module definition
├── LICENSE                 # MIT License
├── CODE_OF_CONDUCT.md      # Community guidelines
└── README.md              # This file
```

Each module directory contains:
- `main.go` - Example code demonstrating the concept
- Comments explaining key points and concepts
- Practical, runnable examples

---

## 🎯 Key Learning Outcomes

By completing this repository, you will:

✓ Understand Go's syntax and core language features
✓ Master data structures (arrays, slices, maps)
✓ Write functions and understand scope
✓ Implement object-oriented patterns with structs and interfaces
✓ Build concurrent applications with goroutines and channels
✓ Work with files and packages
✓ Apply modern Go features like generics
✓ Be ready to build real-world applications

---

## 🤝 Contributing

Contributions are welcome! If you find bugs, have suggestions, or want to add improvements:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/improvement`)
3. Commit your changes (`git commit -m 'Add improvement'`)
4. Push to the branch (`git push origin feature/improvement`)
5. Open a Pull Request

Please follow the [Code of Conduct](CODE_OF_CONDUCT.md) and ensure your code is well-commented and follows Go conventions.

---

## 📚 Additional Resources

- [Official Go Documentation](https://golang.org/doc/)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go Blog](https://blog.golang.org/)
- [Go Playground](https://play.golang.org/) - Try Go in your browser

---

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

## 👤 Author

**Abhishek Singh** ([@isinghabhishek](https://github.com/isinghabhishek))

---

## 📞 Support & Questions

If you have questions or need help:
- Open an [Issue](https://github.com/isinghabhishek/golang_basic/issues)
- Check existing discussions
- Review the code comments and examples

---

## ⭐ Show Your Support

If you find this repository helpful, please give it a ⭐ Star! It helps others discover this learning resource.

Happy Coding! 🎉
