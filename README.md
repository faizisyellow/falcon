![falcon logo](/assets/images/falcon-logo.png)


![GitHub tag (latest by date)](https://img.shields.io/github/v/tag/faizisyellow/falcon)

Falcon is a CLI tool to generate GO REST APIs
the generated project already comes with Authentication,
Swagger documentation and following Handler-Service-Repository Pattern.

## Features

- Create New Project
- Options to choose router and Database (Currently only Chi Router and Mysql Database is available)



## Prerequisite

- [Go v1.24.3](https://go.dev/doc/install) 
- [Migrate](https://github.com/golang-migrate/migrate) 
- [Swag](https://github.com/swaggo/swag) 
- [Air](https://github.com/air-verse/air)


## Installation

Falcon requires [Go](https://go.dev/doc/install) v1.24+ to run.

Download it in:
- [Binaries](https://github.com/faizisyellow/falcon/releases) are available for Linux, macOS, and Windows
    ```
    // linux
    tar -xvzf  falcon_Linux_x86_64.tar.gz 
    sudo mv falcon /usr/local/bin/
    ```

## Usage

Let's get started shall we?

1. Create Go Module in your project directory.
2. Create new project.
   ```
   falcon init 
   ```
3. Run ``` Go mod tidy ```.
4. Rename **.env.sample** to **.env** and fill in the environment variable.
5. Run migration ```make migrate-up ```. Make sure you already create database.
6. Run swag doc generator ```make gen-docs```
7. Run air ```air``` to started the server. Your swagger documentation now in <http://localhost:8080/v1/swagger/index.html>
