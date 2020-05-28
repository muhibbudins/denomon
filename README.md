# Denomon

## Also read in [Portuguese](https://github.com/muhibbudins/denomon/blob/master/README-pt-BR.md) 

Simple watcher file for build an application using deno

[![asciicast](https://asciinema.org/a/kkoeCdKB5bKgCLY7XzHsmbFZy.png)](https://asciinema.org/a/kkoeCdKB5bKgCLY7XzHsmbFZy)

> Note : The pointer printed only on asciinema, but not on the bash / zsh

### Installation

Install using Wget

```bash
$ wget -O - https://raw.githubusercontent.com/muhibbudins/denomon/master/install.sh | sh
```

or, CURL
```bash
$ curl -s https://raw.githubusercontent.com/muhibbudins/denomon/master/install.sh | sh
```

### Usage

```bash
$ denomon <options> <file>
```

Example:

- Showing help message

```bash
denomon --help
```

- Single command to spawn current folder recursively

```bash
$ denomon
```

- Spawn file with permission [see note below]

```bash
$ denomon --allow net,read server.ts
```

- Spawn file with specific folder and permission [see note below]

```bash
$ denomon --dir fixtures --allow net,read server.ts
```

- Spawn specific folder with permission

```bash
$ denomon --dir fixtures --allow net,read
```

> Note : If you set denomon to spawn single file, all changes at root folder will trigger reload on main file.


### Options

#### --version

Showing denomon version

#### --help

Showing help message

#### --dir

Assign directory to watch

#### --allow

Assign permission for the files

#### --unstable

Assign unstable parameter to deno

### Features

- Auto build for single file
- Watching all files in folder recursively
- Auto reload build for child process (ie. net)

### License

This project is under MIT License

### Stargazers over time

[![Stargazers over time](https://starchart.cc/muhibbudins/denomon.svg)](https://starchart.cc/muhibbudins/denomon)
      