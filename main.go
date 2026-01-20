package main

import "fmt"

func main() {
	x := 10
	fmt.Println(x)
}

//Задание №10
//создаём в /GolandProjects/Learning_EM/.git/hooks файл pre-commit
//#!/bin/sh
//
//echo "running go test..."
//go test ./...
//if [ $? -ne 0 ]; then
//  echo "go test failed, aborting commit"
//  exit 1
//fi
//
//echo "running golangci-lint..."
//golangci-lint run ./...
//if [ $? -ne 0 ]; then
//  echo "golangci-lint failed, aborting commit"
//  exit 1
//fi
//
//echo "pre-commit checks passed"
//chmod +x .git/hooks/pre-commit
//
//
//теперь при комитах прогоняется go test ./... golangci-lint run ./...
//если ошибка - откат попытки комита
//No files committed, 1 file failed to commit: Задание №10 (first test) running go test... # Learning_EM [Learning_EM.test] .\main.go:4:2: declared and not used: x FAIL Learning_EM [build failed] FAIL go test failed, aborting commit
//исправляем ошибку и теперь норм коммитится
