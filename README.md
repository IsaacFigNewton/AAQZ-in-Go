# AAQZ-in-Go

## Getting Started
1. [Install Go](https://go.dev/dl/)

## Adding a module
1. Create a new directory for the new module's packages
2. In your IDE terminal, run
  1. `cd <module-name>`
  2. `go mod init <module-name>/<package-name>`
5. Make a new go file, such as `<package-name>.go`
   - Add your code to the file
   - Make sure that `package <package-name>` is written at the very top
6. In your IDE terminal, run
  1. `cd ..`
  2. `cd <main-module-name>`
9. Modify the main package (ex: `main.go`)
    - add `<module-name>/<package-name>` to the list of imports
10. In your IDE terminal, run
  1. `go mod edit -replace <module-name>/<package-name>=../<module-name>`
  2. `go mod tidy`
  3. `go run .`

## Helpful Resources
- [Tutorials](https://go.dev/doc/tutorial/)
- [Go Docs](https://go.dev/ref/spec)
