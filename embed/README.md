# Golang Embed

## Brief overview

In version 1.16 of golang, it became possible to embed files into a binary at compile time using the `//go:embed` directive and the ability to access embedded files using the `embed` package.

To use the directive, you must import the `embed` package. At least with the _ sign.
First, we need to initialize the variable in which our files will be stored using the // go: embed directive. The directive is placed only before the variable declaration and initializes it with contents of the file or the file system.

In the next simple example shown the way to include version information from a `version.txt` file:

```go
package main

import (
    _ "embed"
    "fmt"
    "strings"
)

var (
    Version string = strings.TrimSpace(version)
    //go:embed version.txt
    version string
)

func main() {
    fmt.Printf("Version %q\n", Version)
}
```

## Usage details

- The directive `//go:embed` contains paths to files separated by a space and path patterns can be used. Paths cannot use `..`, that is, only files that are inside the current package are searched. To get access to files from a folder that is outside of our package, you need to get the contents of that files in that folder using `embed` and transfer it to the package we need.
- The directive can only be used for global variables (package scope), but not with local variables.
- When compiling, if no files are found, we will receive an error `no matching files found`.
- Variables can store the contents of files in the following types: `string`, `[]byte` or `embed.FS`.
- For `string` and `[]byte` types, there can be only one pattern in the directive, according to which only one file can be found. Otherwise, an error will occur.
- For a file system, multiple files can be used in a tree view.
- FS implements the `FS` interface in `io/fs`, so this type can be used in any package that understands file systems, for example `net/http`, `text/template`, `html/template`.

    ```go
    type FS interface {
        Open(name string) (File, error)
    }
    ```

## Useful links

- [Effective Go/Embedding](https://golang.org/doc/effective_go#embedding)
- [How to Use //go:embed](https://blog.carlmjohnson.net/post/2021/how-to-use-go-embed/)
- [Package embed documentation](https://golang.org/pkg/embed/)
- [Source file src/embed/embed.go](https://golang.org/src/embed/embed.go)