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

func main() {
	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	var file string = ""
	var directory string = ""

	dirPtr := flag.String("d", ".", "Directory assignment")
	helpPtr := flag.Bool("h", false, "Help message")

	flag.Parse()

	if (len(flag.Args()) == 0 && *dirPtr == ".") || *helpPtr == true {
		help()
		return
	}

	if len(flag.Args()) > 0 {
		file = flag.Args()[0]
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
					color.Green("[denomon] build: %s", event.Name)

					var stderr bytes.Buffer

					options := []string{"run", event.Name}

					if len(flag.Args()) > 1 {
						options = []string{"run"}
						options = append(options, flag.Args()[1:]...)
						options = append(options, event.Name)
					}

					cmd := exec.Command("deno", options...)
					cmd.Stdout = os.Stdout
					cmd.Stderr = os.Stderr
					fmt.Println(cmd.Stdout)

					err := cmd.Run()

					if err != nil {
						color.Red("[denomon] error: %s", fmt.Sprint(err)+": "+stderr.String())
					}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				color.Red("[denomon] error: %s", err.Error())
				os.Exit(1)
			}
		}
	}()

	err = watcher.Add(directory)
	if err != nil {
		color.Red("[denomon] error: %s", err.Error())
		os.Exit(1)
	}
	<-done
}
