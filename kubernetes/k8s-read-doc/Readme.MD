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


### Concept -> Overview -> Objects in kubernetes -> Namespace

Namespaces are intended for use in environments with many users spread across multiple teams, or projects. 

Initial Namesapces:
 - default
 - kube-node-lease
 - kube-public
 - kube-system

 ```sh
kubectl get namespace # k get ns
#You can permanently save the namespace for all subsequent kubectl commands in that context.
kubectl config set-context --current --namespace=<insert-namespace-name-here>
 ```
all namespace names must be valid RFC 1123 DNS labels(63 char, lower alpha num char , "-").

Namespaces and DNS . Public Domain Confusion:

Example:

#### Namespace and Service Setup:
- You create a namespace called com and a service named myapp.
- Kubernetes creates a DNS entry myapp.com.svc.cluster.local.

#### Public Domain Conflict:
- Your app inside the cluster tries to access the real public website myapp.com.
- Kubernetes checks internal DNS first (bcz it's prioritize internal dns) and thinks myapp.com refers to the service myapp in the com namespace.

#### The Problem:
- Instead of reaching the public website myapp.com, your app connects to the internal service myapp in the com namespace.

#### Why This Is Dangerous:
- It can block access to public services and APIs.
- Attackers can create namespaces like com, org, or net to intercept traffic.

#### How to Prevent:
- Avoid naming namespaces after public TLDs like com, org, net, etc.
- Restrict who can create namespaces.
- Use admission webhooks to block such namespaces.


Not all objects are in a namespace

```sh
# To see which Kubernetes resources are and aren't in a namespace:
# In a namespace
kubectl api-resources --namespaced=true

# Not in a namespace
kubectl api-resources --namespaced=false
```

#### The Kubernetes control plane sets an immutable label kubernetes.io/metadata.name on all namespaces. The value of the label is the namespace name.



### Concept -> Overview -> Objects in kubernetes -> Annotations

We can store some information (you can say key value pair db :D ) and attatch it to the pod. 

```sh
"metadata": {
  "annotations": {
    "key1" : "value1",
    "key2" : "value2"
  }
}
```

### Concept -> Overview -> Objects in kubernetes -> Field Selectors

```sh
kubectl get pods --field-selector status.phase=Running
```

Supported field selectors vary by Kubernetes resource type. All resource types support the metadata.name and metadata.namespace fields.

```sh
#Using unsupported field selectors produces an error. For example:
kubectl get ingress --field-selector foo.bar=baz
```

#### List of supported fields

| Kind                      | Fields                                                                                            |
|----------------------------|---------------------------------------------------------------------------------------------------|
| **Pod**                    | spec.nodeName<br>spec.restartPolicy<br>spec.schedulerName<br>spec.serviceAccountName<br>spec.hostNetwork<br>status.phase<br>status.podIP<br>status.nominatedNodeName |
| **Event**                  | involvedObject.kind<br>involvedObject.namespace<br>involvedObject.name<br>involvedObject.uid<br>involvedObject.apiVersion<br>involvedObject.resourceVersion<br>involvedObject.fieldPath<br>reason<br>reportingComponent<br>source<br>type |
| **Secret**                 | type                                                                                              |
| **Namespace**              | status.phase                                                                                      |
| **ReplicaSet**             | status.replicas                                                                                   |
| **ReplicationController**  | status.replicas                                                                                   |
| **Job**                    | status.successful                                                                                 |
| **Node**                   | spec.unschedulable                                                                                |
| **CertificateSigningRequest** | spec.signerName                                                                                 |



You can use the =, ==, and != operators with field selectors (= and == mean the same thing).

```sh
kubectl get services  --all-namespaces --field-selector metadata.namespace!=default

#chained selectors
kubectl get pods --field-selector=status.phase!=Running,spec.restartPolicy=Always 

#Multiple resource types
kubectl get statefulsets,services --all-namespaces --field-selector metadata.namespace!=default

```

### Concept -> Overview -> Objects in kubernetes -> Finalizers

How finalizers work 
When you create a resource using a manifest file, you can specify finalizers in the metadata.finalizers field. When you attempt to delete the resource, the API server handling the delete request notices the values in the finalizers field and does the following:

- Modifies the object to add a metadata.deletionTimestamp field with the time you started the deletion.
- Prevents the object from being removed until all items are removed from its metadata.finalizers field
- Returns a 202 status code (HTTP "Accepted")

When you DELETE an object, Kubernetes adds the deletion timestamp for that object.
You can remove existing finalizers (deleting an entry from the finalizers list) but you cannot add a new finalizer. You also cannot modify the deletionTimestamp for an object once it is set. 

After the deletion is requested, you can not resurrect this object. The only way is to delete it and make a new similar object

Owner references are added to dependent objects (e.g., Pods created by a Job) to establish relationships, ensuring Kubernetes knows which resources to clean up when the owner (e.g., Job) is deleted.

Finalizers in Foreground and Orphan Deletion:

<b>Foreground Deletion:</b> When you delete a resource (like a controller), Kubernetes waits for its dependents (e.g., Pods) to be deleted first. It uses a foreground finalizer to ensure this, and the dependents must have ownerReferences.blockOwnerDeletion=true. Once the dependents are gone, the owner resource gets deleted.

<b>Orphan Deletion:</b> In this case, when you delete an owner resource, Kubernetes does not delete its dependents. It uses an orphan finalizer to ensure the dependents remain even after the owner is deleted.

### Concept -> Overview -> Objects in kubernetes -> Recommended Labels

| Key                              | Description                                                        | Example      | Type   |
|----------------------------------|--------------------------------------------------------------------|--------------|--------|
| app.kubernetes.io/name           | The name of the application                                         | mysql        | string |
| app.kubernetes.io/instance       | A unique name identifying the instance of an application            | mysql-abcxyz | string |
| app.kubernetes.io/version        | The current version of the application (e.g., a SemVer 1.0, revision hash, etc.) | 5.7.21      | string |
| app.kubernetes.io/component      | The component within the architecture                               | database     | string |
| app.kubernetes.io/part-of        | The name of a higher level application this one is part of          | wordpress    | string |
| app.kubernetes.io/managed-by     | The tool being used to manage the operation of an application       | Helm         | string |


```sh
#example
apiVersion: apps/v1
kind: StatefulSet
metadata:
  labels:
    app.kubernetes.io/name: mysql
    app.kubernetes.io/instance: mysql-abcxyz
    app.kubernetes.io/version: "5.7.21"
    app.kubernetes.io/component: database
    app.kubernetes.io/part-of: wordpress
    app.kubernetes.io/managed-by: Helm
```



### Concept -> Containers -> Images

A container image represents binary data that encapsulates an application and all its software dependencies.

container image can also include a registry hostname:
- `fictional.registry.example:10443/imagename`

- if dont specify the hostname , k8s assumes it docker public registry
- After the image name part you can add a tag or digest 
-  It can be up to 128 characters long.
- And must follow the next regex pattern: [a-zA-Z0-9_][a-zA-Z0-9._-]{0,127}
- If you don't specify a tag, Kubernetes assumes you mean the tag latest



Image pull policies:
- IfNotPresent: the image is pulled only if it is not already present locally.
- Always: What "Always" does: When you set imagePullPolicy: Always, Kubernetes will always check with the image registry (like Docker Hub) every time it runs a container. What Kubernetes checks for: It asks the registry, "What is the digest of this image?" (A digest is a unique fingerprint for a specific version of the image.)
- Never: the kubelet does not try fetching the image. If the image is somehow already present locally, the kubelet attempts to start the container; otherwise, startup fails.

Note: You should avoid using the :latest tag when deploying containers in production as it is harder to track which version of the image is running and more difficult to roll back properly.

To make sure the Pod always uses the same version of a container image, you can specify the image's digest; replace <image-name>:<tag> with <image-name>@<digest> (for example, image@sha256:45b23dee08af5e43a7fea6c4cf9c25ccf269ee113168c19722f87876677c5cb2).

- you omit imagePullPolicy,  default : `IfNotPresent`
- you omit imagePullPolicy, if you set image tag to `latest`, default: `Always`
- you omit imagePullPolicy, if you omit image tag , default: `Always`
