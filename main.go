package main

import (
	"github.com/coreos/go-etcd/etcd"
	"log"
)

func main() {
	c := etcd.NewClient([]string{"http://192.168.59.3:49159"})
	comChannel := make(chan *etcd.Response)
	_, err := c.Watch("/some-message", 0, false, comChannel, nil)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case msg := <-comChannel:
			log.Printf("%s", msg.Node.Key)
		}
	}
}
