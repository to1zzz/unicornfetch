# 🦄 UnicornFetch

<div align="center">

![GO](https://img.shields.io/github/go-mod/go-version/to1zzz/unicornfetch?style=for-the-badge&logo=go&label=Go&color=00ADD8)
[![Linux Support](https://img.shields.io/badge/Linux-FCC624?style=for-the-badge&logo=linux&logoColor=black)](https://kernel.org/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge)](LICENSE)
[![GitHub last commit](https://img.shields.io/github/last-commit/to1zzz/unicornfetch?style=for-the-badge&color=blueviolet)](https://github.com/to1zzz/unicornfetch/commits/main)

**A minimalistic brainrot fetch utility for Linux, written in Go.**

<br/>
</div>

## ✨ Features

- **Fast & Lightweight** – Single Go binary with zero dependencies.
- **Clean & Colorful** – Displays system info next to a hungry unicorn ASCII art.
- **Linux Optimized** – Built specifically for Linux-based systems.
- **Easy to Install** – Simple build from source and copy to your `$PATH`.
- **Extensible** – Pure Go code, easy to modify and customize.
- **init** – Сan show init out of the box.
  
## 🖥️ Showcase

Here is UnicornFetch in action on a Gentoo Linux system:

```text
$ unicornfetch

          \/`-.,    OS          Gentoo Linux
       _ _\         Kernel      6.6.16-gentoo-dist
      ('> ('>       Uptime      2d 4h 23m
      /\"( /\"(     Packages    1423 (Gentoo)
      \_)` \_)`     Init        OpenRC
      mrf  mrf      WM          Niri
                    CPU         AMD Ryzen 9 7950X
unicorn is hungry   GPU         AMD Radeon RX 6900 XT
                    Memory      4.2 / 31.3 GiB
                    Disk        128G / 512G (25%)
                    Terminal    kitty

```
## 📦 Installation

### Prerequisites

* [Go](https://go.dev/dl/) (version 1.21+ recommended)
* `git` (to clone the repository)

### Steps

1. **Clone the repository:**
```bash
git clone https://github.com/to1zzz/unicornfetch.git
cd unicornfetch

```

2. **Build the binary:**
```bash
go build -o unicornfetch

```

3. **Move to a directory in your `PATH` (optional but recommended):**
```bash
sudo cp unicornfetch /usr/local/bin/

```

Now you can run `unicornfetch` from anywhere in your terminal.

## 🚀 Usage

Simply run the compiled binary:

```bash
./unicornfetch

```

Or if you've added it to your `PATH`:

```bash
unicornfetch

```

## 🛠️ Build from source (without installing)

If you prefer not to install the binary system-wide, you can run it directly from the project directory:

```bash
go run .

```

## 📜 License

This project is licensed under the **MIT License**. See the [LICENSE](LICENSE) file for details.

---
