# go-testfs-insensitive

On Windows or macOS:

```console
$ go test -v
=== RUN   TestFS
    fs_test.go:33: failed to test file system case insensitively: TestFS found errors:
        expected but not found: FOO/BAR/FILE.TXT
--- FAIL: TestFS (0.00s)
FAIL
exit status 1
FAIL	github.com/jakebailey/go-testfs-insensitive	0.006s
```
