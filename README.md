# API_GO

This is a Go API project for managing scholar control of students. It uses MySQL as the database and Docker to facilitate deployment.

## Table of Contents

- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Usage](#usage)

## Features

- Student management with CRUD (Create, Read, Update, Delete) operations.
- Web interface to interact with the API.
- Easy deployment using Docker of the.

## Requirements
- [Docker](https://www.docker.com/get-started)
- [Go](https://golang.org/doc/install)

## Installation

To install and run this project, follow these steps:

1. Clone the repository: 
git clone https://github.com/SebasChips/API_GO.git
2. Navigate to the project directory: 
cd API_GO
3. Build the Docker image for MySQL: 
docker build -t mi_mysql:latest .
4. Run the MySQL container:
docker run --name contenedor_mysql -d -p 3306:3306 mi_mysql:latest
5. Navigate to the cmd directory:
cd cmd
6. Run the Go application:
go run main.go

## Usage

Once the application is running, you can access the web interface in your browser at the following address:

http://localhost:8000/alumnos.html
