# subaru-steering-angle-sensor

Steering angle sensor V416 Tokai Rika Japan

On the underside of the PCB it has an eeprom 93LC66BI - it is a version of 93C66 really

Working to fix fault code P1767 on an apparently otherwise working steering angle sensor

It seems like once a sensor has been spun too far one way (e.g. if you have the steering rack disconnected) it flags a permanent fault

This might be a good thing if the car had been in a crash, so you would know something was extremely bent or damaged for that to happen.


Installing a new sensor fixes this, but it is just a software fault

Looking at the insides of the sensor, it is not easy to actually break these which I have seen said numerous times on the internet


Changing the EEPROM data section near the top of the file from a5a5 5a5a back to ffff ffff cleared the fault, the DCCD module is no longer reporting an error

I would think there might be a way to clear this fault by using a diagnostic cable and sending a certain message to the DCCD module, but I have never seen or heard of a tool that can do this, and Subaru themselves claimed that it is not possible to reset the steering sensors in any way.
