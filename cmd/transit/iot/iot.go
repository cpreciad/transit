package main

import (
	"github.com/cpreciad/transit/pkg"
    "go.bug.st/serial"
    "log"
    "time"
)

const (
    portName = "/dev/serial0"
    //portName = "/dev/ttyS0"
    //portName = "/dev/ttyAMA0"
)

type driver struct {
    port serial.Port
}

func main(){
    d := newDriver()
    tick := time.Tick(1 * time.Minute)

    for next := range tick{
        str := fetch.DisplayDuboceIoT()
        n, err := d.port.Write([]byte(str))
        if err != nil{
            log.Fatal(err)
        }
        log.Printf("%v bytes written to %s: %d", next, portName, n)
    }

}

func newDriver() *driver{
    mode := &serial.Mode{
        BaudRate: 9600,
    }
    port, err := serial.Open(portName, mode)
    if err != nil{
        log.Fatal(err)
        return nil
    }

    return &driver{
        port: port,
    }
    
}
