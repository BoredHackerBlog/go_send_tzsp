# go_send_tzsp
Capture packets, tzsp encapsulate them, and send them to another host for storage or parsing

Problem: I needed something running on a remote low resource (low disk/storage) vps or iot device that would capture packets. 

There are multiple ways to solve my problem but I wanted to have this as an option. 

Inspired by Scratch-n-Sniff - https://github.com/nickvsnetworking/Scratch-n-Sniff, NoStarch BlackHat Go - https://nostarch.com/blackhatgo, and Greynoise Community Slack. ❤️

I picked go because I can't write code in C/C++ and I can't put python on capture source/remote device. 

<img width="678" alt="image" src="https://github.com/user-attachments/assets/a1c444d8-6c46-46b0-8fa3-40a635b65231">


# Usage
