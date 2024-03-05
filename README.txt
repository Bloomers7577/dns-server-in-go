A simple dns server written in Go. 

To build go into /cmd/dns-resolver and type go build and ./dns-resolver <port-name> <blocked-file> 

by default it will bind to port 53 and use blocked domains in blocked.txt 

There is no cache implemented, every query will go to root servers.

This code is heavily (90%) based on https://www.youtube.com/watch?v=V3EAssIsQNI&t=2447s and https://github.com/wardviaene/golang-for-devops-course 

Deployed on VM and test with dig. 