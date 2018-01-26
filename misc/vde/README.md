# HOWTO: unprivileged netns using VDE

No root privilege nor SUID binary is required.

## Install VDE (vdeplug4)

- https://github.com.rd235/s2argv-execs
- https://github.com.rd235/vdeplug4
- https://github.com.rd235/libslirp
- https://github.com.rd235/vdeplug_slirp

An older version of VDE (VDE2) is available at Debian apt repo, but I haven't tested VDE2.

## Demo

### Terminal 1 [runc]

```
$ ./prepare-rootfs.sh
$ runc run foo
runc# ip tuntap add name tap0 mode tap 
runc# ip link set tap0 up 
```

- The example `config.json` contains `network` namespace, `CAP_NET_ADMIN`, and `CAP_NET_RAW` capabilities.
- You may need to fix `uidMappings` and `gidMappings` according to your UID/GID.
- No need to use runROOTLESS for this demo. Plain runC is enough.

### Terminal 2 [vxvde-slirp]

```
$ vde_plug vxvde:// slirp://
```

### Terminal 3[vxvde-netns-tap]
```
$ vde_plug vxvde:// = nsenter -- -t $(runc state foo | jq -r .pid) -n -U --preserve-credentials vde_plug tap://tap0
```

### Terminal 1 [runc]

```
runc# ip link set tap0 up
runc# ip addr add 10.0.2.100/24 dev tap0
runc# ip route add default via 10.0.2.2 dev tap0
runc# echo "nameserver 10.0.2.3" > /etc/resolv.conf
runc# apk add --no-cache w3m
runc# w3m https://rootlesscontaine.rs/
```
