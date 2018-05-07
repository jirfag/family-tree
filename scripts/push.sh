#!/usr/bin/env bash
cd ..
GOOS=linux GOARCH=amd64  go build ./family-tree.go
docker build -t registry.cn-hangzhou.aliyuncs.com/fredliang/family-tree  .
docker push  registry.cn-hangzhou.aliyuncs.com/fredliang/family-tree
cd scripts
