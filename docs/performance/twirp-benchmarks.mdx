---
title: Twirp Benchmarks
---

As with our other benchmarks, this will be as close to an apples-to-apples comparison as possible.
These benchmarks are publicly available at [https://github.com/loopholelabs/frpc-go-benchmarks](https://github.com/loopholelabs/frpc-go-benchmarks), and we encourage you to run them for yourselves.

To make sure our benchmark is fair, we'll be using the exact same proto3 file as the input for both fRPC and Twirp.
Moreover, we'll be using **the exact same service implementation for both the Twirp and fRPC** servers - Twirp uses protobufs for serialization and its interface is very similar to gRPC. Because fRPC was designed
to feel familiar to the gRPC interface, using the same service implementation is extremely simple.

```protobuf Benchmark Proto3 File
option go_package = "/";

service BenchmarkService {
  rpc Benchmark(Request) returns (Response);
}

message Request {
  string Message = 1;
}

message Response{
  string Message = 1;
}
```

We'll be running a number of different benchmarks with an increasing number of concurrent clients to show off both the throughput and the scalability of fRPC when compared to Twirp.

The following benchmarks were performed on a bare metal host running Debian 11, 2x AMD EPYC 7643 48-Core CPUs, and 256GB of DDR4 memory. The benchmarks were performed over a local network to avoid inconsistencies due to network latency.

# Client Throughput

For our first set of benchmarks we'll have a number of concurrently connected clients and each client will make RPCs to the fRPC or
Twirp server using a randomly generated fixed-sized message, and then wait for a response before repeating.

In each of our benchmark runs we're increasing the number of concurrently connected clients and we're measuring the average throughput of each client to see how well fRPC and Twirp scale. We're also
running a number of separate benchmarks, each with an increasing message size.

## 32-Byte Messages

![32-byte messages](/images/twirp/32byte.png)

Starting with **32-byte** messages, the results are very similar to the ones with gRPC. fRPC consistently and
substantially outperforms Twirp - though the performance drop off for Twirp is significantly steeper than what we saw with gRPC.

In the case of 8192 connected clients, Twirp's performance drops to only 4 RPCs/second per client while fRPC is able to handle 112 RPC/second.
This means fRPC is 28x more performant than Twirp.

Twirp seems to be relatively capable when there is a small number of connected clients, but quickly falls off as the number of clients increases.

## 512-Byte Messages

![512-byte messages](/images/twirp/512byte.png)

When changing the message size to **512-bytes**, we can see an extremely sharp drop in Twirp's throughput, while fRPC seems to fare
much better. In the case of 8192 connected clients, fRPC's throughput is still 98 RPCs/second while Twirp drops to only 4 - meaning
fRPC performs almost 25x better than Twirp.

## 128-KB Messages

![128-KB messages](/images/twirp/128kb.png)

With larger **128-KB** messages, we continue to see the same pattern as before - throughput for individual clients of both frameworks
drops as more clients are added, but fRPC performs far better than Twirp - in this case between 2-7x better

## 1-MB Messages

![1-MB messages](/images/twirp/1mb.png)

With the largest message size of the benchmark, **1MB**, the pattern from our previous runs continues. In this case, fRPC
seems to perform between 3-6x better than Twirp and we're guessing that our bare metal host is starting the become the bottleneck as
we increase the number of clients.

# Server Throughput

Now let's look at how fRPC servers scale compared to Twirp as we increase the number of connected clients. Twirp makes
use of the standard `net/http` server so we're really comparing against that.

For this benchmark, we're going to make it so that each client repeatedly sends 10 concurrent RPCs in order to
saturate the underlying TCP connections and the accompanying RPC server.

![server throughput](/images/twirp/throughput.png)

As before, we can see that fRPC consistently outperforms Twirp - but as we increase the number of clients beyond 1024,
we actually saw the Twirp clients begin to fail. We couldn't get our benchmark to run for more than 1024 clients, which is
why the benchmark reports a 0 for those runs.

At 1024 clients, though, the fRPC is easily able to handle more than 60x more RPCs/second than Twirp is.

These benchmarks show off just a small portion of fRPCs capabilities, and we encourage everyone to run
these for themselves. We also have [benchmarks comparing fRPCs messaging format with protobuf and other serialization frameworks](https://github.com/loopholelabs/polyglot-go-benchmarks).
