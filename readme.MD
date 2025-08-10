# buildGit

A minimal custom Git-like version control system implemented in Go.  
This project demonstrates basic concepts of content-addressable storage, trees, and commits, similar to how Git works internally.

## Features

- **init**: Initialize a `.cgit` object database in the current directory.
- **hash**: Hash a file and store its content as a blob object.
- **cat**: Print the content of a stored object by its object ID.
- **write-tree**: Store the directory structure as a tree object.
- **get-tree**: Restore a directory structure from a tree object.
- **commit**: Create a commit object referencing a tree and a commit message.

## Getting Started

### Prerequisites

- Go 1.22 or newer

### Installation

Clone this repository and build the project:

```sh
git clone <your-repo-url>
cd buildGit
go build
```

### Usage

Run commands using `go run main.go [command] [args...]`:

#### Initialize repository

```sh
go run main.go init
```

Creates a `.cgit` directory for storing objects.

#### Hash a file

```sh
go run main.go hash <file-path>
```

Hashes the file and stores it as a blob object.

#### Show object content

```sh
go run main.go cat <object-id> blob
```

Prints the content of the blob object.

#### Write a tree object

```sh
go run main.go write-tree <directory-path>
```

Stores the directory structure as a tree object and prints its object ID.

#### Restore a tree

```sh
go run main.go get-tree <target-directory> <tree-object-id>
```

Restores the directory structure from the specified tree object.

#### Commit

```sh
go run main.go commit "<commit-message>" <directory-path>
```

Creates a commit object referencing the tree of the given directory.

## Object Storage

All objects are stored in `.cgit/object/` as files named by their SHA-256 hash.

- **blob**: Stores file content.
- **tree**: Stores directory structure (list of blobs and subtrees).
- **commit**: Stores reference to a tree and a commit message.

## Project Structure

- [`main.go`](main.go): CLI entry point.
- [`command/command.go`](command/command.go): Implements commands.
- [`data/data.go`](data/data.go): Handles object storage and retrieval.
- [`enum/objectType.go`](enum/objectType.go): Defines object types.
- [`util/util.go`](util/util.go), [`util/constant.go`](util/constant.go): Utility functions and constants.
- [`it/`](it/): Integration tests for commands.

## Running Tests

```sh
go test ./...
```

## Example

```sh
go run main.go init
go run main.go hash somefile.txt
go run main.go write-tree mydir
go run main.go commit "Initial commit" mydir
```

## License

MIT (or specify your license)