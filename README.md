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

# 2. Configuration Management

Configuration Management is the process of managing application settings outside the source code. Instead of hardcoding values such as server ports, database credentials, API keys, and environment settings, we store them in configuration files or environment variables.

This makes applications more secure, maintainable, and easier to deploy across different environments such as Development, Testing, and Production.

## Why Do We Need Configuration Management?

Without configuration management, developers often hardcode values directly into the application:

```go
const PORT = "8080"
const DB_HOST = "localhost"
const DB_USER = "postgres"
const DB_PASSWORD = "password"
```

This approach has several drawbacks:

- Sensitive information is exposed in source code.
- Changing settings requires modifying and redeploying code.
- Different environments require different configurations.
- Configuration becomes difficult to manage as the project grows.

Configuration management solves these problems by separating configuration from application logic.

## Method 1: Environment Variables

Environment variables store configuration values outside the application.

### Advantages

- Good for secrets and sensitive information.
- Supported by all operating systems.
- Easy to override in deployment environments.
- Widely used in cloud-native applications.

### Disadvantages

- Difficult to manage when the number of configuration values increases.
- Does not support hierarchical structures naturally.
- Configuration can become hard to read.

---

## Install godotenv

To load environment variables from a `.env` file:

```bash
go get github.com/joho/godotenv
```

---

## Create a `.env` File

```env
APP_ENV=development
PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=students_db
```

---

### Configuration Loader

Create:

```bash
config/
└── env.go
```

### config/env.go

```go
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()

	if err != nil {
		log.Println(".env file not found")
	}
}

func GetEnv(key string) string {
	return os.Getenv(key)
}
```

### Using Environment Variables

```go
package main

import (
	"fmt"

	"student-api/config"
)

func main() {
	config.LoadEnv()

	port := config.GetEnv("PORT")

	fmt.Println("Server running on port:", port)
}
```

## Method 2: File-Based Configuration

As applications grow, it becomes difficult to manage everything using environment variables.

A configuration file allows related settings to be grouped together in a structured format.

Common formats include:

| Format | Extension    |
| ------ | ------------ |
| YAML   | .yaml / .yml |
| JSON   | .json        |
| TOML   | .toml        |
| XML    | .xml         |

YAML is commonly used in Go projects because it is clean, readable, and supports nested structures.

---

### Why YAML?

Instead of:

```env
SERVER_ADDRESS=localhost:8080
SERVER_TIMEOUT=4s
SERVER_IDLE_TIMEOUT=60s

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=students_db
```

We can organize the same configuration as:

```yaml
http_server:
  address: localhost:8080
  timeout: 4s
  idle_timeout: 60s

database:
  host: localhost
  port: 5432
  user: postgres
  password: password
  dbname: students_db
```

Benefits:

- Better organization
- Easier maintenance
- Human-readable format
- Supports nested configuration
- Cleaner structure for large applications

### Install YAML Package

```bash
go get gopkg.in/yaml.v3
```

### Project Structure

```bash
.
├── cmd
│   └── students-api
│       └── main.go
│
├── config
│   ├── env.go
│   ├── config.go
│   └── local.yaml
│
├── .env
├── go.mod
└── .gitignore
```

### Create Configuration File

### config/local.yaml

```yaml
env: "local"

http_server:
  address: localhost:8080
  timeout: 4s
  idle_timeout: 60s

storage_path: "storage/storage.db"

# or

database:
  host: localhost
  port: 5432
  user: postgres
  password: password
  dbname: students_db
```

### Create Configuration Structure

### config/config.go

```go
package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Env string `yaml:"env"`

	HTTPServer struct {
		Address     string `yaml:"address"`
		Timeout     string `yaml:"timeout"`
		IdleTimeout string `yaml:"idle_timeout"`
	} `yaml:"http_server"`
}

func MustLoad(path string) *Config {
	data, err := os.ReadFile(path)

	if err != nil {
		panic(err)
	}

	var cfg Config

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}

	return &cfg
}
```

### Loading Configuration

```go
package main

import (
	"fmt"

	"student-api/config"
)

func main() {
	cfg := config.MustLoad("config/local.yaml")

	fmt.Println("Environment:", cfg.Env)
	fmt.Println("Server Address:", cfg.HTTPServer.Address)
}
```

## Combining Both Methods (Recommended)

In real-world applications, both methods are typically used together.

### Step 1: Store the configuration file path in `.env`

```env
CONFIG_PATH=config/local.yaml
```

### Step 2: Load the environment variable

```go
config.LoadEnv()

configPath := config.GetEnv("CONFIG_PATH")
```

### Step 3: Load the YAML file

```go
cfg := config.MustLoad(configPath)
```

This approach provides flexibility because different environments can load different configuration files.

For example:

Development

```env
CONFIG_PATH=config/local.yaml
```

Testing

```env
CONFIG_PATH=config/test.yaml
```

Production

```env
CONFIG_PATH=config/prod.yaml
```

### Best Practices

- Never hardcode secrets.
- Store API keys, tokens, and passwords in environment variables.
- Use YAML files for application settings.
- Keep environment-specific configurations separate.
- Add `.env` to `.gitignore`.
- Validate configuration during application startup.

---
