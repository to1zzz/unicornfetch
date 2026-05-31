````md
# 🦄 unicornfetch

A lightweight system fetch utility written in Go with a unicorn-themed ASCII logo.

<p align="center">
  <img src="./assets/preview.png" width="700">
</p>

---

## ✨ Features

- Fast and lightweight
- Written in Go
- Unicorn ASCII art
- Linux-friendly output
- Simple and minimal design

---

## 📦 Installation

### Build from source

```bash
git clone https://github.com/to1zzz/unicornfetch.git
cd unicornfetch
go build -o unicornfetch
````

### Install via Go

```bash
go install github.com/to1zzz/unicornfetch@latest
```

---

## 🚀 Usage

```bash
unicornfetch
```

### Options

```bash
unicornfetch --help
unicornfetch --version
unicornfetch --no-color
```

---

## ⚙️ Makefile

```makefile
build:
	go build -o unicornfetch .

run:
	go run .

install:
	go install .

clean:
	rm -f unicornfetch
```

---

## 📁 Project Structure

```text
.
├── cmd/
│   └── unicornfetch/
│       └── main.go
├── internal/
│   └── fetch/
│       └── fetch.go
├── assets/
│   └── preview.png
├── go.mod
├── Makefile
└── README.md
```

---

## 🧠 Motivation

A small personal project inspired by tools like fastfetch and neofetch, but with a more playful aesthetic.

---

## 🧪 Development

```bash
git clone https://github.com/to1zzz/unicornfetch.git
cd unicornfetch
go run ./cmd/unicornfetch
```

---

## 🔧 Build

```bash
go build -o unicornfetch ./cmd/unicornfetch
```

---

## 🚀 CI (GitHub Actions)

```yaml
name: build

on: [push, pull_request]

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
      - uses: actions/checkout@v4

      - uses: actions/setup-go@v5
        with:
          go-version: "1.22"

      - run: go build ./...
```

---

## 📜 License

MIT

```
```
