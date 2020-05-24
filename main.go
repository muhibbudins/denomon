package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
)

var pid int = 0
var command exec.Cmd
var args []string
var file string = ""
var path string = ""
var permission []string

func help() {
	fmt.Printf("[DENOMON HELP]\n\n")
	fmt.Printf("Usage:\n\n")
	fmt.Printf("denomon <options> <file>\n\n")
	fmt.Printf("Options:\n\n")
	fmt.Printf("--help Print this help message\n")
	fmt.Printf("--dir Assign directory to watch\n")
	fmt.Printf("--allow Assign permission for the files\n\n")
}

func build(file string) bool {
	color.Green("[denomon] build: %s", file)

	if pid > 0 {
		command.Process.Kill()
	}

	var stderr bytes.Buffer

	options := []string{"run", file}

	if len(permission) > 0 {
		options = []string{"run"}
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
		log.Fatal(err)
	}

	dirPtr := flag.String("dir", ".", "Assign directory to watch")
	helpPtr := flag.Bool("help", false, "Print this help message")
	allowPtr := flag.String("allow", "", "Assign permission for the files")

	flag.Parse()

	if *allowPtr != "" {
		allowSlice := strings.Split(*allowPtr, ",")

		for i := range allowSlice {
			allowSlice[i] = strings.TrimSpace(allowSlice[i])
		}

		permission = allowSlice
	}

	args = flag.Args()

	if (len(args) == 0 && *dirPtr == ".") || *helpPtr == true {
		help()
		return
	}

	if len(args) > 0 {
		file = args[0]
	}

	if file != "" {
		path = dir + "/" + *dirPtr + "/" + file
		build(path)
	} else {
		path = dir + "/" + *dirPtr
	}

	color.Cyan("[denomon] watching: %s", path)

	watcher, err := fsnotify.NewWatcher()

	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}

				if event.Op&fsnotify.Write == fsnotify.Write {
					repeat := build(event.Name)

					if repeat {
						build(event.Name)
					}
				}
			case _, ok := <-watcher.Errors:
				if !ok {
					return
				}
				color.Red("[denomon] error: %s", err.Error())
				os.Exit(1)
			}
		}
	}()

	if err = watcher.Add(path); err != nil {
		color.Red("[denomon] error: %s", err.Error())
		os.Exit(1)
	}
	<-done
}
