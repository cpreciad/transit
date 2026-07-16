package main

import (
	"github.com/cpreciad/transit/pkg"
    "go.bug.st/serial"
    "log"
    "time"
)

const (
    portName         = "/dev/serial0"
    TICK_INTERVAL    = 25 * time.Second
    TIMEOUT_INTERVAL = 2 * time.Minute
)

type driver struct {
    port serial.Port
}

func main(){
    d := newDriver()
    tick := time.Tick(TICK_INTERVAL)
    timeout := time.Tick(TIMEOUT_INTERVAL)

    log.Print("Good Morning! Today is a fresh day, so stay grateful and count your blessings")

    for {
        select{
        case <-tick:
            data, err := fetch.DisplayDuboceIoT()
            if err != nil{
                log.Printf("fetch error: %v", err)
                d.signalError()
            }

            n, err := d.port.Write(data)
            if err != nil{
                log.Printf("failed to write to %s: %v", portName, err)
                d.signalError()
            }
            log.Printf("bytes written to %s: %d", portName, n)
        case <-timeout:
            log.Printf("shutting down serial port %s...", portName)
            d.signalShutdown()
            log.Print("goodbye!")
            return
        }
    }
}

func newDriver() *driver{
    mode := &serial.Mode{
        BaudRate: 9600,
    }
    port, err := serial.Open(portName, mode)
    if err != nil{
        log.Fatalf("failed to open port %s: %v. shutting down...", portName, err)
        return nil
    }

    return &driver{
        port: port,
    }
}

func (d *driver) signalShutdown(){
    data := []byte("done for the day\ngoodbye!")
    _, err := d.port.Write(data)
    if err != nil{
        log.Printf("failed to signal shutdown to %s: %v", portName, err)
    }
}

func (d *driver) signalError(){
    data := []byte("fetch error...\ncheck logs")
    _, err := d.port.Write(data)
    if err != nil{
        log.Printf("failed to signal error to %s: %v", portName, err)
    }
}
