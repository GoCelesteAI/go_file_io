package main

import (
  "bufio"
  "fmt"
  "os"
)

func main() {
  // Open file for reading
  file, err := os.Open("data.txt")
  if err != nil {
    fmt.Println("Error opening file:", err)
    return
  }
  defer file.Close()

  // Create a scanner for line-by-line reading
  scanner := bufio.NewScanner(file)

  fmt.Println("Reading file line by line:")
  fmt.Println("==========================")

  lineNum := 1
  for scanner.Scan() {
    line := scanner.Text()
    fmt.Printf("%d: %s\n", lineNum, line)
    lineNum++
  }

  // Check for scanner errors
  if err := scanner.Err(); err != nil {
    fmt.Println("Error reading file:", err)
    return
  }

  fmt.Println("\nTotal lines:", lineNum-1)

  // Reopen file to count words
  file2, _ := os.Open("data.txt")
  defer file2.Close()

  scanner2 := bufio.NewScanner(file2)
  scanner2.Split(bufio.ScanWords)

  wordCount := 0
  for scanner2.Scan() {
    wordCount++
  }

  fmt.Println("Total words:", wordCount)
}
