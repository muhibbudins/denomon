# Denomon

Simple watcher file for build an application using deno

[![asciicast](https://asciinema.org/a/kkoeCdKB5bKgCLY7XzHsmbFZy.png)](https://asciinema.org/a/kkoeCdKB5bKgCLY7XzHsmbFZy)

> Note : The pointer printed only on asciinema, but not on the bash / zsh

### Installation

Download Wget

```bash
$ wget -O denomon https://raw.githubusercontent.com/muhibbudins/denomon/master/denomon
```

or, CURL
```bash
$ curl -LJO https://raw.githubusercontent.com/muhibbudins/denomon/master/denomon
```

Move the denomon binary into **/usr/bin** or **/usr/local/bin**

```bash
$ mv denomon /usr/local/bin
```

Make it executable

```bash
$ sudo chmod +x /usr/local/bin/denomon
```

### Usage

```bash
$ denomon <options> <file>
```

Example

```bash
$ denomon --dir fixtures --allow net express.ts
or
$ denomon --allow net fixtures/express.ts
or
$ denomon fixtures/mod.ts
```

### Options

#### --dir

To set target directory

#### --allow

To set allowed permission

### Features

- Auto build for single file
- Watching all files in folder
- Auto reload build for child process (ie. net)

### Limitations

#### --dir

Refer to this [issue](https://github.com/fsnotify/fsnotify/issues/18), currently this tools only can watch file in one folder. But not recursively.

### License

This project under MIT License