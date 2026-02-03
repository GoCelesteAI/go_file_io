package main

import (
  "fmt"
  "os"
  "time"
)

func main() {
  // Open file for appending (create if not exists)
  file, err := os.OpenFile("log.txt",
    os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  if err != nil {
    fmt.Println("Error opening file:", err)
    return
  }
  defer file.Close()

  // Write log entries
  timestamp := time.Now().Format("2006-01-02 15:04:05")
  entry := fmt.Sprintf("[%s] Application started\n", timestamp)

  _, err = file.WriteString(entry)
  if err != nil {
    fmt.Println("Error writing to file:", err)
    return
  }

  fmt.Println("Appended log entry")

  // Add more entries
  events := []string{"User logged in", "Data processed", "Task completed"}

  for _, event := range events {
    timestamp = time.Now().Format("2006-01-02 15:04:05")
    entry = fmt.Sprintf("[%s] %s\n", timestamp, event)
    file.WriteString(entry)
    fmt.Println("Logged:", event)
  }

  // Read and display the log
  content, _ := os.ReadFile("log.txt")
  fmt.Println("\nLog file contents:")
  fmt.Println(string(content))
}
