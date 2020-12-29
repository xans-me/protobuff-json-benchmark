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
BenchmarkJSONHTTP-4                 5391            232441 ns/op            9045 B/op        117 allocs/op
BenchmarkProtobuffGRPC-4            5277            237833 ns/op            9401 B/op        190 allocs/op
PASS
ok      github.com/xans-me/protobuff-json-benchmark     6.567s
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
go tool pprof protobuffgrpc.cpu
go tool pprof jsonhttp.cpu
```

My results show that Protobuf consumes less ressources, around **30% less**.

