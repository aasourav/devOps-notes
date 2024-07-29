## Pods, Deployment, Replicaset, Services

```sh
    kubectl run newdeploy -n alpha --image nginx --dry-run -o yaml
    kubectl run newdeploy -n alpha --image nginx --dry-run -o yaml > pod.yaml
```

```sh
    #command line documentation
    kubectl explain pod
    kubectl explain pod.metadata
```
