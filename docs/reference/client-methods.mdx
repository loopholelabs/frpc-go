---
title: Client Methods
---

# NewClient

The `NewClient(tlsConfig *tls.Config, logger *zerolog.Logger) (*Client, error)` method is used to create a new fRPC client.

It takes two arguments:

- `tlsConfig`: a `*tls.Config` that will be used to configure TLS for the underlying connection. This can be left as `nil` if no TLS is required.
- `logger`: a `*zerolog.Logger` that will be used to log all events. This can be left as `nil` if no logging is required.

It returns an fRPC `*Client` on success and an error otherwise.

# Connect

The `Connect(addr string) error` method is used to initiate a connection to an fRPC server.
If a `*tls.Config` was provided when the client was created (using `NewClient`)
it will be used to configure TLS for the connection.

An error will be returned if the connection fails.

<Note>
  The Connect function should only be called once. If FromConn was already
  called on this client, Connect will return an error.
</Note>

# FromConn

The `FromConn(conn net.Conn) error` method is used to create a new fRPC client from an existing net.Conn. This is useful if you want to reuse an existing connection, or
if you have a custom transport that you want to use. If a `*tls.Config` was provided when the client was created (using `NewClient`), it will be ignored.

An error will be returned if the connection fails.

<Note>
  The FromConn function should only be called once. If Connect was already
  called on this client, FromConn will return an error."
</Note>

# Closed

The `Closed() bool` method is used to check if the client is closed. This method will return `true` if the client is closed or has not yet been initialized, and `false` otherwise.

# Error

The `Error() error` method is used to check if the client has encountered an error. This method will return an `error` if the client has encountered an error, or `nil` otherwise.

This method is meant to be used to check if the client encountered an error that caused it to close.

# Close

The `Close() error` method is used to close the client. It will return an `error` if the client encounters an error while closing (or if it is already closed), and will cancel any pending RPCs.

# WritePacket

The `WritePacket(p *packet.Packet) error` method is used to write a raw frisbee packet to the underlying connection. Normal fRPC operations should not use this method, however it is available
when extending fRPC with custom protocols or messaging patterns directly for use with the underlying Frisbee library.

# Flush

The `Flush() error` method is used to flush the underlying connection. Normal fRPC operations should not use this method, however it is available
when extending fRPC with custom protocols or messaging patterns directly for use with the underlying Frisbee library.

# CloseChannel

The `CloseChannel() <- chan struct{}` method is used to signal to a listener that the Client has been closed.
The returned channel will be closed when the client is closed, and the `Error()` method can be used to check if the connection was closed due to an error.

# Raw

The `Raw() (net.Conn, error)` method is used to get the underlying `net.Conn` from the fRPC Client. This is useful if you want to extend fRPC with custom protocols or messaging patterns directly for use with the underlying Frisbee library.

# Logger

The `Logger() *zerolog.Logger` method is used to get the logger that was provided when the client was created. This is useful if you want to extend fRPC with custom protocols or messaging patterns directly for use with the underlying Frisbee library.

# Generated Methods

When generating the fRPC Client, each service in the `.proto` file also results in a generated `service` Client.
Then, for each RPC in the service, a method is generated on the appropriate the `service` client.
For example, if the `.proto` file contains the following service definition:

```proto
service MyService {
  rpc MyMethod(MyRequest) returns (MyResponse) {}
}
```

Then the generated `service` Client method is:

```go
func (c *MyService) MyMethod(ctx context.Context, req *MyRequest) (res *MyResponse, err error) {
  ...
}
```

And it's meant to be invoked using:

```go
res, err := c.MyService.MyMethod(ctx, req)
```

Unlike gRPC, where each service requires creating a new RPC Client, fRPC creates a single client for all your services.
