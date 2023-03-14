#!/bin/bash

go fmt cmd/vending_machine/*.go
go fmt cmd/test_harness/*.go
go fmt internal/api/*.go
go fmt internal/vending_machine/*.go
go fmt pkg/error_catching/*.go
