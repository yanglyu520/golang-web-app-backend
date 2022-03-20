# net package

## listen func

- Listen announces on the local network address.
- The network must be "tcp", "tcp4", "tcp6", "unix" or "unixpacket".

```go
func Listen(network, address string) (Listener, error)
```

## type listen

Func Accept will return a connection, and we need to close the listener after all func is done, we can also return the listener's network address.

```go
type Listener interface {
	// Accept waits for and returns the next connection to the listener.
	Accept() (Conn, error)

	// Close closes the listener.
	// Any blocked Accept operations will be unblocked and return errors.
	Close() error

	// Addr returns the listener's network address.
	Addr() Addr
}
```
- A Listener is a generic network listener for stream-oriented protocols.
- Multiple goroutines may invoke methods on a Listener simultaneously.

## type Conn

the conn interface has inherited writer and reader interfaces

```go
type Conn interface {
	// Read reads data from the connection.
	Read(b []byte) (n int, err error)

	// Write writes data to the connection.
	Write(b []byte) (n int, err error)

	// Close closes the connection.
	Close() error

	// LocalAddr returns the local network address, if known.
	LocalAddr() Addr
...

```


## note
run `telnet <IP ADDRESS OF SERVER PC> <PORT> or telnet localhost 8080` to start the web server on the tcp level
