---
title: Optimizations
---

With **Zero Allocations** in the hot-path and an 8-Byte Packet Header, the network overhead of Frisbee is significantly
lower than that of existing protocols like gRPC (performance comparisons available in our [gRPC Benchmarks](/performance/grpc-benchmarks))
which use HTTP/2 under the hood.

This, combined with the substantial performance gains over [protobufs](https://github.com/protocolbuffers/protobuf) that come with using
[polyglot](http://github.com/loopholelabs/polyglot-go) for serialization and deserialization, makes fRPC a great choice for high-performance, high-scalability, and high-reliability RPC frameworks.

We originally designed Frisbee for our own messaging needs at [Loophole Labs](https://loopholelabs.io), where we needed
to send both large and small amounts of data in a latency-sensitive manner. We also needed it to be massively scalable,
able to handle thousands of concurrent connections and able to send millions of messages.

To fulfill these requirements we spent a lot of time architecting the Frisbee data path to be extremely fast and extremely efficient.

# Data Path Optimizations

Our optimizations began with the `Packet` package, which efficiently recycles the byte buffers
that are used throughout Frisbee to hold interstitial data. These make use of the [polyglot](http://github.com/loopholelabs/polyglot-go) library
to implement `Encoding` and `Decoding` packages, which read and write directly from the `Packet.packet`
structure. By recycling the `Packet.packet` structures throughout Frisbee, we can significantly reduce
the number of allocations in the encoding and decoding functions.

Most of our other optimizations center around our network I/O. Actually reading and writing data from a TCP socket
is extremely slow, and so Frisbee makes an effort to maximize the amount of data that we read or write to a TCP socket
while avoiding any additional latency.

All these optimizations - as well as Frisbee's architecture, make it feasible to use Frisbee (as well as fRPC)
in both latency-sensitive applications like in real-time streaming, as well as high-throughput applications like PUB/SUB systems.

# Scheduling Optimizations

By default, fRPC creates a new goroutine for each incoming RPC. This is a very similar approach to the one used by gRPC,
and is a good choice for high-throughput applications where handling the RPC can be a blocking operation (like querying a remote
database).

However, fRPC can also be configured to create a single goroutine to handle all the RPCs from
each incoming connection. This is a good choice for applications that require very low latency and where the handlers are not blocking operations (such as metrics streaming).

<Note>
  In our benchmarks we've tested both approaches, though it should be noted that
  the single-goroutine approach is not as efficient as the multi-goroutine
  approach when the blocking time of the RPC handler is high.
</Note>

# Why TCP?

Many of the recently released wire protocols like [Wireguard](https://www.wireguard.com/) and [QUIC](https://datatracker.ietf.org/doc/html/rfc9000)
use UDP under the hood instead of TCP. Unlike TCP, UDP is an unreliable transport mechanism and provides no guarantees
on packet delivery.

There are benefits to using UDP and implementing packet delivery mechanisms on top of it - QUIC, for example, uses
UDP to solve the [head-of-line blocking](https://calendar.perfplanet.com/2020/head-of-line-blocking-in-quic-and-http-3-the-details/)
problem.

For Frisbee, however, we wanted to make use of the existing performance optimizations that networking software and hardware
have for TCP traffic, and we wanted the strong guarantees around packet delivery that TCP already provides.

<Note>
  It's important to note that while Frisbee and fRPC were designed to be used
  with TCP connections, there's no reason developers can't use other transports.
  As long as the transport fulfills the 'net.Conn' interface, it will work as
  expected with Frisbee and fRPC.
</Note>
