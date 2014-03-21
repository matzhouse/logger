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

## Simple visual

<pre>
write:
node --->  logger  |---->  redis_1
                   |---->  redis_2

read:
node_1 --|  
         |
         |             |---->  redis_1 (node_2 read)
         |--> logger --|
         |             |---->  redis_2 (node_1 read)
node_2 --|
</pre>
