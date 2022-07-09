# multik8s

![CI workflow](https://github.com/eliasbokreta/multik8s/actions/workflows/main.yml/badge.svg)

`multik8s` is a tool to access pod information across multiple Kubernetes contexts at once.

## Usage
- Follow logs from pods in a given namespace :
    - `multik8s get logs [-n namespace] [-p pod] [--follow] [--tail nbr]`
- List pods in a given namespace :
    -  `multik8s get pods [-n namespace] [-p pod]`

## Limitations
Currently, the namespace should be the same on all targeted clusters.

## Setup
Tested on ARM64 using Go 1.18.3.

`make build`
