#!/usr/bin/env bash

# In Go 1.6 through 1.8, the ./... matched also the vendor directory.
go test $(go list ./... | grep -v /vendor/)