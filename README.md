# GoEnv
A very simple environment variables management library for Go.
This was meant for my personal use, expecially because this is actually the first time I approached Go.
Note that there are mre complete and mature libraries out there.

# Usage
## As a library
```
go get github.com/riccardooliva91/goenv
```

You should call the `Load` method early in your boostrap:

```go
...
import "github.com/riccardooliva91/goenv"

func main() {
    // Pass false as second argument to avoid loading in OS variables
    // However, note that in that case os vars won't be available in the env map

    // Will look for a .env file in the same folder
    env := goenv.Load([], true)

    // Alternatively, you can pass a list of file paths to load:
    env := goenv.Load(["../.env", "../../environment.env", "/.env", ...], true)
    ...
}
```

The environment will only load once. Subsequential calls of the `Load` method will have no effect.
The environment is kept in-memory in a `map[string]string` map. This contains __both the OSvariables AND the env files variables__.
Please note that __the env files DO override the OS variables as well as themselves__ if the same key is present in more then one file, and the last loaded wins.

```go
// You can use either the provided methods or treat the env as a map
env.Set("key", "example")
env["key"] = "example"

// Value: "example"
env.Get("key")
env["key"]

env.Isset("key") // true
env.Isset("notset") // false

env.Delete("key")
delete(env, "key")
```


## As a command
```
go install github.com/riccardooliva91/goenv/cmd/goenv-merge@latest
```

The `goenv-merge` command allows you to merge multiple .env files of your choice into a single one. Note that this command __won't load OS variables__:

```
# outputs the result in a .env file in the same directory
goenv-merge file1 file2 file3

# custom destination
goenv-merge file1 file2 file3 -d /.myenv
# or: goenv-merge file1 file2 file3 --destination /.myenv

# overwrite an existing destination file, if that exists
goenv-merge file1 file2 file3 -o
# or: goenv-merge file1 file2 file3 --overwrite

# complete example
goenv-merge file1 file2 file3 -d /.myenv -o
# same as: goenv-merge file1 file2 file3 -o -d /.myenv
```

# Writing .env files
You can use either a `KEY=VALUE` format or a yaml-like `KEY:VALUE` one. Formats can be mixed within the same file (ypu shouldn't, but you can ;) )
```
# Comments are allowed like this, even inline
KEY = VALUE # Spaces are allowed
KEY="VALUE" # Double quotes will be removed
KEY="VA"LUE" # Will parse as: VA"LUE
KEY= VALUE" # Will parse as VALUE
KEY: VALUE
KEY:VALUE
# For the Yamlish the validations are the same as in the KEY=VALUE format
```