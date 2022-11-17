# subaru-steering-angle-sensor - SAS

## Sensor

This applies to at least Hawkeye era steering angle sensors - vaguely 2006 year cars, it could apply to many more

The sensor in the pictures in this repo is marked V416 Tokai Rika Japan, but this seems to work for all sensors

This is not for cars which just have a clock spring, I think it may only apply to the A-DCCD cars


## Fault code

It seems like once a sensor has been spun too far one way it flags a permanent fault, in this case fault code P1767 was showing on an apparently otherwise working steering angle sensor

This can happen if you have the steering rack disconnected and connect it back with it spun 360 degrees or more.
There is gearing and a special pattern inside the sensor so it knows if you go "too far" one way, it might only trigger when you put full lock of steering on, or the car may work it out from the wheel speed sensor data when it sees you are going straight with the wheels but have 360 degrees of steering wheel on

This might be a good thing if the car had been in a crash, so you would know something was extremely bent or damaged for that to happen, but not if it was just an accident after disconneting the steering knuckle

Installing a new sensor fixes this, but it is just a software issue that can be fixed

## Test steering angle sensor

If you use FreeSSM and connect to the DCCD controller you can test your sensor.
Initially it is normal for the steering sensor value to be strange value like 0 or 4096(?), then once you move the steering wheel it should start to read real values.
If you can turn the wheel and it seems to get sensible readings that follow with how you are moving the wheel then it is likely that your steering sensor is working fine.
In my case there would not be any fault until I moved the wheel.

## Technical investigation

In technical terms I intercepted the CAN data from the steering sensor to the DCCD controller. Normally the steering sensor sends packets which have the steering angle in them. I can't quite remember now but when the permanent fault has been set, one of the can packet data bytes had an extra bit set as a fault flag.
I set up dummy CAN messages and sent them to the DCCD without that specific bit set, and this would never cause a fault on the DCCD, but still showed the steering sensor angle
This in iteself did not fix the issue but at least it seemed that the sensor itself knew there was a fault somehow, and was poisoning the data sent to the DCCD

I opened up the sensor to have a look insides, it is not easy to actually break these which I have seen said numerous times on the internet, I really doubt you would break it from dropping it etc

On the underside of the PCB it has an eeprom chip marked 93LC66BI - it is a version of the widely used 93C66 chip really

Changing the EEPROM data section near the top of the file from a5a5 5a5a back to ffff ffff cleared the fault, the DCCD module is no longer reporting an error

Other sensors people have sent me to repair have had slightly different fault data in, but it is always a pattern of 2 hex byte repeated, then another 2 of the data bytes with the upper and lower 4 bits swapped, so 5a 5a a5 a5 or 4b 4b b4 b4 for example
Wiping these by setting them back to all ff's always seems to work

I would think there might be a way to clear this fault by using a diagnostic cable and sending a certain message to the DCCD module, but I have never seen or heard of a tool that can do this, and Subaru themselves claimed that it is not possible to reset the steering sensors in any way.

## Further investigation

The rest of the data on the steering angle sensor eeprom seems to be directly related to the actual steering angle

It seems like there are two lookup tables which I assume take the current steering wheel angle and convert it into actual car wheel angle

I have seen people making complicated add-on systems to correct the steering angle sensor data when they have fitted a steering rack with a different ratio, I would think you can do the same thing here by just editing the steering angle sensor maps to match a sensor from whatever car your rack is made for
