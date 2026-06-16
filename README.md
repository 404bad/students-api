# Student API - Go Project Setup

## 1. Check Go Version

Before starting, verify that Go is installed on your system:

```bash
go version
```

This command displays the currently installed Go version.

---

## 2. Initialize the Go Module

Create a new Go module:

```bash
go mod init github.com/your-username/student-api
```

This generates a `go.mod` file containing the module name and Go version.

Example:

```go
module github.com/your-username/student-api

go 1.24.0
```

The module is important because it acts as the unique identifier for your project and is later used when importing packages within the application.

---

## 3. Typical Go Project Structure

A common folder structure for a Go application is:

```bash
.
├── cmd
│   └── students-api
│       └── main.go
├── go.mod
└── README.md
```

### main.go

```go
package main

import "fmt"

func main() {
    fmt.Println("WELCOME TO STUDENT API")
}
```

---

## 4. Running the Application

There are multiple ways to execute a Go application.

### Option 1: Build and Run the Executable

Build the project:

```bash
go build -o student-api cmd/students-api/main.go
```

Run the executable:

```bash
./student-api
```

### Option 2: Run Directly (Recommended During Development)

Instead of building first, you can run the application directly:

```bash
go run cmd/students-api/main.go
```

This compiles and executes the program in a single command.

---

## 5. Initialize Git Repository

Before creating the Git repository, create a `.gitignore` file.

Initialize Git:

```bash
git init
git add . #Add all files
git commit -m "Initial project setup" #Create the first commit:
git remote add origin <repository-url> #Add the remote repository:
git branch -M main
git push -u origin main #Push the code:

gh repo create # create repo and push to github using gthub cli
#  then follow on
```

---
