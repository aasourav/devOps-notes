### Overall configurations steps
    - Setting up the Static IPV4 on all nodes.
    - Disabling swap & Setting up hostnames.
    - Installing Kubernetes components on all nodes.
    - Initializing the Kubernetes cluster.
    - Configuring Kubectl.
    - Configure Calico Network operator.
    - Print Join token & add worker nodes in the cluster.
    - Deploy Applications.

#### 1. Setting static IPV4 on all nodes (Master & Workers)

```sh
sudo vim /etc/netplan/01-network-manager-all.yaml
```

```yaml
network:
  version: 2
  renderer: networkd
  ethernets:
    ens33: # this is driver name, it can be diffrent
      dhcp4: no
      addresses:
        - 192.168.10.245/24
      routes:
        - to: default
          via: 192.168.10.1
      nameservers:
          addresses: [8.8.8.8, 8.8.4.4]
```
```sh
sudo netplan apply
```

#### Remove swap (Master & worker)

```sh
sudo apt update
sudo swapoff -a

# comment the swap command in fstab file
sudo sed -i '/^\/swap.img/ s/^/#/' /etc/fstab
sudo init 6 # have to test without using it
```

#### change hostname (Master & worker)
```sh
hostnamectl set-hostname master
```

#### Configure modules (Master & worker)

```sh
cat <<EOF | sudo tee /etc/modules-load.d/k8s.conf
overlay
br_netfilter
EOF
```

```sh
sudo modprobe br_netfilter
sudo modprobe overlay
```

#### Configure networking (Master & worker)

```sh
cat <<EOF | sudo tee /etc/sysctl.d/k8s.conf
net.bridge.bridge-nf-call-iptables  = 1
net.bridge.bridge-nf-call-ip6tables = 1
net.ipv4.ip_forward                 = 1
EOF
```


#### Install containerd (Master & worker)
```sh
sudo apt-get update
sudo apt-get install -y containerd
```

#### modify containerd configuration (Master & worker)

```sh
sudo mkdir -p /etc/containerd
sudo containerd config default | sudo tee /etc/containerd/config.toml
sudo sed -i 's/SystemdCgroup \= false/SystemdCgroup \= true/g' /etc/containerd/config.toml
cat /etc/containerd/config.toml
```

```sh
sudo systemctl restart containerd.service
sudo systemctl status containerd
```


#### Install Kubernetes Management Tools (Master & Worker).
```sh
sudo apt-get update
sudo apt-get install -y ca-certificates curl
sudo apt-get install -y apt-transport-https ca-certificates curl


curl -fsSL https://pkgs.k8s.io/core:/stable:/v1.29/deb/Release.key | sudo gpg --dearmor -o /etc/apt/keyrings/kubernetes-apt-keyring.gpg

echo 'deb [signed-by=/etc/apt/keyrings/kubernetes-apt-keyring.gpg] https://pkgs.k8s.io/core:/stable:/v1.29/deb/ /' | sudo tee /etc/apt/sources.list.d/kubernetes.list

```

```sh
sudo apt-get update
sudo apt-get install -y kubelet kubeadm kubectl
sudo apt-mark hold kubelet kubeadm kubectl
```

#### Initialization the kubernetes cluster (Master)

```sh
sudo kubeadm init --apiserver-advertise-address=172.17.17.200 --pod-network-cidr=192.168.0.0/16 --cri-socket /run/containerd/containerd.sock --ignore-preflight-errors Swap
```

#### Configure Kubectl (Master )
```sh
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
```

### Install Calico Networking for on premises deployment (Master)

```sh
kubectl create -f https://raw.githubusercontent.com/projectcalico/calico/v3.24.1/manifests/tigera-operator.yaml
```
```sh
curl https://raw.githubusercontent.com/projectcalico/calico/v3.24.1/manifests/custom-resources.yaml -O


kubectl create -f custom-resources.yaml # we can customize this file
```

#### Print joing token for worker node to join cluster (Master)
```sh
kubeadm token create --print-join-command
```
then you will see a command to join worker nodes

ex:

```sh
kubeadm join 172.17.18.200:6443 --token 5g5jo2.agl26wfzkujgjt3s --discovery-token-ca-cert-hash ha256:57795a664200425258ed0619af960fe476d1ae93f99182a3d710ce1185468d3f
```