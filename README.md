# Github API protobuf

This project contains the [protobuf](https://github.com/google/protobuf) definition files for entities and events in the [Github API](https://developer.github.com/).

## Building

You can generate language bindings by running the Makefile with the following targets:

- `protoc-go` Go bindings
- `protoc-python` Python bindings
- `protoc-java` Java bindings
- `protoc-js` JavaScript bindings
- `protoc-ruby` Ruby bindings
- `protoc-php` PHP bindings

Code will be generated to the `build/VERSION/LANGUAGE` such as `build/1.0.0/go`

## Testing

You can run `make test` which will generate `go test` tests and run against the Go bindings to test Github API events from the documentation.

## License

Licensed under the MIT License.
