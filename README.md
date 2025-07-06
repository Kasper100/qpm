# qpm — Quiet Pacman

> **Made for minimalism.**

`qpm` is a minimalist wrapper for Arch Linux’s `pacman`, built to streamline package management with quiet, clean output. No clutter, no noise — just packages.

---

## WARNING
### THIS PROJECT WAS MADE TO LEARN GOLANG. 
### EXPECT BUGS

---

## Features

* 🛠 Install, remove, and update packages quietly
* 📜 Install/remove packages from `.qpm` files (space or newline format)
* 🤫 Minimal terminal output — no distractions
* 🔀 Use `--yay` to switch to `yay` for AUR support
* 🖥 `--out` flag to show raw pacman/yay output (debugging)

---
## Requirement
* bash
* golang / go
``` bash
sudo pacman -S go bash
```
## Installation
Bash only:
``` bash
git clone https://github.com/Kasper100/qpm.git
mv qpm .qpm
cd .qpm
./install.sh
qpm -V
```
---

## Usage

```bash
qpm [command] [packages]
```

### Commands

| Command     | Description                         |
| ----------- | ----------------------------------- |
| `-S <pkgs>` | Install packages                    |
| `-R <pkgs>` | Remove packages                     |
| `-U`        | Update system                       |
| `-?`        | Search for packages                 |
| `-F`        | Install packages from a `.qpm` file |
| `-RF`       | Remove packages from a `.qpm` file  |
| `-V`        | Show qpm version                    |
| `-H`        | Show help message                   |

### Options

| Option  | Description                   |
| ------- | ----------------------------- |
| `--yay` | Use `yay` instead of `pacman` |
| `--out` | Show full command output      |

---

## .qpm Files

You can store packages in a simple `.qpm` text file using either of the following formats:

**Space-separated**

```
cmatrix cowsay fastfetch btop
```

**Newline-separated**

```
cmatrix
cowsay
fastfetch
btop
```

Then install them with:

```bash
qpm -F my-packages.qpm
```

Or remove them:

```bash
qpm -RF my-packages.qpm
```

---

## License

qpm is free software under the GNU General Public License.
It respects your freedom and comes with the rights to use, study, share, and modify the code.
This ensures the program stays free and open for everyone.
