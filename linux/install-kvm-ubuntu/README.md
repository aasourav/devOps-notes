REF: https://www.youtube.com/watch?v=qCUmf5gyOYY

```sh
sudo apt update
```

### Check virtualization is enabled (if yes then value should be greater than 0 )
```sh
egrep -c '(vmx|svm)' /proc/cpuinfo
```

### Install the cpu-checker package
```sh
sudo apt install -y cpu-checker
kvm-ok # this check virtualization is enabled
```

### Install KVM
```sh
sudo apt install -y qemu-kvm virt-manager libvirt-daemon-system virtinst libvirt-clients bridge-utils
```

### Enable and start virtualization daemon
```sh
sudo systemctl enable --now libvirtd
sudo systemctl start libvirtd
sudo systemctl status libvirtd
```

### Add Local User to the KVM and Libvirt Group
```sh
sudo usermod -aG kvm $USER
sudo usermod -aG libvirt $USER
```

### Cinfig for bridge network
```sh
sudo vi /etc/netplan/01-netcfg.yaml

network:
  version: 2
  renderer: NetworkManager
  ethernets:
    ens33:
      dhcp4: false
      dhcp6: false
  
  bridges:
    br0:
      interfaces: [ens33]
      dhcp4: false
      addresses: [192.168.2.120/24]

      routes:
to: default
          via: 192.168.2.1
          metric: 100
      nameservers:
        addresses: [8.8.8.8]
```

### Run the virtual machin (kvm)
```sh
sudo virt-manager
```