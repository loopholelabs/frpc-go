---
title: Concepts
---

fRPC is, at its core, a code generator - one which uses the Frisbee messaging framework as its underlying transport mechanism. It hooks into
the `protoc` compiler and generates an RPC framework that matches the `proto3` spec provided to it.

Frisbee was designed to allow developers to define their own messaging protocols, while having a library that would handle
all the lower level implementation for them.

<Tooltip tip="Request/Reply systems where the request is sent to a remote service which then replies back to the caller">RPC Frameworks</Tooltip>
are implementations of the Request/Reply pattern, and so fRPC generates the necessary
Frisbee code to handle that messaging pattern.

There are three main components to fRPC:

- The Message Types
- The Client
- The Server

# Message Types

One of the challenges with any messaging system is that the messages must be serialized and deserialized into formats that
can be transmitted over the wire. With a code generator like fRPC, that means we need to take your `proto3`
message definitions and generate the accompanying `structs` in Go. We then need to create consistent, performant,
and safe serialization and deserialization functions for those structs.

To do this, fRPC makes use of the [Polyglot](https://github.com/loopholelabs/polyglot-go) library, which is a high-performance
serialization framework that recycles byte buffers and can serialize and deserialize data with almost no allocations.
This makes serialization and deserialization extremely fast, while also allowing us to minimize the accompanying memory allocations.

[polyglot-go](https://github.com/loopholelabs/polyglot-go) library type comes with a number of
`encode` and `decode` methods for various types, that fRPC chains together to create the
serialization and deserialization functions for your `proto3` message definitions.

We're also actively working on a [polyglot-rs](https://github.com/loopholelabs/polyglot-rs) library, which is a Rust
implementation of `Polyglot`, as well as [polyglot-ts](https://github.com/loopholelabs/polyglot-ts) which is a
TypeScript (and Javascript) implementation of `Polyglot`.

# The Client

The fRPC Client is a simple wrapper around the `frisbee.Client` type, and contains generated helper
functions for creating and sending requests to an fRPC Server and then returning the accompanying response.

It's also possible to deviate from those helper functions and access the underlying `frisbee.Client` directly.
This allows you to do things like turn Frisbee off (and thus retrieve the underlying TCP connection).

# The Server

The fRPC Server is a simple wrapper around the `frisbee.Server` type, and contains generated helper
functions for handling incoming requests and returning the accompanying response based on the handlers you've passed in
to the constructor.

Similar to the Client, it's also possible to deviate from those helper functions and access the underlying
`frisbee.Server` directly. This allows you to do things like turn Frisbee off (and thus retrieve the
underlying TCP connection), or write your own middleware functions for incoming or outgoing packets.

# Accessing Frisbee Directly

As we've mentioned before, it's possible to access the underlying [Frisbee](https://github.com/loopholelabs/frisbee-go) primitives from both the
client and the server. This is why fRPC is more flexible than other RPC frameworks, and why it's possible to
do things like send a few RPC requests using fRPC and then reuse that underlying TCP connection for something like an
HTTP proxy.

fRPC generates a `frisbee.HandlerTable` that allows Frisbee to route incoming packets to the correct
handler functions. It's possible to override this table using the `frisbee.Server.SetHandlerTable()`
method (which is exposed in the generated `frpc.Server` type).

To learn more about how [Frisbee](https://github.com/loopholelabs/frisbee-go) works and how you can leverage it from within the generated fRPC
code, check out the [frisbee-go Github Repository](https://github.com/loopholelabs/frisbee-go).
