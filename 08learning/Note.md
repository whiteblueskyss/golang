# Go Learning Notes

## Compilation Optimization

- Go only compiles packages that have changed.
- Unchanged files won't be recompiled each time.
- This speeds up compile time.

---

## Go Commands

### `go mod tidy`
- Cleans up the `go.mod` file by removing unused dependencies and adding missing ones.

### `go get`
```bash
go get github.com/fatih/color
```
- Downloads and installs external packages.
- Example: Installing the `color` package for terminal color output.
