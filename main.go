package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/radovskyb/watcher"
)

var pid int = 0
var command exec.Cmd
var args []string
var file string = ""
var path string = ""
var permission []string
var mainFile string = ""
var isMainFile bool = false
var isUnstable bool = false

func help() {
	fmt.Printf("[DENOMON HELP]\n\n")
	fmt.Printf("Usage:\n\n")
	fmt.Printf("denomon <options> <file>\n\n")
	fmt.Printf("Options:\n\n")
	fmt.Printf("--version Showing denomon version\n")
	fmt.Printf("--help Print this help message\n")
	fmt.Printf("--dir Assign directory to watch\n")
	fmt.Printf("--allow Assign permission for the files\n\n")
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func build(file string) bool {
	if !fileExists(file) {
		return false
	}

	color.Green("[denomon] build: %s", file)

	if pid > 0 {
		command.Process.Kill()
	}

	var stderr bytes.Buffer

	options := []string{"run"}

	if isUnstable {
		options = append(options, "--unstable")
	}

	options = append(options, file)

	if len(permission) > 0 {
		options = []string{"run"}

		if isUnstable {
			options = append(options, "--unstable")
		}

		for i := 0; i < len(permission); i++ {
			options = append(options, "--allow-"+permission[i])
		}
		options = append(options, file)
	}

	cmd := exec.Command("deno", options...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	fmt.Print("\033[s")
	fmt.Print(cmd.Stdout)
	fmt.Print("\033[u\033[K")

	if err := cmd.Start(); err != nil {
		color.Red("[denomon] error: %s", fmt.Sprint(err)+": "+stderr.String())
	}

	pid = cmd.Process.Pid
	command = *cmd

	return false
}

func main() {
	dir, err := os.Getwd()

	if err != nil {
		color.Red("[denomon] error: %s", err.Error())
	}

	dirPtr := flag.String("dir", "", "Assign directory to watch")
	helpPtr := flag.Bool("help", false, "Print this help message")
	allowPtr := flag.String("allow", "", "Assign permission for the files")
	versionPtr := flag.Bool("version", false, "Showing denomon version")
	unstablePtr := flag.Bool("unstable", false, "Assign unstable option for Deno")

	flag.Parse()

	if *unstablePtr {
		isUnstable = true
	}

	if *allowPtr != "" {
		allowSlice := strings.Split(*allowPtr, ",")

		for i := range allowSlice {
			allowSlice[i] = strings.TrimSpace(allowSlice[i])
		}

		permission = allowSlice
	}

	args = flag.Args()

	if *versionPtr == true {
		color.Green("[denomon] version %s", "0.1.3")
		return
	}

	if *helpPtr == true {
		help()
		return
	}

	if len(args) > 0 {
		file = args[0]
	}

	path = dir + "/"

	if *dirPtr != "" {
		path = dir + "/" + *dirPtr + "/"
	}

	if file != "" {
		isMainFile = true
		mainFile = path + file

		build(mainFile)
	}

	color.Cyan("[denomon] watching: %s", path)

	w := watcher.New()
	w.IgnoreHiddenFiles(true)

	go func() {
		for {
			select {
			case event := <-w.Event:
				if !event.IsDir() {
					if isMainFile {
						repeat := build(mainFile)

						if repeat {
							build(mainFile)
						}
					} else {
						repeat := build(event.Path)

						if repeat {
							build(event.Path)
						}
					}
				}
			case err := <-w.Error:
				color.Red("[denomon] error: %s", err.Error())
			case <-w.Closed:
				return
			}
		}
	}()

	if err := w.AddRecursive(path); err != nil {
		color.Red("[denomon] error: %s", err.Error())
	}

	if err := w.Start(time.Millisecond * 100); err != nil {
		color.Red("[denomon] error: %s", err.Error())
	}
}
