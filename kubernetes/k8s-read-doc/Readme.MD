### Concept -> Overview -> Objects in kubernetes -> Kubernetes Object management
```sh
kubectl create deployment nginx --image nginx #this is called imparative command
kubectl create -f nginx.yaml
kubectl delete -f nginx.yaml -f redis.yaml # create also work
kubectl replace -f nginx.yaml #Update the objects defined in a configuration file by overwriting the live configuration
kubectl apply -f directoryName/ #deploy all manifest inside directoryName
kubectl diff -f directoryName/ #see the diff before apply (it look what will happen when we will apply)
kubectl apply -R -f directoryName/ # if it has sub directory then it will deploy subdirectory too. Mean it will deploy recursively
kubectl diff -R -f directoryName/ # same
```

### Concept -> Overview -> Objects in kubernetes -> Object name and ID

DNS Subdomain Names
-------------------
```sh
my-service.my-namespace.svc.cluster.local  #this is full domain
my-service.my-namespace # this is sub domain under `my-service.my-namespace.svc.cluster.local`
```

#### DNS Hierarchy in Kubernetes
1. Top-level domain: `.local` or `.cluster.local` (in a Kubernetes cluster)
2. Service domain: `svc.cluster.local`
3. Namespace subdomain: `my-namespace.svc.cluster.local`
4. Service subdomain: `my-service.my-namespace.svc.cluster.local`

Rules:
sub domain container no more than 253 character
contain only lowercase alphanumeric characters, '-' or '.'
start and end with an alphanumeric character


Accessing the Kubernetes API from Outside the Cluster
```sh
    kubectl proxy
    curl http://localhost:8001/api/v1/namespaces/default/pods/my-app-pod # now we can access this
```

UID: 
Every object created over the whole lifetime of a Kubernetes cluster has a distinct UID. It is intended to distinguish between historical occurrences of similar entities.


### Concept -> Overview -> Objects in kubernetes -> Labels and Selector

```sh
# environment=production,tier!=frontend we can search pods bu using label (-l)
kubectl get pods -l environment=production,tier=frontend
kubectl get pods -lapp=guestbook,role=replica

kubectl label pods -l app=nginx tier=fe #update labels

kubectl apply -f examples/guestbook/all-in-one/guestbook-all-in-one.yaml
kubectl get pods -Lapp -Ltier -Lrole # show pods with labels
# NAME                           READY  STATUS    RESTARTS   AGE   APP         TIER       ROLE
# guestbook-fe-4nlpb             1/1    Running   0          1m    guestbook   frontend   <none>
# guestbook-fe-ght6d             1/1    Running   0          1m    guestbook   frontend   <none>
# guestbook-fe-jpy62             1/1    Running   0          1m    guestbook   frontend   <none>

"metadata": {
  "labels": {
    "key1" : "value1", #[a-z0-9A-Z],must be 63 characters or less (can be empty), - , _ , . is allowed
    "key2" : "value2" #[a-z0-9A-Z]
  }
}

# Selector
selector:
 matchLabels:
  component: redis
#advanced selector:
selector:
  matchLabels:
    component: redis
  matchExpressions:
    - { key: tier, operator: In, values: [cache] }
    - { key: environment, operator: NotIn, values: [dev] }
```





