package main

import (
	"fmt"
  "time"
  "github.com/go-daq/canbus"
)

const (
	frameSep  = "#"
	maxUint32 = uint64(^uint32(0))
)

func sendSASData(sck *canbus.Socket) {
  // looking at the proxy/scaler code
  // either byte 0 or 1
  // the lower 4 bits are bitfield/status
  // on startup the reading is 800C
  // all other valid values when moving wheel end with E
  // 0xC = 0b1100
  // 0xE = 0b1110
  // so bit 0 doesn't change
  // bit 1 is for whether or not we know the position/wheel has moved a bit
  // bit 2 and 3 always 1
  // on "broken" SAS the value is 0x8008
  // 0x8 = 0b1000
  // so is bit 2 being 0 indicating a fault?
  // the SAS has an 8 pin eeprom on the back of the PCB
  // it must be storing the fact that there has been a fault in the EEPROM
  // there is very likely a CAN packet that would clear the fault, otherwise I will make changes to EEPROM until the fault clears
  sequence := 0
  id := 0x002
  // data := []byte{0,0,0,0,0,0,0,0}
  // data := []byte{8,128,0,0,0,0,0,0}
  // data := []byte{8,128,0,0,0,0,0,1} // this is what the broken sensor sends before moving it
  data := []byte{0x0C,0x80,0,0,0,0,0,1} // this is apparently what a working sensor reads before being moved at all
  // data := []byte{0,0,0,0,0,0,0,1}
  // data := []byte{0,0,0,0,0,0,0,1}
  // angle := 10 // degrees
  // rawValue := angle * 16
  // data[0] = byte(rawValue & 0xFF)
  // data[1] = byte((rawValue >> 8) & 0xFF)

  for {
    data[2] = byte(sequence & 0xF)
    checksum := calculateChecksumBytes(data)
    data[2] = data[2] | (checksum << 4)
    // fmt.Println(data)

  	_, err := sck.Send(uint32(id), data)
  	if err != nil {
  		fmt.Println("error sending data")
      fmt.Println(err)
  	} else {
      fmt.Println("Sent SAS data 0x002")
    }

    sequence++
    if sequence > 0xF {
      sequence = 0
    }
    time.Sleep(10 * time.Millisecond)
  }

}

func sendAbs513(sck *canbus.Socket) {
  // VDC/ABS CAN-ID 0x513
  // Sent every 20 ms = 1/50 second.
  //
  // Byte# (0-7)	Description
  // 0 & 1	Wheel Speed Front Left, uint16, little endian (x*0.05625 [km/h])
  // lowest non-zero value seen: 0x0022 = 1.9 km/h
  // 2 & 3	Wheel Speed Front Right
  // 4 & 5	Wheel Speed Rear Left
  // 6 & 7	Wheel Speed Rear Right
  id := 0x513
  data := []byte{0,0,0,0,0,0,0,0}

  for {
    _, err := sck.Send(uint32(id), data)
    if err != nil {
      fmt.Println("error sending data")
      fmt.Println(err)
    } else {
      fmt.Println("Sent ABS data 0x513")
    }
    time.Sleep(20 * time.Millisecond)
  }

}

func sendAbs501(sck *canbus.Socket) {
  // VDC/ABS CAN-ID 0x501
  // Sent every 20 ms = 1/50 second.
  //
  // Byte# (0-7)	Description
  // 0
  // 1
  // 2	Torque Reduction ??? (x*1.6 [Nm])
  // 3	Torque Allowed ??? (x*1.6 [Nm])
  // 4	bit 2: ???
  // bit 1: ???
  // bit 0: Request Torque Down
  // 5	msg counter
  // 6
  // 7
  id := 0x501
  data := []byte{0,0,0,0,0,0,0,0}
  counter := 0

  for {
    data[5] = byte(counter & 0xFF)
    _, err := sck.Send(uint32(id), data)
    if err != nil {
      fmt.Println("error sending data")
      fmt.Println(err)
    } else {
      fmt.Println("Sent ABS data 0x501")
    }
    counter++
    if counter > 0xFF {
      counter = 0
    }
    time.Sleep(20 * time.Millisecond)
  }

}
