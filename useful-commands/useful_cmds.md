```
    //aes
    make generate manifest  -> aesl api-resource build
    sudo hostname new_host_name -> change hostname

    //kubernetes
    kubens -> switch namespace, it's also inside in kubectx
    kubectx -> tools to switch contexts(cluster) on kubectl faster
    kubectl port-forward service/ingress-nginx-controller 8080:80 -n core-ns --address 0.0.0.0
    mongosh "mongodb://ratul:1234@172.18.255.202:3342/admin?retryWrites=true&w=majority"
    mongodb://databaseAdmin:databaseAdmin123456@host:port/admin?retryWrites=true&w=majority // this is for docker
```