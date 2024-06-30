[Official Documentation](https://github.com/kubernetes-sigs/metrics-server)

### Installation

```sh
#for prod
 kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml

 #for dev we have to modify the yaml file
 #first 
 curl -L -o components.yaml https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
 # then edit the compnents.yaml file
```

now we have to add `--kubelet-insecure-tls` under the 
```
spec:
      containers:
      - args:
        - --cert-dir=/tmp
        - --secure-port=10250
        - --kubelet-preferred-address-types=InternalIP,ExternalIP,Hostname
        - --kubelet-use-node-status-port
        - --metric-resolution=15s
        - --kubelet-insecure-tls=true
```
