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

*This is just a training project and these are the first lines I've written
in Golang, use it as your own risk. Any help would be appreciated.*

## Downloads

[In this link](https://github.com/artberri/gitcleaner/releases) you will find the packages for every supported platform. Please download the proper package for your operating system and architecture. You can also download older versions.

## Installation

To install Gitcleaner, find the [appropriate package](https://github.com/artberri/gitcleaner/releases)
for your system and download it.

After downloading Gitcleaner, unzip the package. Gitcleaner runs as a single binary named `gitcleaner`. Any other files in the package can be safely removed and Gitcleaner will still function.

The final step is to make sure that the Gitcleaner binary is available on the PATH. See [this page](https://stackoverflow.com/questions/14637979/how-to-permanently-set-path-on-linux) for instructions on setting the PATH on Linux and Mac. [This page](https://stackoverflow.com/questions/1618280/where-can-i-set-path-to-make-exe-on-windows) contains instructions for setting the PATH on Windows.

### Windows

CMD and Powershell are not supported. You need a [Bash](https://en.wikipedia.org/wiki/Bash_(Unix_shell)) shell in order to run Gitcleaner. You can download [Git Bash](https://git-scm.com/download/win) or [Cygwin](https://www.cygwin.com/) if you don't have one yet.

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
