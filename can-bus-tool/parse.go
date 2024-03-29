package main

import (
  "fmt"
)

func parse(id uint32, data []int) {

  switch(id) {

    // DCCD sends, periodic/keepalive
    case 0x428:
      // the first and last numbers increase then roll to 0
      // they are not (always?) lined up
      fmt.Println("0x428 packet - DCCD alive stream")
      fmt.Println(data)
      break

      // steering angle sensor, sending current sensor angle and status
    case 0x002:
      // -5830 for roughly a full turn
      // 5706 the other way full turn
      // /16 to get degrees
      // 360 degrees = 5760
      // also correlates with initial startup value of 0x8000 which shows as 2048(d) in freessm
      fmt.Println("0x002 - Steering angle sensor - steering angle")
      fmt.Println(data)

      steeringAngle := (data[1] << 8) | data[0]
      // sign it
      if steeringAngle & 0b1000000000000000 > 0 {
        steeringAngle = steeringAngle - 0xffff
      }
      fmt.Println(steeringAngle)

      fmt.Print("Steering angle in degrees ")
      fmt.Println(steeringAngle/16)

      checksum := data[2] >> 4
      fmt.Print("checksum ")
      fmt.Println(checksum)

      sequence := data[2] & 0b1111
      fmt.Print("sequence ")
      fmt.Println(sequence)

      calc_checksum := calculateChecksum(data)
      fmt.Print("calculated checksum ")
      fmt.Println(calc_checksum)

      if checksum != calc_checksum {
        fmt.Println("Checksum looks wrong !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
      } else {
        fmt.Println("Checksum looks good")
      }
      break


    default:
      fmt.Print("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!! UNHANDLED PACKET, ID: ")
      fmt.Println(id)
      fmt.Println(data)
      break
  }

}

func calculateChecksum(data []int) int {
    calc_checksum := (data[0] & 0xF)
    calc_checksum ^= (data[0] >> 4)
    calc_checksum ^= (data[1] & 0xF)
    calc_checksum ^= (data[1] >> 4)
    calc_checksum ^= (data[2] & 0xF) // don't do the checksum nibble too
    calc_checksum ^= (data[3] & 0xF)
    calc_checksum ^= (data[3] >> 4)
    calc_checksum ^= (data[4] & 0xF)
    calc_checksum ^= (data[4] >> 4)
    calc_checksum ^= (data[5] & 0xF)
    calc_checksum ^= (data[5] >> 4)
    calc_checksum ^= (data[6] & 0xF)
    calc_checksum ^= (data[6] >> 4)
    return calc_checksum
}
func calculateChecksumBytes(data []byte) byte {
    calc_checksum := (data[0] & 0xF)
    calc_checksum ^= (data[0] >> 4)
    calc_checksum ^= (data[1] & 0xF)
    calc_checksum ^= (data[1] >> 4)
    calc_checksum ^= (data[2] & 0xF) // don't do the checksum nibble too
    calc_checksum ^= (data[3] & 0xF)
    calc_checksum ^= (data[3] >> 4)
    calc_checksum ^= (data[4] & 0xF)
    calc_checksum ^= (data[4] >> 4)
    calc_checksum ^= (data[5] & 0xF)
    calc_checksum ^= (data[5] >> 4)
    calc_checksum ^= (data[6] & 0xF)
    calc_checksum ^= (data[6] >> 4)
    return calc_checksum & 0xF
}
