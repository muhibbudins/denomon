package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
)

var pid int = 0
var command exec.Cmd
var args []string
var file string = ""
var directory string = ""
var permission permissionFlag

func help() {
	fmt.Println("[DENOMON HELP]")
	fmt.Println("Usage:")
	fmt.Println("denomon <options> <file> <allowed-permission>")
	fmt.Println("Options:")
	fmt.Println("-h Show this help message")
	fmt.Println("-d Assign a directory")
	fmt.Println("Allowed Permision:")
	fmt.Println("Use --allow-net etc. sequentially after file name")
}

func build(path string) bool {
	color.Green("[denomon] build: %s", path)

	if pid > 0 {
		command.Process.Kill()
	}

	var stderr bytes.Buffer

	options := []string{"run", path}

	if len(permission) > 0 {
		options = []string{"run"}
		for i := 0; i < len(permission); i++ {
			options = append(options, "--allow-"+permission[i])
		}
		options = append(options, path)
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

type permissionFlag []string

func (i *permissionFlag) String() string {
	return "my string representation"
}

func (i *permissionFlag) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func main() {
	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	flag.Var(&permission, "allow", "Some description for this param.")
	dirPtr := flag.String("dir", ".", "Directory assignment")
	helpPtr := flag.Bool("help", false, "Help message")

	flag.Parse()

	args = flag.Args()

	if (len(args) == 0 && *dirPtr == ".") || *helpPtr == true {
		help()
		return
	}

	if len(args) > 0 {
		file = args[0]
	}

	if file != "" {
		directory = dir + "/" + *dirPtr + "/" + file
	} else {
		directory = dir + "/" + *dirPtr
	}

	color.Cyan("[denomon] watching: %s", directory)

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

	if err = watcher.Add(directory); err != nil {
		color.Red("[denomon] error: %s", err.Error())
		os.Exit(1)
	}
	<-done
}
