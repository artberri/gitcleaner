# gitcleaner

[![Build Status](https://travis-ci.org/artberri/gitcleaner.svg?branch=master)](https://travis-ci.org/artberri/gitcleaner)
[![GoDoc](https://godoc.org/github.com/artberri/gitcleaner?status.svg)](https://godoc.org/github.com/artberri/gitcleaner)
[![Go Report Card](https://goreportcard.com/badge/artberri/gitcleaner)](https://goreportcard.com/report/artberri/gitcleaner)
[![Coverage Status](https://coveralls.io/repos/github/artberri/gitcleaner/badge.svg?branch=master)](https://coveralls.io/github/artberri/gitcleaner?branch=master)

`gitcleaner` is a command line tool to ease the cleaning of your Git repository history.

Recommended for:

- Heavy/big file removal
- Secrets removal

**Disclaimer**:

This is just a training project and these are the first lines I've written 
in Golang, use it as your own risk. Any help would be appreciated.

## Usage

```bash
gitcleaner [global options] command [command options] [/path/to/your/repo]
```

If no path argument is given the current path will be used.

### Available commands

```bash
# List all comands options
gitcleaner help
```

```bash
# List specific command options
gitcleaner help <command>
```

#### List Command

List heavier file objects in the repository history

```bash
gitcleaner list [command options] [/path/to/your/repo]
```

Options:

Option            | Shortname | Description
---               | ---       | ---
`--humanreadable` | `--hr`    | Outputs the size in a readable format
`--unique`        | `-u`      | Outputs the size of the whole history of each file
`--lines NUM`     | `-n NUM`  | Output a maximum of NUM files, 0 = no limit (default: 10)

Recommended usage:

```bash
gitcleaner list -u --hr # List files with the heavier history size in human readable format
```

## And Windows?

There is not support for Windows on CMD now, but you can run it on Windows
using the Git Bash.

## License

gitcleaner. Git housekeeping utility.

Copyright (C) 2017 Alberto Varela <alberto@berriart.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.
