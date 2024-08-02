### Why we need this ?

to maintain which pod have to deploy in which kubernetes nodes. which kubernetes pod should avoid certain pods to be scheduled on them , which kubernetes nodes should become unschedule(downtime for customers) for all the pods.

```sh
    #kind create cluster with name with configuration
    kind create cluster --name=ahsan-sourav --config config-name.yaml

    # delete cluster
    kind delete cluster --name=ahsan-sourav

    kubectl config view | grep clustername

    kubectl use-context context-name #getting ctx name from top^ command
```

### the config.yaml

```yaml
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
  - role: control-plane
  - role: worker
  - role: worker
```

### Node Selector

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      nodeSelector: # here this is
        foo: bar # this should be match 100% in node label
      containers:
        - name: nginx
          image: nginx:1.14.2
          ports:
            - containerPort: 80
```

## how to add node lablel

```sh
    kubectl label nodes <node-name> <label-key>=<label-value>
    # or you can edit the nodes and add the label manually
```

## Node Affinity

the diff between `nodeSelector` and `affinity`

`nodeSelector` -> hard match ( if don't find specific label node the pod is not being scheduled )

`nodeAffinity` ->

1. Preferred = "if find the label use that node. if not found it can scheduled any node"

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      affinity:
        nodeAffinity:
          preferredDuringSchedulingIgnoredDuringExecution:
            - weight: 1
              preference:
                matchExpressions:
                  - key: foo
                    operator: In
                    values:
                      - bar
      containers:
        - name: nginx
          image: nginx:1.14.2
          ports:
            - containerPort: 80
```

2. Required = "almost same to the nodeSelector"

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
  labels:
    app: nginx
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
              - matchExpressions:
                  - key: foo
                    operator: In
                    values:
                      - bar1
```

## Taint

let's say you have k8s node and you don't want to anything scheduled on particular node. using taint

three types ot taint

1. NoSchedule = make the node ilde (pod will not schedule on that node). used for backup or upgrades
2. NoExecute = (little bit danger) all the pods on that particular node will stop working
3. PreferNoSchedule = let's say there is a node that parformance issues (there is some issue with the node). and you want that not to recommend to schedule pod on that node. but if no way to schedule on other nodes only then it will be schedule for pod.

```sh
    kubectl taint nodes <node-name> key1=value1:<taint-name>
```

## Toleration

is basically an execption. that is given to certain pods. ex: we have 3 nodes in the no schedule status (we have tainted the nodes with NoSchedule). for some reason we still have hight priority pods they should definately run. so what we can do? . to these pods we can add a <b>Toleration</b> / Execption so that they can run any tainted node.

how did we do Toleration. like before we set taint (`kubectl taint nodes <node-name> key1=value1:<taint-name>`) where we add key=value . if this key=value match then those pod will schedule.(also matched effect)

Example:

```sh
apiVersion: v1
kind: Pod
metadata:
  name: nginx
  labels:
    env: test
spec:
  containers:
  - name: nginx
    image: nginx
    imagePullPolicy: IfNotPresent
  tolerations:
  - key: "key1"
    operator: "Exists"
    value: "value1"
    effect: "NoSchedule"


```

for more info:

- [affinity , anti affinity](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/)
- [Assigning Pods to Nodes using Affinity and Anti-Affinity](https://pushkar-sre.medium.com/assigning-pods-to-nodes-using-affinity-and-anti-affinity-df18377244b9)
