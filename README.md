logger
======

A go library for logging data from anywhere using channels

Every instance of logger comes back to a single channel.

# Possible Uses
Monitoring lots of go routines doing the same thing
Single write to file from lots of outputs

This library is more of a test than anything.
It creates a logging system that passes messages over a Fan-In type channel system.

## Example

```go
lg := logger.NewLogger()

var message string

for i := 50; i >= 0; i-- {

	message = fmt.Sprintf("Logging, number %d", i)

	lg.Log(message)
	time.Sleep(200 * time.Millisecond)
}
```
