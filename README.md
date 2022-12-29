[![Go](https://github.com/cwxstat/ipsecrt/actions/workflows/go.yml/badge.svg)](https://github.com/cwxstat/ipsecrt/actions/workflows/go.yml)
# ipsecrt
Go program to add route entries used in /etc/ppp/ip-{up/down}

# Example of /etc/ppp/ip-up

```bash
cat /etc/ppp/ip-up 
#!/bin/sh
now=`date +%Y-%m-%d_%Hh%Mm%Ss`
logfile=/var/log/ip-up.log

echo "$0 called at $now with following params:" >> $logfile
echo "The VPN interface (e.g. ppp0): $1" >> $logfile
echo "Unknown, was 0, in my case: $2" >> $logfile
echo "IP of the VPN server: $3" >> $logfile
echo "VPN gateway address: $4" >> $logfile
echo "Regular (non-vpn) gateway for your lan connections: $5" >> $logfile

# Add 34.160.0.0 range to routing table on VPN interface
#/sbin/route add -net 34.160.0.0/16 -interface 192.168.1.1 >> $logfile 2>&1
/usr/local/bin/ipsecrt  zoomadd >> $logfile 2>&1
/usr/local/bin/ipsecrt githubadd >> $logfile 2>&1

```

# Example of /etc/ppp/ip-down

```bash
cat /etc/ppp/ip-down 
#!/bin/sh
now=`date +%Y-%m-%d_%Hh%Mm%Ss`
logfile=/var/log/ip-up.log

echo "Down called: /etc/ppp/ip-down" >> $logfile
echo "$0 called at $now with following params:" >> $logfile
echo "The VPN interface (e.g. ppp0): $1" >> $logfile
echo "Unknown, was 0, in my case: $2" >> $logfile
echo "IP of the VPN server: $3" >> $logfile
echo "VPN gateway address: $4" >> $logfile
echo "Regular (non-vpn) gateway for your lan connections: $5" >> $logfile

/usr/local/bin/ipsecrt  zoomdel >> $logfile 2>&1
/usr/local/bin/ipsecrt githubdel >> $logfile 2>&1

```
