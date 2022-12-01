---
title: Architecture
---

The architecture of fRPC is based on the standard Server/Client model that a lot of other RPC frameworks are follow.
The idea is that the `Client` makes a connection with the `Server`, and then sends a structured
request. Based on the request type, the `Server` runs a handler that then returns a response or an error.
The `Server` then forwards that response object (or the error) back to the `Client`.

From the perspective of the `Client`, they have simply called a function and received a response. The fact
that the request is serialized and transmitted to the `Server` is hidden for simplicity.

To dig into how the underlying architecture of both the `Server` and `Client` work, it is first
important to understand that the underlying [Frisbee](https://github.com/loopholelabs/frisbee-go) protocol does not have any notion of a request
or response. When Frisbee sends a `Packet` of data, it does not wait for a response. This makes
the protocol suitable for a number of use cases (like real-time streaming), but also means that Request/Reply semantics
need to be implemented in the application logic - in this case, the code that fRPC generates.

# Server Architecture

The generated fRPC `Server` is based on the RPC `Services` that are defined in the `proto3`
file that is passed to the `protoc` compiler. Developers are responsible for implementing the generated
`Service` interfaces and passing that into the `Server` constructor.

The `Server` then takes that implementation and creates a `handler table` that maps the request type to the
accompanying function in the provided `Service` implementation.

When it receives a request, it looks up the request type in the `handler table` and calls the accompanying
function with the deserialized Request object. The function then returns a Response object that
is serialized and sent back to the `Client`.

# Client Architecture

The generated fRPC `Client` is also based on the RPC `Services` that are defined in the
`proto3` file that is passed to the `protoc` compiler. Based on the RPC Calls defined in those services,
fRPC generates a number of `Client` helper functions - one for each possible RPC Call.

As mentioned before, Frisbee does not have any notion of a request or response - this means that we must implement
the ability to wait for a response from the `Server` in the application logic. We need to also be able to map
those incoming responses to the correct ongoing request.

To achieve this, fRPC Clients make use of an `in-flight` requests table that maps a request ID to a channel
that can be listened to for a response. When an RPC function is called, it generates a request ID, serializes the request
object and sends it to the `Server`. When a response object is received from the `Server`, it is
deserialized and request ID is looked up in the `in-flight` requests table.

The response is then pushed into the channel associated with the request ID, where it is read by the RPC function
that made the request in the first place. This response unblocks the RPC caller and the response object (or an error)
is returned.
