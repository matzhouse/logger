package logger

import (
	"log"
)

type logmessage struct {
	Message string
	Mtype   string
}

type logger struct {
	c chan logmessage
}

var chanOut, chanErr chan logmessage
var errorReader, outReader bool

func init() {

	chanOut = loggerFactory("out")
	//chanErr = loggerFactory("err")

	errorReader = false
	outReader = false

}

// creates the channels for loggers to send on
func loggerFactory(mtype string) chan logmessage {

	c := make(chan logmessage)

	go func(c chan logmessage) { // We launch the goroutine from inside the function.
		for {
			s := <-c
			log.Println(s)
		}
	}(c)

	return c // Return the channel to the caller.
}

func NewLogger() (l *logger) {

	// set up a logger with a fan in go routine that returns the logger

	l = &logger{}

	l.c = make(chan logmessage)

	// start fan in function
	go func(in, out chan logmessage) {
		for {
			out <- <-in
		}
	}(l.c, chanOut)

	return l

}

func (l logger) Log(message string) {

	lm := logmessage{}

	lm.Message = message

	l.c <- lm

}
