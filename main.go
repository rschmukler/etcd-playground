package main

import (
  "log"
  "github.com/coreos/go-etcd/etcd"
)

func main() {
  machines := []string{"192.168.59.3:49159"}
  c := etcd.NewClient(machines)
  _, err := c.Get("some-message", false, false)
  if err != nil {
    log.Fatal(err)
  }
  comChannel := make(chan *etcd.Response)
  c.Watch("/some-message", 0, false, comChannel, nil)
  for {
    select {
      case msg := <- comChannel:
        log.Printf("%s", msg.Node.Key)
        break
    }
  }
}
