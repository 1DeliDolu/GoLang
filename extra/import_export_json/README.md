# JSON Import/Export Example

This project demonstrates how to import and export JSON data in Go.

## Running the Application

In order to run the application correctly, you need to include all Go files in the directory.

### Using Go directly:

```bash
go run *.go
```

Or specifically:

```bash
go run main.go json-export.go json-import.go
```

### Using the convenience script:

```bash
chmod +x run.sh
./run.sh
```

## Error Explanation

If you run just `go run main.go`, you'll get errors like:

```
# command-line-arguments
./main.go:16:2: undefined: ExportJSON
./main.go:19:2: undefined: ImportJSON
```

This happens because the functions are defined in separate files, and you need to include all files when running the application.
