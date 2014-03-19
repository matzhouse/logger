logger
======

A go library for logging from using channels


This library is more of a test than anything.
It creates a logging system that passes messages over a Fan-In type channel system.

## Example

```
lg := logger.NewLogger()

var message string

for i := 50; i >= 0; i-- {

	message = fmt.Sprintf("Logging, number %d", i)

	lg.Log(message)
	time.Sleep(200 * time.Millisecond)
}
```