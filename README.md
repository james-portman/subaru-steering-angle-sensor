# subaru-steering-angle-sensor

Steering angle sensor V416 Tokai Rika Japan

On the underside of the PCB it has an eeprom
93LC66BI
93LC66B really then I for industrial temperature range

Working to fix fault code P1767 on an apparently working steering angle sensor

It seems like once a sensor has been spun too far one way (with the steering rack disconnected) it flags a permanent fault

Installing a new sensor fixes this, but I think it is just a software fault

Looking at the insides of the sensor, it is not easy to actually break these


Changing the EEPROM data a5a5 5a5a back to ffff ffff cleared the fault
DCCD module no longer reporting error

I would think there is a CAN packet that would clear the steering sensor fault, which would be sent by the DCCD module after you doing a normal diagnostic call to the DCCD
