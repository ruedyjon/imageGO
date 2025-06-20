#!/bin/bash

# Build the Go binary from source
cd pkg/imageGO
go build -o ../../imagego
cd ../..

# Continue with Node.js install
npm install
