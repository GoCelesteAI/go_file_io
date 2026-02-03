# Go File I/O

Demo code from **Go Tutorial Lesson 22: File I/O**

Watch the full tutorial: [YouTube Link]

## Topics Covered

1. **Reading Files** - `os.Open`, `io.ReadAll`, `os.Stat`
2. **Writing Files** - `os.Create`, `WriteString`
3. **Line-by-Line Reading** - `bufio.Scanner`, `ScanWords`
4. **Appending to Files** - `os.OpenFile` with flags

## Files

| File | Description |
|------|-------------|
| `file_read.go` | Read entire file, get file info |
| `file_write.go` | Create and write to file |
| `file_scanner.go` | Line-by-line and word-by-word reading |
| `file_append.go` | Append log entries to file |
| `sample.txt` | Sample text file for reading demo |
| `data.txt` | CSV-like data for scanner demo |

## Running the Examples

```bash
# Reading files
go run file_read.go

# Writing files
go run file_write.go

# Line-by-line with Scanner
go run file_scanner.go

# Appending to files
go run file_append.go
# Run again to see appending in action
go run file_append.go
```

## Key Concepts

### Reading Files
```go
file, err := os.Open("file.txt")
if err != nil {
    // handle error
}
defer file.Close()

content, err := io.ReadAll(file)
```

### Writing Files
```go
file, err := os.Create("output.txt")
if err != nil {
    // handle error
}
defer file.Close()

file.WriteString("Hello, File I/O!\n")
```

### Line-by-Line Reading
```go
scanner := bufio.NewScanner(file)
for scanner.Scan() {
    line := scanner.Text()
    fmt.Println(line)
}
```

### Appending to Files
```go
file, err := os.OpenFile("log.txt",
    os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
if err != nil {
    // handle error
}
defer file.Close()

file.WriteString("New log entry\n")
```

## Channel

Subscribe for more Go tutorials: [GoCeleste AI](https://youtube.com/@GoCelesteAI)
