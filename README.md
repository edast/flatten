Utility to flatten json into key = value pairs. 
It might be used to compare different json objects.

It accepts json as first cli attribute, or can read from stdin:
```
❯ echo '{"foo":"bar","baz":{"bar":1}}' | ./flatten
foo = bar
baz.bar = 1
```
 
Run `go build -o flatten main.go` to build binary.

For array values in json - multiple lines with the same key will be printed:
```shell
❯ echo '{"foo":"bar","baz":{"bar":[1,2,3,4,5]}}' | flatten
foo = bar
baz.bar = 1
baz.bar = 2
baz.bar = 3
baz.bar = 4
baz.bar = 5
```