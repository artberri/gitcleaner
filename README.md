# gitcleaner

`gitcleaner` is a command line tool to ease the cleaning of your Git repository history.

Recommended for:

- Heavy/big file removal
- Secrets removal

**Disclaimer**:

This is just a training project and these are the first lines I've written in Golang, use it as your own risk. Any help would be appreciated.

## Usage

```bash
gitcleaner [global options] command [command options] [/path/to/your/repo]
```

If no path argument is given the current path will be used.

```bash
# List commands
gitcleaner help
```

gh help <command>
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
gitcleaner list -u --hr
```
