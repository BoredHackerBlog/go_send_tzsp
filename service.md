# use as service
Tested on Ubuntu 26.10.

I was using tzsp2pcap (https://github.com/thefloweringash/tzsp2pcap) to listen for tzsp packets. It seems to be able to handle input from multiple machines sending tzsp packets.

Dstip was running tzsp2pcap

- Copy compiled binary to /usr/bin/go_send_tzsp
- add /etc/systemd/system/sendtzsp.service

```
[Unit]
Description=Send TZSP packets
After=network.target

[Service]
Type=simple
User=root
ExecStart=/usr/bin/go_send_tzsp -dstip 192.168.1.2 -iface ens3 -filter 'port 22'

Restart=always
RestartSec=5s

[Install]
WantedBy=multi-user.target
```

- run the following:

```
sudo systemctl daemon-reload
sudo systemctl start sendtzsp
sudo systemctl enable sendtzsp
```

- run `sudo systemctl daemon-reload` every time service file is changed to change config
