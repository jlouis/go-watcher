# Watcher.
Simple go program that watches a directory (not recursively) and applies an action whenver a change is detected.
Inside the file action.go, there are a few examples. An action can be anything that implements an interface like :

```golang
type Action interface {
	Do(string) error
}
```

# CMD watcher
The executable command provided is configured to watch for msgpack files and ship them to s3. 
So it exits if the credentials to s3 are not found in the environment.

## building
This project uses gb to manage reproducible builds. Grab gb from getgb.io.
Clone this repo and then:

```
gb build
```

Then a binary will be created in ./bin/ folder.


## usage

```
 ./watcher - h

>>> Usage of ./watcher:
  -path string
        Directory to watch for changes. Not recursively. (default ".")
```

