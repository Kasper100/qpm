package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

var versionnumbers = "0.0.1" // still 0.0.1, because its not released yet
var versionname = "Cheers"   // random name related to how i felt
var version = versionnumbers + " (" + versionname + ")"

func showversion() {
	banner := `
.-----.-----.--------.
|  _  |  _  |        |   ver: %s
|__   |   __|__|__|__|
   |__|__|            
Quiet PacMan: Made for `
	fmt.Printf(banner+"\x1b[4mminimalism\x1b[0m\n", version)
}

// thanks chatgpt 👍👍
func containsString(slice []string, value string) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}

func removeString(slice []string, target string) []string {
	for i, v := range slice {
		if v == target {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func usepac(paccmd string, pkgs []string, returnstr string, pkgmanager string, out bool) {
	installpkgs := append([]string{pkgmanager, paccmd}, pkgs...)
	installpkgs = append(installpkgs, "--noconfirm")
	if pkgmanager == "yay" {
		execCmd := exec.Command("yay", installpkgs...)
		output, err := execCmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error executing command: %s\n", err)
			return
		}
		if out {
			fmt.Printf("%s\n", string(output))
		}
	} else {
		execCmd := exec.Command("sudo", installpkgs...)
		output, err := execCmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Error executing command: %s\n", err)
			return
		}
		if out {
			fmt.Printf("Output: %s\n", string(output))
		}
	}
	fmt.Printf(returnstr + "\n")
}
func TranslateFileToPKGS(file *os.File) []string {
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading file: %s\n", err)
		return []string{}
	}
	content := string(data)
	words := strings.Fields(content)
	return words
}

func showhelp() {
	fmt.Println("Usage: qpm [command] [packages]")
	fmt.Println("Commands:")
	fmt.Println("  -S <pkgs>     Install packages")
	fmt.Println("  -R <pkgs>     Remove packages")
	fmt.Println("  -U            Update system")
	fmt.Println("  -?            Search for packages")
	fmt.Println("  -F            Download packages from a .qpm file")
	fmt.Println("  -RF           Remove packages from a .qpm file")
	fmt.Println("  -V            Shows qpm version")
	fmt.Println("  -H            Show this help message")
	fmt.Println("Options:")
	fmt.Println("  --yay         Switches Package Manager from pacman to yay")
	fmt.Println("  --out         Shows output, mostly used for debugging")
}

func main() {
	// Setup Ctrl+C handler
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigChan
		fmt.Println("\nReceived signal:", sig)
		cmd := exec.Command("sudo", "rm", "-rf", "/var/lib/pacman/db.lck")
		output, err := cmd.CombinedOutput()

		if err != nil {
			fmt.Println("Failed to delete lock file:", err)
		} else {
			fmt.Println(string(output))
			fmt.Println("Deleted lock file: /var/lib/pacman/db.lck")
		}

		os.Exit(0)
	}()
	if len(os.Args) < 2 {
		showhelp()
		return
	}
	cliargs := os.Args[1:2]
	rawargs := os.Args[2:]
	pkgmanager := "pacman"
	if len(cliargs) == 0 {
		fmt.Println("No command provided. Please provide a command.")
		return
	}
	pkgs := rawargs
	if containsString(pkgs, "--yay") {
		pkgmanager = "yay"
		pkgs = removeString(pkgs, "--yay")
	}
	showoutput := false
	if containsString(pkgs, "--out") {
		showoutput = true
		pkgs = removeString(pkgs, "--out")
	}
	switch cliargs[0] {
	case "-S":
		if len(pkgs) == 0 {
			fmt.Println("No packages provided. Please provide packages to install.")
			return
		}
		returnedstr := ""
		if len(pkgs) > 1 {
			returnedstr = "\nThe packages are installed"
		} else {
			returnedstr = "\nThe package is installed"
		}
		fmt.Printf("Installing packages: %s\n", strings.Join(pkgs, ", "))
		usepac("-S", pkgs, returnedstr, pkgmanager, showoutput)
	case "-R":
		if len(pkgs) == 0 {
			fmt.Println("No packages provided. Please provide packages to remove.")
			return
		}
		returnedstr := ""
		if len(pkgs) > 1 {
			returnedstr = "\nThe packages are removed"
		} else {
			returnedstr = "\nThe package is removed"
		}
		fmt.Printf("Removing packages: %s\n", strings.Join(pkgs, ", "))
		usepac("-Rns", pkgs, returnedstr, pkgmanager, showoutput)
	case "-?":
		// search for packages
		if len(pkgs) == 0 {
			fmt.Println("No packages provided. Please provide packages to search.")
			return
		}
		returnedstr := "\n\x1b[47;30mSearch results are above\x1b[0m"
		fmt.Printf("Searching packages: %s\n", strings.Join(pkgs, ", "))
		usepac("-Ss", pkgs, returnedstr, pkgmanager, showoutput)
	case "-U":
		fmt.Println("Updating system...")
		returnedstr := "\n\x1b[47;30mSystem updated\x1b[0m"
		usepac("-Syu", []string{}, returnedstr, pkgmanager, showoutput)
	case "-F":
		if len(pkgs) == 0 {
			fmt.Println("No packages provided. Please provide packages to download from a .qpm file.")
			return
		}
		returnedstr := "\n\x1b[47;30mPackages downloaded from %s\x1b[0m"
		file, err := os.Open(pkgs[0])
		if err != nil {
			panic(err)
		}
		defer file.Close()
		TranslatedPKGS := TranslateFileToPKGS(file)
		if len(TranslatedPKGS) == 0 {
			fmt.Println("No packages found in the file.")
			return
		}
		fmt.Printf("Downloading packages from file: %s\n", TranslatedPKGS)
		usepac("-S", TranslatedPKGS, fmt.Sprintf(returnedstr, pkgs[0]), pkgmanager, showoutput)
	case "-RF":
		if len(pkgs) == 0 {
			fmt.Println("No packages provided. Please provide packages to download from a .qpm file.")
			return
		}
		returnedstr := "\n\x1b[47;30mPackages removed from %s\x1b[0m"
		file, err := os.Open(pkgs[0])
		if err != nil {
			panic(err)
		}
		defer file.Close()
		TranslatedPKGS := TranslateFileToPKGS(file)
		if len(TranslatedPKGS) == 0 {
			fmt.Println("No packages found in the file.")
			return
		}
		fmt.Printf("Removing packages from file: %s\n", TranslatedPKGS)
		usepac("-Rns", TranslatedPKGS, fmt.Sprintf(returnedstr, pkgs[0]), pkgmanager, showoutput)
	case "-H":
		showhelp()
	case "-V":
		showversion()
	default:
		fmt.Printf("Unknown command: %s\n", cliargs[0])
		fmt.Println("Use -H for help.")
	}
}
