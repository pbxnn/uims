#!/bin/bash

gogen model datasource -url 'pears:cy7m0yp@,FperzI45@tcp(152.136.131.96:3306)/zxn_test' -t '*' -o './internal/data/dao' -style 'go_zero'
go get github.com/google/wire/internal/wire@v0.5.0
go get github.com/google/wire/cmd/wire@v0.5.0

