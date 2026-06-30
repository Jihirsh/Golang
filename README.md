# Go Learning Journey

This repository documents my journey of learning Go by **writing code myself**, debugging it, and understanding the language through official documentation instead of relying on AI-generated solutions.

## Learning Philosophy

- Think first, Google second.
- Search for concepts, not complete solutions.
- Read official documentation before blogs.
- Write code by hand.
- Debug instead of rewriting.
- Experiment with edge cases.
- Understand *why* code works, not just *how*.

---

# Learning Roadmap

## Phase 1 - Go Fundamentals

### Program Execution

Learn:

- What is `package main`?
- What is `func main()`?
- How does `go run .` work?
- How Go organizes packages.

Resources:

- https://go.dev/doc/code
- https://go.dev/doc/tutorial/getting-started
- https://gobyexample.com/hello-world

---

### Functions

Learn:

- Function syntax
- Parameters
- Return values
- Multiple return values
- Variadic functions

Resources:

- https://go.dev/tour/basics/4
- https://go.dev/tour/basics/8
- https://gobyexample.com/functions
- https://gobyexample.com/multiple-return-values
- https://gobyexample.com/variadic-functions

---

### Error Handling

Learn:

- What is the `error` type?
- Why Go returns errors instead of exceptions.
- `if err != nil`
- Creating custom errors.

Resources:

- https://go.dev/blog/errors-are-values
- https://gobyexample.com/errors
- https://gobyexample.com/custom-errors

---

### Input / Output

Learn:

- `fmt.Scan`
- `fmt.Scanln`
- `fmt.Scanf`
- `fmt.Print`
- `fmt.Println`
- `fmt.Printf`

Resources:

- https://pkg.go.dev/fmt

Experiment with:

- Invalid input
- Empty input
- Multiple values
- Return values (`n`, `err`)

---

### Buffered Input (Later)

Learn:

- `bufio.Reader`
- Reading full lines
- Why `bufio` is preferred over `fmt.Scan` in many cases.

Resources:

- https://pkg.go.dev/bufio
- https://gobyexample.com/line-filters

---

### Go Tour

Resource:

- https://go.dev/tour

---

### Effective Go (After learning basics)

Read after becoming comfortable with Go syntax.

Resource:

- https://go.dev/doc/effective_go

---

# Projects

- [x] Number Guessing Game
- [ ] Calculator CLI
- [ ] Todo List CLI
- [ ] File Reader
- [ ] JSON Parser
- [ ] REST API
- [ ] URL Shortener
- [ ] Concurrent Downloader

---

# Learning Process

For every new topic:

1. Read official documentation.
2. Read Go by Example.
3. Build a small experiment.
4. Try edge cases.
5. Debug unexpected behavior.
6. Only then build a project.

---

# Debugging Checklist

When something doesn't work, ask:

- Am I running the correct file?
- Is this function actually being called?
- What are the variable values?
- What does the function return?
- What assumptions am I making?
- Can I reproduce the bug with a smaller program?

---

# Useful Commands

```bash
go run .
go build
go fmt
go test
go mod init <module-name>
go mod tidy
```

---

# Notes

Every project should teach at least one new concept.

The goal is **not** to finish projects quickly.

The goal is to understand how Go works.
