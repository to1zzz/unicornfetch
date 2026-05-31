🦄 UnicornFetch

A minimalistic fetch utility for Linux, written in Go.

✨ Features

Fast & Lightweight – Single Go binary with zero dependencies.

Clean & Colorful – Displays system info next to a hungry unicorn ASCII art.

Linux Optimized – Built specifically for Linux-based systems.

Easy to Install – Simple build from source and copy to your $PATH.

Extensible – Pure Go code, easy to modify and customize.

🖥️ Showcase

Here is UnicornFetch in action on a Gentoo Linux system:

$ unicornfetch

          \/`-.,    OS          Gentoo Linux
       _ _\         Kernel      6.6.16-gentoo-dist
      ('> ('>       Uptime      2d 4h 23m
      /\"( /\"(     Packages    1423 (Gentoo)
      \_)` \_)`     Init        OpenRC
      mrf mrf       WM          Niri
                    CPU         AMD Ryzen 9 7950X
unicorn is hungry   GPU         AMD Radeon RX 6900 XT
                    Memory      4.2 / 31.3 GiB
                    Disk        128G / 512G (25%)
                    Terminal    kitty


📦 Installation

Prerequisites

Go (version 1.21+ recommended)

git (to clone the repository)

Steps

Clone the repository:

git clone https://github.com/to1zzz/unicornfetch.git
cd unicornfetch


Build the binary:

go build -o unicornfetch


Move to a directory in your PATH (optional but recommended):

sudo cp unicornfetch /usr/local/bin/


Now you can run unicornfetch from anywhere in your terminal.

🚀 Usage

Simply run the compiled binary:

./unicornfetch


Or if you've added it to your PATH:

unicornfetch


🛠️ Build from source (without installing)

If you prefer not to install the binary system-wide, you can run it directly from the project directory:

go run .


📜 License

This project is licensed under the MIT License. See the LICENSE file for details.
