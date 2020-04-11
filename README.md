# tree
![go-version](https://img.shields.io/github/go-mod/go-version/thatisuday/tree?label=Go%20Version) &nbsp;
![CI](https://github.com/thatisuday/tree/workflows/CI/badge.svg?style=flat-square) &nbsp;
![release](https://github.com/thatisuday/tree/workflows/release/badge.svg?style=flat-square)

A command-line tool to print the contents of a relative or an absolute directory in a tree-like format. It conditionally displays file size and file mode in a neat and colorful format.

![demo](/assets/demo.gif)

> This CLI application is made using [Commando](https://github.com/thatisuday/commando).

## Installation
```
$ GO111MODULE=on go get -u "github.com/thatisuday/tree"
```

## Usage
```
$  tree --help

This tool lists the contents of a directory in tree-like format.
It can also display information about files and folders like size, permission and ownership.

Usage:
   tree [dir] {flags}
   tree <command> {flags}

Commands: 
   info                          displays detailed information of a directory
   help                          displays usage informationn
   version                       displays version number

Arguments: 
   dir                           local directory path (default: ./)

Flags: 
   --no-color                    ignore colored output (default: false)
   -h, --help                    displays usage information of the application or a command (default: false)
   --ignore                      ignore directories (separated by comma) (default: .git,node_modules)
   -l, --level                   level of depth to travel (default: 1)
   --mode                        display mode of the each file (default: false)
   --size                        display size of the each file (default: false)
   -v, --version                 displays version number (default: false)
```

## Example
```
$ tree /projects/commando -l=2 --size --mode
├── .DS_Store (6.1kb) (644)
├── .github (755)
|  └── workflows (755)
├── .gitignore (269 bytes) (644)
├── LICENSE (1.1kb) (644)
├── README.md (23.8kb) (644)
├── assets (755)
|  ├── logo.png (285.0kb) (644)
|  └── logo.svg (87.0kb) (644)
├── commando.go (21.5kb) (644)
├── commando_test.go (10.7kb) (644)
├── demo (755)
|  ├── .DS_Store (6.1kb) (644)
|  ├── reactor.cast (6.3kb) (600)
|  ├── reactor.gif (308.3kb) (644)
|  ├── reactor.gif.sh (49 bytes) (644)
|  └── reactor.go (4.2kb) (644)
├── go.mod (94 bytes) (644)
├── go.sum (533 bytes) (644)
├── templates.go (1.5kb) (644)
└── tests (755)
   ├── empty-exec-name.go (172 bytes) (644)
   ├── invalid-default-value.go (259 bytes) (644)
   ├── missing-action-function.go (208 bytes) (644)
   └── valid-registry.go (4.4kb) (644)
```
