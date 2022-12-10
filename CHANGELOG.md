# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/), and this project adheres
to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [v0.7.2] - 2022-12-10

### Fixes

- Fixed a bug where stream handlers would be generated for a proto file without streams

## [v0.7.1] - 2022-12-10

### Changes

- Fixed a bug when generating fRPC with streams where sometimes stream messages would be received out of order.
- Removed the Trunk linter

## [v0.7.0] - 2022-09-28

### Features

- fRPC now uses the `VarInt` encoding format under the hood (added in [polyglot-go v0.5.0](https://github.com/loopholelabs/polyglot-go)) which should help reduce the number of bytes an RPC call is serialized to
- A new `CloseError` type has been added which, when returned by an RPC call, causes the connection to be closed after the message is written. This can be useful for authentication or connection management.
- Streaming is now available! The API matches gRPC's so it should be a drop-in replacement!

### Changes

- The [polyglot-go](https://github.com/loopholelabs/polyglot-go) dependency has been bumped to `v0.5.0`
- The [frisbee-go](https://github.com/loopholelabs/frisbee-go) dependency has been b umped to `v0.7.0`

## [v0.6.0] - 2022-08-24 (Beta)

### Changes

- Refactoring the generated code to use the [polyglot-go](https://github.com/loopholelabs/polyglot-go) library to generate message encode/decode functions. ([#3](https://github.com/loopholelabs/frpc-go/pull/3))

### Fixes

- Fixed an issue with the generated code that caused compilation issues when the names of two methods in different services
  were the same ([#5](https://github.com/loopholelabs/frpc-go/issues/5))

## [v0.5.1] - 2022-07-20 (Beta)

### Fixes

- Fixed an issue where if the number of services is 0 the RPC Generator would
  crash ([#101](https://github.com/loopholelabs/frisbee-go/issues/101))

> Changelogs for [v0.5.0] and before can be found at https://github.com/loopholelabs/frisbee-go

[unreleased]: https://github.com/loopholelabs/frpc-go/compare/v0.7.2...HEAD
[v0.7.2]: https://github.com/loopholelabs/frpc-go/releases/tag/v0.7.2
[v0.7.1]: https://github.com/loopholelabs/frpc-go/releases/tag/v0.7.1
[v0.7.0]: https://github.com/loopholelabs/frpc-go/releases/tag/v0.7.0
[v0.6.0]: https://github.com/loopholelabs/frpc-go/releases/tag/v0.6.0
[v0.5.1]: https://github.com/loopholelabs/frpc-go/releases/tag/v0.5.1
[v0.5.0]: https://github.com/loopholelabs/frisbee-go/compare/v0.4.6...v0.5.0
