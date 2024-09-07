# go_send_tzsp
Capture packets, tzsp encapsulate them, and send them to another host for storage or parsing

Problem: I needed something running on a remote low resource (low disk/storage) vps or iot device that would capture packets. 

There are multiple ways to solve my problem but I wanted to have this as an option. 

Inspired by Scratch-n-Sniff - https://github.com/nickvsnetworking/Scratch-n-Sniff, NoStarch BlackHat Go - https://nostarch.com/blackhatgo, and Greynoise Community Slack. ❤️

I picked go because I can't write code in C/C++ and I can't put python on capture source/remote device. 

<img width="678" alt="image" src="https://github.com/user-attachments/assets/a1c444d8-6c46-46b0-8fa3-40a635b65231">

# Requirements
- go (I'm using go1.23.1 on my host)
- gopacket
- libpcap-dev
- gcc

# Building
- Clone the repo
- Cd into the repo folder
- run `go mod init go_send_tzsp`
- run `go mod tidy`
- run `go build .`

# Usage
- On the server where you're capturing/storing/parsing pcaps, you can run either tcpdump or wireshark.
- In wireshark/tcpdump, make sure to filter and only capture on udp port 37008 or whatever port you're using as dst port. Wireshark on port 37008 should automatically recognize tzsp and parse accordingly.
- On the VPS/honeypot/IOT device, run `./go_send_tzsp -iface eth0 -filter 'port 8080' -dstip 10.0.0.100`

- go_send_tzsp command line options
```
iface - Interface - required - example: eth0
filter - Filter - required - example: 'port 8080'
dstip - Destination IP of capture/parse server - required (where you're running wireshark/tcpdump)
dstport - Destination port of capture/parse server - default 37008/udp
```

# Screenshots
go_send_tzsp running on 10.0.0.53. 10.0.0.53 is also running web server on 8000
<img width="855" alt="image" src="https://github.com/user-attachments/assets/aec73322-d59a-43a4-84c6-8039989f08d4">

Destination capture/parsing host 10.0.0.8 running wireshark. Its showing http request being sent from 10.0.0.61 to 10.0.0.53:8000
<img width="1439" alt="image" src="https://github.com/user-attachments/assets/53511f41-d651-4165-87f1-57739f0f7bed">


