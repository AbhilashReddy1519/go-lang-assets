# Complete Go CLI Command Guide

> Learn every important Go command, when to use it, why it exists, and common examples.

---

# Go Command Overview

Run:

```bash
go help
```

Output:

```text
bug
build
clean
doc
env
fix
fmt
generate
get
install
list
mod
run
test
tool
version
vet
work
```

---

# 1. go build

Builds source code into an executable binary.

## Why?

Compile code without running it.

## Example

```bash
go build
```

Creates:

```text
./myproject
```

Build specific file:

```bash
go build main.go
```

Build for Linux:

```bash
GOOS=linux GOARCH=amd64 go build
```

## Common Usage

```bash
go build ./...
```

Build every package.

---

# 2. go run

Compiles and immediately executes code.

## Why?

Quick testing during development.

```bash
go run main.go
```

Multiple files:

```bash
go run *.go
```

Run package:

```bash
go run .
```

---

# 3. go install

Build and install binary.

```bash
go install
```

Installs to:

```text
$GOPATH/bin
```

Install external tool:

```bash
go install golang.org/x/tools/cmd/stringer@latest
```

---

# 4. go get

Download or upgrade dependencies.

Add dependency:

```bash
go get github.com/gin-gonic/gin
```

Upgrade:

```bash
go get -u
```

Specific version:

```bash
go get github.com/gin-gonic/gin@v1.10.0
```

---

# 5. go test

Run tests.

```bash
go test
```

All packages:

```bash
go test ./...
```

Verbose:

```bash
go test -v
```

Coverage:

```bash
go test -cover
```

Coverage report:

```bash
go test -coverprofile=cover.out
```

Benchmark:

```bash
go test -bench=.
```

---

# 6. go fmt

Formats source code.

```bash
go fmt
```

Entire project:

```bash
go fmt ./...
```

## Why?

Enforces standard Go style.

---

# 7. go vet

Finds suspicious code.

```bash
go vet
```

Checks:

* Wrong printf arguments
* Unreachable code
* Misused locks
* Struct tag errors

---

# 8. go clean

Removes build artifacts.

```bash
go clean
```

Remove module cache:

```bash
go clean -modcache
```

Remove test cache:

```bash
go clean -testcache
```

---

# 9. go version

Shows installed version.

```bash
go version
```

Example:

```text
go version go1.25.0 linux/amd64
```

---

# 10. go env

Displays Go environment variables.

```bash
go env
```

Specific variable:

```bash
go env GOPATH
```

Important variables:

```text
GOROOT
GOPATH
GOOS
GOARCH
GOMODCACHE
GOCACHE
```

---

# 11. go list

Lists packages/modules.

```bash
go list
```

All packages:

```bash
go list ./...
```

All modules:

```bash
go list -m all
```

---

# 12. go doc

Displays documentation.

```bash
go doc fmt
```

Function docs:

```bash
go doc fmt.Println
```

---

# 13. go fix

Automatically updates old Go code.

```bash
go fix
```

Useful after language upgrades.

---

# 14. go generate

Runs code generators.

Source:

```go
//go:generate stringer -type=Status
```

Run:

```bash
go generate
```

---

# 15. go tool

Accesses internal Go tools.

List:

```bash
go tool
```

Examples:

```bash
go tool compile
go tool asm
go tool pprof
go tool trace
```

---

# 16. go bug

Creates bug reports.

```bash
go bug
```

Collects system information.

---

# 17. go work

Workspace management.

Create:

```bash
go work init
```

Add modules:

```bash
go work use ./service-a
go work use ./service-b
```

File:

```text
go.work
```

Useful for monorepos.

---

# 18. go mod

Module management.

---

## go mod init

```bash
go mod init project
```

Create module.

---

## go mod tidy

```bash
go mod tidy
```

Clean dependencies.

---

## go mod download

```bash
go mod download
```

Download dependencies.

---

## go mod vendor

```bash
go mod vendor
```

Create vendor directory.

---

## go mod verify

```bash
go mod verify
```

Verify integrity.

---

## go mod graph

```bash
go mod graph
```

Dependency graph.

---

## go mod why

```bash
go mod why package
```

Why dependency exists.

---

## go mod edit

```bash
go mod edit
```

Modify go.mod programmatically.

---

# Important Development Workflow

## New Project

```bash
go mod init myapp
go run .
```

---

## Add Dependency

```bash
go get github.com/gin-gonic/gin
```

---

## Clean Dependencies

```bash
go mod tidy
```

---

## Format Code

```bash
go fmt ./...
```

---

## Static Analysis

```bash
go vet ./...
```

---

## Run Tests

```bash
go test ./...
```

---

## Build Binary

```bash
go build
```

---

## Install Binary

```bash
go install
```

---

# Most Used Commands in Real Projects

Daily:

```bash
go run .
go test ./...
go build
go fmt ./...
go vet ./...
go mod tidy
go get
```

Weekly:

```bash
go list
go doc
go env
go mod graph
go mod why
```

Rare:

```bash
go generate
go tool
go fix
go bug
go work
```

---

# Go Command Cheat Sheet

| Command     | Purpose              |
| ----------- | -------------------- |
| go build    | Compile              |
| go run      | Compile + Execute    |
| go install  | Install Binary       |
| go get      | Manage Dependencies  |
| go test     | Testing              |
| go fmt      | Formatting           |
| go vet      | Static Analysis      |
| go clean    | Cleanup              |
| go env      | Environment          |
| go list     | Package Listing      |
| go doc      | Documentation        |
| go generate | Code Generation      |
| go tool     | Internal Tools       |
| go work     | Workspace Management |
| go mod      | Module Management    |
| go version  | Version Info         |
| go bug      | Bug Reports          |
| go fix      | Upgrade Old Code     |

```

This covers about **95% of commands used by Go developers**. The remaining commands are mostly low-level tooling (`compile`, `link`, `asm`, `cgo`, `nm`, `objdump`, `pprof`, `trace`, etc.) that are accessed through `go tool`.
```
