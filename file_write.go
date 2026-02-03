package main

import (
  "fmt"
  "os"
)

func main() {
  // Create a new file (overwrites if exists)
  file, err := os.Create("output.txt")
  if err != nil {
    fmt.Println("Error creating file:", err)
    return
  }
  defer file.Close()

  // Write a string to the file
  _, err = file.WriteString("Hello, File I/O!\n")
  if err != nil {
    fmt.Println("Error writing to file:", err)
    return
  }

  // Write more content
  lines := []string{
    "Line 1: Go makes file handling simple.\n",
    "Line 2: Always close your files.\n",
    "Line 3: Use defer for cleanup.\n",
  }

  for _, line := range lines {
    _, err = file.WriteString(line)
    if err != nil {
      fmt.Println("Error writing line:", err)
      return
    }
  }

  fmt.Println("Successfully wrote to output.txt")

  // Read and display what we wrote
  content, _ := os.ReadFile("output.txt")
  fmt.Println("\nFile contents:")
  fmt.Println(string(content))
}
