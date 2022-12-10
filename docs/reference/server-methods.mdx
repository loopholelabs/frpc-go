---
title: Server Methods
---

# NewServer

The `NewServer(... Services, *tls.Config, *zerolog.Logger) (*Server, error)` method is used to create a new fRPC Server. It takes a list of structs that implement
the RPC methods corresponding to the services defined in the `.proto` file.

For example, if the `.proto` file contains the following service definition:

```proto
service MyService {
  rpc MyMethod(MyRequest) returns (MyResponse) {}
}

service OtherService {
  rpc OtherMethod(MyRequest) returns (MyResponse) {}
}
```

Then the generated `service implementation ` interfaces are expected:

```go
type MyService interface {
	MyMethod(context.Context, *MyRequest) (*MyResponse, error)
}

type OtherService interface {
    OtherMethod(context.Context, *MyRequest) (*MyResponse, error)
}
```

And the corresponding generated `NewServer` method would be:

```go
func NewServer(myService MyService, otherService OtherService, tlsConfig *tls.Config, logger *zerolog.Logger) (*Server, error) {
    ...
}
```

The generated `NewServer` method also takes the additional two arguments:

- `tlsConfig`: a `*tls.Config` that will be used to configure TLS for fRPC server. This can be left as `nil` if no TLS is required.
- `logger`: a `*zerolog.Logger` that will be used to log all events. This can be left as `nil` if no logging is required.

It returns an fRPC `*Server` on success and an error otherwise.

<Note>
  For long-running RPC handlers it's important to pay attention to the passed on
  context.Context - it will be cancelled when the server is shutdown and
  handlers are expected to return as soon as that happens.
</Note>

# SetBaseContext

The `SetBaseContext(f func() context.Context) error` method is used to set the base context for all incoming RPCs. This is useful if you want to set a common context for all incoming RPCs.

# SetOnClosed

The `SetOnClosed(f func(*frisbee.Async, error)) error` method is used to set the callback function that will be called when a connection to an fRPC client is closed. This is useful if you want to do any cleanup when a connection to a client is closed.

# SetHandlerTable

The `SetHandlerTable(handlerTable frisbee.HandlerTable) error` method is used to set the handler table for the fRPC server. This is useful if you want to set a custom handler table for the fRPC server, and is commonly used to extend the fRPC server with custom handlers for
alternative messaging patterns. In order to avoid breaking the fRPC functionality, it's recommended to first use the `GetHandlerTable` method to retrieve the base handler table, modify it, and then use the `SetHandlerTable` method to set the modified handler table.

# GetHandlerTable

The `GetHandlerTable() frisbee.HandlerTable` method is used to retrieve the handler table for the fRPC server. This is useful if you want to retrieve and extend handler table for the fRPC server, and is commonly used with the `SetHandlerTable` method.

# SetPreWrite

The `SetPreWrite(f func()) error` method is used to set the pre-write callback function for the fRPC server. This is useful if you want to handle metrics or do some logging before a request is written to the client.

# SetConcurrency

The `SetConcurrency(concurrency uint64)` method is used to set the concurrency for the fRPC server. This is useful if you want to set a maximum number of concurrent goroutines that can be spawned by the fRPC server (across all clients).

Setting this value to `0` will result in the fRPC server spawning an unlimited number of goroutines to handle incoming RPCs, and setting the value to `1` will result in the fRPC server spawning a single goroutine (per fRPC Client) to handle incoming RPCs.

All other values will result in the fRPC server spawning (at maximum) the specified number of goroutines to handle incoming RPCs.

# Start

The `Start(addr string) error` method is used to start the fRPC server. It takes the address to listen on as an argument and will return an error if the server fails to start.

# ServeConn

The `ServeConn(conn net.Conn)` method is used to serve a given net.Conn. It takes the connection to serve as an argument and is a non-blocking method - it will return immediately after starting to serve the connection in a separate goroutine, and if it
encounters an error, the `OnClosed` callback function will be called with the error.

# Logger

The `Logger() *zerolog.Logger` method is used to retrieve the logger for the fRPC server. This is useful if you want to retrieve and extend the fRPC server.

# Shutdown

The `Shutdown() error` method is used to shutdown the fRPC server. It will return an error if the server fails to shutdown, and it will clean up all goroutines spawned by the server before returning. Any active connections will be closed before the server is shutdown, and any active RPCs
will be cancelled. The contexts given to the RPCs will be cancelled as well.

# Generated Interfaces

When generating the fRPC Server, each service in the `.proto` file requires a `service implementation` that
fulfills the RPC methods defined for the service.

For example, if the `.proto` file contains the following service definition:

```proto
service MyService {
  rpc MyMethod(MyRequest) returns (MyResponse) {}
}
```

then the generated `service implementation ` interface looks like this:

```go
type MyService interface {
	MyMethod(context.Context, *MyRequest) (*MyResponse, error)
}
```

This is a similar function signature to the one gRPC would generate, making it easy to reuse the service implementation from gRPC.
