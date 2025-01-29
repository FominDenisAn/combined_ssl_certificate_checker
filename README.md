# SSL Certificate Checker

## Overview

This repository contains both the backend server written in Go and the frontend client written in Vue.js. The project is designed to monitor and manage information about security certificates (SSL/TLS) for various hosts and their associated services.

## Project Structure

The project is organized as follows:

combined_ssl_certificate_checker/
├── go/ # Backend server in Go
│ ├── main.go
│ ├── go.mod
│ └── go.sum
└── vue/ # Frontend client in Vue.js
├── package.json
├── src/
│ ├── App.vue
│ └── components/
│ ├── ImportExport.vue
│ ├── CertificateInfo.vue
│ ├── ServiceList.vue
│ └── HelloWorld.vue
└── public/
└── index.html

## Prerequisites

### For Go Server

Ensure you have Go installed on your system. You can download it from [here](https://golang.org/dl/).

### For Vue.js Client

Ensure you have Node.js and npm installed on your system. You can download them from [here](https://nodejs.org/).

## Installation Steps

### Clone the Repository

git clone https://github.com/FominDenisAn/combined_ssl_certificate_checker.git

cd combined_ssl_certificate_checker

###Install Dependencies

#For Go Server

cd go

go get github.com/fatih/color

go get github.com/shirou/gopsutil/v3

go get github.com/vbauerster/mpb/v7

go mod tidy

###For Vue.js Client

cd vue
npm install

###Running the Application

You can run both the Go server and Vue.js client using concurrently.

#Update package.json in Vue.js Client

Make sure you have the following script in your vue/package.json:

{
  "scripts": {
    "serve": "vue-cli-service serve",
    "build": "vue-cli-service build",
    "lint": "vue-cli-service lint",
    "start": "concurrently \"cd ../go && go run main.go\" \"npm run serve\""
  }
}

#Install concurrently if not already installed:

npm install concurrently --save-dev

#Run the Application

cd vue

npm run start

#This command will start both the Go server and the Vue.js client simultaneously.

###Usage Examples

#Retrieve list of services:

curl http://localhost:8080/services

#Retrieve certificate information for a specific host:

curl "http://localhost:8080/certificate?host=host1.example.com"

#Import service data:

curl -X POST -H "Content-Type: application/json" -d '{"Service A": ["host1.example.com", "host2.example.com"]}' http://localhost:8080/import

#Export service data:

curl http://localhost:8080/export

###Components

###Vue.js Components

ServiceList.vue: Displays the list of services and their associated hosts. Allows selecting a host to view its certificate information.

CertificateInfo.vue: Displays detailed information about the selected host's certificate.

ImportExport.vue : Provides functionality to import and export service data in JSON format.


#License

This project is licensed under the MIT License - see the LICENSE file for details.
