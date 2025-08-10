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

## Object Structure

Your custom Git implementation uses three main object types, each stored as a file named by its SHA-256 hash in `.cgit/object/`:

### 1. **Blob**
- **Purpose:** Stores the raw content of a file.
- **Format:**
  ```
  blob\x00<file content>
  ```
- **Example:**
  ```
  blob\x00Hello, world!
  ```

### 2. **Tree**
- **Purpose:** Represents a directory and its contents (files and subdirectories).
- **Format:**
  Each line describes one entry:
  ```
   tree\x00<objectType> <objectId> <fileName>
  ```
  - `objectType`: `blob` for files, `tree` for subdirectories
  - `objectId`: SHA-256 hash of the referenced object (blob or tree)
  - `fileName`: Name of the file or directory in the parent directory

- **Example:**
  ```
  tree\x00blob 123abc... test.txt
  tree 456def... subdir
  ```

### 3. **Commit**
- **Purpose:** Records a snapshot of the project, referencing a tree and including a commit message.
- **Format:**
  ```
  commit\x00tree <treeObjectId>

  commit <commit message>
  ```
- **Example:**
  ```
  commit\x00tree 456def...

  commit Initial commit
  ```

---

All objects are stored in `.cgit/object/` as files named by their hash. The first part of each object file indicates its type (`blob`, `tree`, or `commit`). Trees reference blobs (files) and other trees (subdirectories). Commits reference a tree and store a message.

---

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
