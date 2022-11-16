package main

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/gwangyi/udp-forward"
)

var dst1 = func () *net.UDPAddr {
    addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:5555")
    if err != nil {
        panic(err)
    }
    return addr
}()

var dst2 = func () *net.UDPAddr {
    addr, err := net.ResolveUDPAddr("udp", "127.0.0.1:5556")
    if err != nil {
        panic(err)
    }
    return addr
}()

func router(incoming *net.UDPAddr) *net.UDPAddr {
    fmt.Println("Routing: ", incoming)
    if incoming.Port % 2 == 0 {
        return dst1
    } else {
        return dst2
    }
}

func main() {
    f, err := forward.Forward(forward.WithDestination("127.0.0.1:5555"))
    if err != nil {
        panic(err)
    }
    fmt.Println("Listening: ", f.LocalAddr())
    done := make(chan os.Signal, 1)
    signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
    fmt.Println("Blocking, press ctrl+c to continue...")
    <-done
    f.Close()
}
