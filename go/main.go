package main

import (
  "fmt"
  "os/exec"
  "time"
  "github.com/go-daq/canbus"
)

var ecuPacketsReceived = 0

func printout() {
  for {

    fmt.Print("ECU packets received: ")
    fmt.Println(ecuPacketsReceived)

    time.Sleep(5 * time.Second)
  }
}


// main - loop to receive can packets, sends them off for parsing
func main() {
  go connectLoop()
  go printout()

  for { time.Sleep(100 * time.Second) }
}

func connectLoop() {
  for {
    connect()
    time.Sleep(1 * time.Second)
  }
}

func connect() {
  // hmm
  // /sys/class/net/can0
  cmd := exec.Command("ifconfig", "can0", "down")
  cmd.Start()
  cmd.Wait()
  cmd = exec.Command("ip", "link", "set", "can0", "type", "can", "restart-ms", "100")
  cmd.Start()
  cmd.Wait()
  cmd = exec.Command("ip", "link", "set", "can0", "type", "can", "bitrate", "500000")
  cmd.Start()
  cmd.Wait()
  cmd = exec.Command("ifconfig", "can0", "up")
  cmd.Start()
  cmd.Wait()

  sck, err := canbus.New()
  defer sck.Close()
  err = sck.Bind("can0")
  if err != nil {
    fmt.Println("CAN: there was an error connecting")
    return
  }
  fmt.Println("CAN: seems connected")
  for {
    id, data, err := sck.Recv() // TODO: does this block? some way to flag to UI we're not receiving?
    if err != nil {
      fmt.Print("CAN: ")
      fmt.Println(err)
      return
    } else {
      ecuPacketsReceived++
      intData := convertData(data)
      parse(id, intData)
    }
  }
}

// convert byte array to an int array
func convertData(data []byte) []int {
  output := make([]int, len(data))
  for i := 0 ; i < len(data) ; i++ {
    output[i] = int(data[i])
  }
  return output
}
