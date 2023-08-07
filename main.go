package main

import (
	"log"
	"math/rand"
	"net"
	"time"
)

var network = "tcp4"
var adress = ":7777"

func main() {
	ln, err := net.Listen(network, adress)
	if err != nil {
		log.Println(err)
	}
	defer ln.Close()

	for {
		c, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}

		go HandleConn(c)
	}

}

func HandleConn(c net.Conn) {

	for {

		time.Sleep(time.Second * 3)

		i := rand.Intn(len(goprv))

		c.Write([]byte(goprv[i] + "\n"))
	}

	//c.Close()
}

var goprv = []string{
	"Don't communicate by sharing memory, share memory by communicating.",
	"Concurrency is not parallelism.",
	"Channels orchestrate; mutexes serialize.",
	"The bigger the interface, the weaker the abstraction.",
	"Make the zero value useful.",
	"interface{} says nothing.",
	"Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.",
	"A little copying is better than a little dependency.",
	"Syscall must always be guarded with build tags.",
	"Cgo must always be guarded with build tags.",
	"Cgo is not Go.",
	"With the unsafe package there are no guarantees.",
	"Clear is better than clever.",
	"Reflection is never clear.",
	"Errors are values.",
	"Don't just check errors, handle them gracefully.",
	"Design the architecture, name the components, document the details.",
	"Documentation is for users.",
	"Don't panic.",
}
