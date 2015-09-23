# Shakers
🐹 + 🐙 = 😽 [![Circle CI](https://circleci.com/gh/vdemeester/shakers.svg?style=svg)](https://circleci.com/gh/vdemeester/shakers)

A collection of `go-check` Checkers to ease the use of it.

## Building and testing it

You need either [docker](https://github.com/docker/docker), or `go`
and `godep` in order to build and test shakers.

### Using Docker and Makefile

You need to run the ``test-unit`` target. 
```bash
$ make test-unit
docker build -t "shakers-dev:master" .
# […]
docker run --rm -it   "shakers-dev:master" ./script/make.sh test-unit
---> Making bundle: test-unit (in .)
+ go test -cover -coverprofile=cover.out .
ok      github.com/vdemeester/shakers   0.015s  coverage: 96.0% of statements

Test success
```

### Using `godep`

The idea behind `godep` is the following :

- when checkout(ing) a project, **run `godep restore`** to install
  (`go get …`) the dependencies in the `GOPATH`.
- if you need another dependency, `go get` it, import and use it in
  the source, and **run `godep save ./...` to save it in
  `Godeps/Godeps.json` and vendoring it in `Godeps/_workspace/src`.

```bash
$ godep restore
$ godep go test -v -cover ./...
=== RUN   Test
OK: 26 passed
--- PASS: Test (0.01s)
PASS
coverage: 96.0% of statements
ok      _/home/vincent/src/github/vdemeester/shakers    0.009s
```
