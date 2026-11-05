I tested everything with ubuntu 26.10 or 26.04. It seems to work but tests were short.
In my testing, i had multiple hosts sending tzsp packets to tzsp2pcap and it seemed to work just fine. 


You need software that will capture packets and encapsulate them, like the one in this repo.

It can be ran as a service: https://github.com/BoredHackerBlog/go_send_tzsp/blob/main/service.md

Just make sure you not accidentlly sniff the tzsp port 37008 or whatever you use.

Next you need tzsp2pcap: https://github.com/thefloweringash/tzsp2pcap

I think I just installed gcc & libpcap-dev and ran make. copy the binary into /usr/bin. 

This also can be ran as a service.

Install tcpreplay. It'll replay packets that tzsp2pcap converted to pcap.

Create a dummy interface

file /etc/netplan/01-netcfg.yaml

```
network:
  version: 2
  renderer: networkd
  dummy-devices:
    tzsp0: {}
```

create a service

file /etc/systemd/system/tzspreplay.service 

```
[Unit]
Description=replay TZSP packets
After=network.target

[Service]
Type=simple
User=root
ExecStart=/bin/sh -c '/usr/bin/tzsp2pcap -f | /usr/bin/tcpreplay --topspeed -i tzsp0 -'

Restart=always
RestartSec=5s

[Install]
WantedBy=multi-user.target
```

OR

Try this: there is some issue with mtu size for the config above. This was used in one of the articles i read.

```
[Unit]
Description=replay TZSP packets
After=network.target

[Service]
Type=simple
User=root
ExecStart=/bin/sh -c '/usr/bin/tzsp2pcap -f | /usr/bin/tcpreplay-edit --topspeed --mtu=$(cat /sys/class/net/tzsp0/mtu) --mtu-trunc -i tzsp0 -'

Restart=always
RestartSec=5s

[Install]
WantedBy=multi-user.target
```

run the following:
```
sudo systemctl daemon-reload
sudo systemctl start sendtzsp
sudo systemctl enable sendtzsp
```

then have suricata monitor tzsp0.


Turns out there are a lot of people doing this with mikrotik:
https://serinformatica.ar/en/insights/security-onion-mikrotik-integration/

https://bløgg.no/2020/11/traffic-capturing-and-streaming-with-mikrotik-revisited/

https://github.com/laciKE/AtomIDS/

https://forum.mikrotik.com/t/suricata-ids-ips-integration-with-mikrotik-now-with-ossec/101160/169

https://github.com/zzbe/mikrocata
