# [WIP] Denomon

Simple watcher file for build an application using deno

### Installation

Download Wget

```bash
$ wget https://raw.githubusercontent.com/muhibbudins/denomon/denomon
```

or, CURL
```bash
$ curl https://raw.githubusercontent.com/muhibbudins/denomon/denomon
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

TODO

### Limitations

TODO

### License

This project under MIT License