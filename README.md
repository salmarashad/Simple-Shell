# Simple-Shell
A lightweight, interactive shell implementation written in Go. This is created following the basic structure of [this tutorial](https://blog.init-io.net/post/2018/07-01-go-unix-shell/) with all the suggested added features.

## Features
- Command execution with full terminal support
- Directory navigation using `cd`
- Command history navigation with up/down arrow keys
- Input editing with backspace
- Current working directory display in prompt

## Prerequisites
- Go 1.13 or higher
- `github.com/mattn/go-tty` package

## Installation & Usage

### Clone the repository
```bash
git clone https://github.com/salmarashad/Simple-Shell.git
```

### Navigate to the project directory
```bash
cd Simple-Shell
```

### Install dependencies
```bash
go get github.com/mattn/go-tty
```

### Start the shell
```bash
go run .
```
