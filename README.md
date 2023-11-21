## Details

```sh
go test -v -p 1 -count 1 ./foo...
```

And compare to

```sh
go test -v p 2 -count 1 ./foo...
```

And compare to

```sh
go test -v -p 10 -count 1 ./foo...
```

## All together

```sh
time go test -p 1 -count 1 ./foo...
time go test -p 2 -count 1 ./foo...
time go test -p 10 -count 1 ./foo...
```

should give

```text
ok  	github.com/cybozu-private/test-go-test/foo	5.104s
ok  	github.com/cybozu-private/test-go-test/foo/bar	5.110s
ok  	github.com/cybozu-private/test-go-test/foo/baz	5.112s
?   	github.com/cybozu-private/test-go-test/foo/utils	[no test files]
go test -p 1 -count 1 ./foo...  0.29s user 0.22s system 3% cpu 16.003 total
ok  	github.com/cybozu-private/test-go-test/foo	5.230s
ok  	github.com/cybozu-private/test-go-test/foo/bar	5.370s
?   	github.com/cybozu-private/test-go-test/foo/utils	[no test files]
ok  	github.com/cybozu-private/test-go-test/foo/baz	5.232s
go test -p 2 -count 1 ./foo...  0.28s user 0.22s system 4% cpu 10.888 total
?   	github.com/cybozu-private/test-go-test/foo/utils	[no test files]
ok  	github.com/cybozu-private/test-go-test/foo	5.093s
ok  	github.com/cybozu-private/test-go-test/foo/bar	5.135s
ok  	github.com/cybozu-private/test-go-test/foo/baz	5.183s
go test -p 10 -count 1 ./foo...  0.33s user 0.30s system 11% cpu 5.447 tota
```