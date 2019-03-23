# alarmista-cli
Command line interface for for the "alarmista" alarm clock 

# command examples

##### write a characteristic
```bash
sudo ./alarmista-cli writec -address 30:AE:A4:02:BC:3A -uuid 9501faf3-b697-40de-ad74-0a10f5e2de2c -value 12345
```
##### read a characteristic
```bash
sudo ./alarmista-cli writec -address 30:AE:A4:02:BC:3A -uuid 9501faf3-b697-40de-ad74-0a10f5e2de2c
```

##### send gpio writes to simulate a button on/off
```bash
sudo ./alarmista-cli gpio -pin 12 -value 0
```
