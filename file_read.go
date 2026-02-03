package main

import (
  "fmt"
  "io"
  "os"
)

func main() {
  // Open a file for reading
  file, err := os.Open("sample.txt")
  if err != nil {
    fmt.Println("Error opening file:", err)
    return
  }
  defer file.Close()

  // Read entire file content
  content, err := io.ReadAll(file)
  if err != nil {
    fmt.Println("Error reading file:", err)
    return
  }

  fmt.Println("File contents:")
  fmt.Println(string(content))

  // Get file info
  info, err := os.Stat("sample.txt")
  if err != nil {
    fmt.Println("Error getting file info:", err)
    return
  }

  fmt.Println("\nFile info:")
  fmt.Println("Name:", info.Name())
  fmt.Println("Size:", info.Size(), "bytes")
  fmt.Println("Modified:", info.ModTime())
}
