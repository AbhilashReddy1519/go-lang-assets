# Local Go development — step-by-step

This guide lists the typical local development steps for Go projects, why each step matters, and the exact commands to run. Use it as a checklist when onboarding onto a Go repository.

**Prerequisites**
- OS: macOS, Linux, or Windows
- Internet access to download Go and modules
- Basic familiarity with command line

**1) Install Go (toolchain)
- Why: The `go` tool compiles, tests, runs, and manages modules. You need a matching toolchain for the project's Go version.
- How to check if already installed:

```bash
go version
```

- Install (official):

macOS / Linux:
```bash
# follow https://go.dev/doc/install — example for Linux (tarball)
wget https://go.dev/dl/go1.20.9.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.20.9.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

Windows (recommended): download the MSI from https://go.dev/dl/ and run the installer.

- Importance: the compiler, `gofmt`, `vet`, and module commands come from this installation.

**2) Use Go Modules (project-scoped dependencies)
- Why: Modules isolate dependencies and define reproducible builds using `go.mod` and `go.sum`.
- Check project root contains `go.mod`.

```bash
ls -la go.mod
cat go.mod
```

- If missing, initialize (only for new projects):

```bash
go mod init github.com/your/repo
```

- Important: Never modify `go.sum` manually; use `go` commands to update it.

**3) Fetch dependencies & verify module graph
- Why: Ensures your local cache has the dependencies and verifies checksums.

```bash
go mod download      # download modules to module cache
go mod tidy          # add/remove modules to match imports
go list -m all       # list all modules in the build list
```

**4) Build the project
- Why: Confirms code compiles before running or testing.

```bash
go build ./...       # build all packages in module
```

- Common flags:
  - `-v` verbose
  - `-o ./bin/app` to write binary to a path

**5) Run the program (development run)
- Why: Quick development iteration.

```bash
go run ./cmd/myapp   # or `go run .` at package main root
```

**6) Run tests
- Why: Verify behavior, prevent regressions.

```bash
go test ./...                # run package tests recursively
go test -v ./pkg/mypkg       # verbose test output for a package
go test -run TestName ./...  # run tests matching pattern
```

- Coverage and race detector:

```bash
go test -coverprofile=cover.out ./...
go test -race ./...           # detect data races (useful for concurrent code)
```

**7) Linting, formatting, and vet
- Why: Ensure consistent style and catch common mistakes.

```bash
gofmt -w .            # format files in-place
go vet ./...           # static checks
# optional: golangci-lint (fast, multi-linter)
golangci-lint run
```

**8) Static analysis and security
- Why: Find potential bugs and security issues before shipping.

```bash
go vet ./...
gosec ./...            # security scanner (separate install)
```

**9) Debugging locally
- Why: Inspect runtime state, step through code.
- Use `delve` (install `dlv`):

```bash
dlv debug ./cmd/myapp -- -arg1=value
```

**10) Development workflows (watch/build loops)
- Why: Faster feedback by rebuilding or running tests on file changes.
- Tools: `air`, `reflex`, or `CompileDaemon`.

Example with `reflex`:
```bash
reflex -r '\.go$' -- sh -c 'clear; go test ./...'
```

**11) Working with multiple modules / monorepos
- Why: Some repos have multiple `go.mod` files; run commands per module root.

```bash
cd ./module-a
go test ./...
```

**12) Environment variables commonly used
- `GOOS` and `GOARCH` — cross-compile targets.
- `GOMAXPROCS` — max OS threads for Go scheduler.
- `GOPROXY` — module proxy (default `https://proxy.golang.org`).
- `GOSUMDB` — checksum database.

Examples:
```bash
GOOS=linux GOARCH=amd64 go build -o bin/app-linux-amd64
GOMAXPROCS=4 go run .
```

**13) CI integration (quick checklist)
- `go test ./...` with `-race` and `-cover` where appropriate.
- Run `gofmt -l .` and `golangci-lint` to enforce style.
- Cache Go module downloads using module cache in runners (`GOMODCACHE`).

**14) Working with private modules
- Why: Access private repos via modules.
- Options: set `GOPRIVATE` and configure `GONOPROXY`/`GONOSUMDB`, or use a private proxy.

```bash
export GOPRIVATE=github.com/myorg/*
```

**15) Releasing and versioning
- Use semantic versions (tags) for modules: `git tag v1.2.3` and push tags.
- `go install ./cmd/myapp@latest` installs from module's versioned tag when publishing.

**16) Useful commands reference
- Build: `go build ./...`
- Run: `go run ./cmd/myapp`
- Test: `go test ./...`
- Format: `gofmt -w .`
- Vet: `go vet ./...`
- Modules: `go mod download`, `go mod tidy`

**17) Common pitfalls & tips
- Rely on modules, not GOPATH: modern Go uses modules by default.
- Keep `go.mod` tidy; run `go mod tidy` before commits.
- Use `gofmt`/`goimports` to avoid style comments in reviews.
- Prefer small, focused packages — easier testing and reuse.
- For reproducible builds, pin CI to a specific Go version.

**18) Editor & IDE integration
- VS Code: `golang.go` extension (gopls, formatting, linting, test runner).
- GoLand: full IDE with built-in tooling.
- Ensure `gopls` is installed and configured for module mode.

**19) Quick onboarding checklist (one-liner summary)
1. Install Go (match project version). 2. `git clone` repo. 3. `cd repo` and `go mod download`. 4. `gofmt -w .` and `go vet ./...`. 5. `go test ./...` 6. `go run ./cmd/myapp`.

---
If you want, I can add a short `scripts/` folder with convenient wrapper scripts (`make`, `scripts/dev.sh`) or wire an example `launch.json` for VS Code. Would you like that? 
