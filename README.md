# go-testfs-insensitive

On Windows:

```console
$ go test
--- FAIL: TestFS (0.01s)
    fs_test.go:23: failed to test file system case insensitively: TestFS found errors:
        expected but not found: FOO/bar/file.txt
FAIL
FAIL    github.com/jakebailey/go-testfs-insensitive     3.446s
FAIL
```
