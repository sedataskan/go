# Cheatsheet for Go Development

-> to run the Go program, use:

```bash
go run main.go
```

-> to build the mod file which contains dependencies, use:

```bash
go mod init <module-name>
```

-> to add a dependency, use:

```bash
go get <package-name>
```

-> to tidy up the module file and remove unused dependencies, use:

```bash
go mod tidy
```

- on Go, maps are unordered collections of key-value pairs. When you print a map, the order of the keys is not guaranteed to be the same every time you run the program.
