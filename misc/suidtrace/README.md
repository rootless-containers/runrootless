# Trace SUID binary executions

## Trace runc / runROOTLESS

```
$ stap suidtrace.stp -c "runc run -b /path/to/bundle foo"
```

## Appendix: how to install Systemtap to Ubuntu

1. Install systemtap

```
$ sudo apt install systemtap
```

2. Add yourself to `stapusr` and `stapdev` groups

```
$ sudo vigr
$ sudo vigr -s
```

3. Set up debug symbol repo ( https://wiki.ubuntu.com/DebuggingProgramCrash )
```
$ sudo tee -a /etc/apt/sources.list.d <<EOF
deb http://ddebs.ubuntu.com artful main restricted universe multiverse
deb http://ddebs.ubuntu.com artful-updates main restricted universe multiverse
deb http://ddebs.ubuntu.com artful-proposed main restricted universe multiverse
EOF
$ sudo apt-key adv --keyserver keyserver.ubuntu.com --recv-keys 5FDFF622
$ sudo apt update
```

4. Install kernel debug symbol pkg

```
$ sudo apt install linux-image-$(uname -r)-dbgsym 
```

5. Execute `stap-prep`

```
$ sudo stap-prep
```
