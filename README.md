### Protobuf or JSON?

This repository contains 2 equal APIs: gRPC using Protobuf and JSON over HTTP. The goal is to run benchmarks for 2 approaches and compare them. APIs have 1 endpoint to create user, containing validation of request. Request, validation and response are the same in 2 packages, so we're benchmarking only mechanism itself. Benchmarks also include response parsing.

### Requirements

 - Go 1.15

### Run tests

Run benchmarks:
```
go test -bench=. -benchmem
```

### Results

```
goos: darwin
goarch: amd64
pkg: github.com/xans-me/protobuff-json-benchmark
BenchmarkJSONHTTP-12               13203             87254 ns/op            9260 B/op        118 allocs/op
BenchmarkProtobuffGRPC-12          10000            146531 ns/op            9615 B/op        189 allocs/op
PASS
ok      github.com/xans-me/protobuff-json-benchmark     4.924s
```

They are almost the same, HTTP+JSON is a bit faster and has less allocs/op.

### CPU usage comparison

This will create an executable `benchmark-grpc-protobuf-vs-http-json.test` and the profile information will be stored in `grpcprotobuf.cpu` and `httpjson.cpu`:

```
go test -bench=BenchmarkProtobuffGRPC -cpuprofile=protobuffgrpc.cpu
go test -bench=BenchmarkJSONHTTP -cpuprofile=jsonhttp.cpu
```

Check CPU usage per approach using:

```
$ go tool pprof protobuffgrpc.cpu
Type: cpu
Time: Dec 30, 2020 at 1:47am (+07)
Duration: 1.81s, Total samples = 2.32s (128.13%)
```

```
$ go tool pprof jsonhttp.cpu
Type: cpu
Time: Dec 30, 2020 at 1:47am (+07)
Duration: 2.21s, Total samples = 2.94s (133.02%)
```

My results show that Protobuf consumes less ressources.

