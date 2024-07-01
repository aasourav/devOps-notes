Once we have a user that is already authenticated.
Now we need authorization.

If a user want to talk with kubernetes cluster it's need first to be authenticate , then it's need authorization.

User Authenticate with config file / token bearer / pass user
Application/ Process authenticate with Service account

For authorization provides 4 basic obj/resources

api group is : `rbac.authorization.k8s.io`

resources are:

1. role -> namespace scoped resource like pod, etc
2. role binding -> namespace scoped resource
3. cluster -> cluster scoped resource ex: nodes, storage classes , pv
4. cluster role binding -> cluster scoped resource

check access ->

```yaml
k get clusterrole
k get role

k auth can-i delete pod -n kube-system --as system:serviceaccount:default:default
#output -> yes / no
```

```yaml
    apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  namespace: default
  name: example-role
rules:
- apiGroups: [""] # "" indicates the core API group
  resources: ["pods","pods/logs"]
  verbs: ["get", "watch", "list"]
- apiGroups: ["apps"]
  resources: ["deployments"]
  verbs: ["get", "watch", "list"]

```

to allow a service account / user ->

```yaml
k create role role_name -n namespace_name --verb delete,list,update -- resource pod,deployment
```

now we have to bind the role for SA or user

```yaml
    apiVersion: rbac.authorization.k8s.io/v1
    kind: RoleBinding
    metadata:
    name: example-rolebinding
    namespace: default
    subjects:
    - kind: User
    name: "example-user" # Name is case sensitive
    apiGroup: rbac.authorization.k8s.io
    roleRef:
    kind: Role
    name: example-role
    apiGroup: rbac.authorization.k8s.io
```

```yaml
#for service account
k create rolebinding rolebinding_name -n namespace_name --role role_name --serviceaccount service_account_name

#for user
k create rolebinding rolebinding_name -n namespace_name --role role_name --user user_name
```

if want to user can delete pod from multiple namespace we have to create role and role binding for them individually for user/sa.

but a case arise there are n number of namespace. what we do then?

```yaml
#Cluster role
    apiVersion: rbac.authorization.k8s.io/v1
    kind: ClusterRole
    metadata:
    name: example-clusterrole
    rules:
    - apiGroups: [""]
    resources: ["pods","pods/logs"]
    verbs: ["get", "watch", "list"]
    - apiGroups: ["apps"]
    resources: ["deployments"]
    verbs: ["get", "watch", "list"]

```

```yaml
    apiVersion: rbac.authorization.k8s.io/v1
    kind: ClusterRoleBinding
    metadata:
    name: example-clusterrolebinding
    subjects:
    - kind: User
    name: "example-user" # Name is case sensitive
    apiGroup: rbac.authorization.k8s.io
    roleRef:
    kind: ClusterRole
    name: example-clusterrole
    apiGroup: rbac.authorization.k8s.io

```

```yaml
# creating cluster role , and binding that cluster role to a user using role binding.

# k create clusterrole cluster_role_name --verb delete --resources pod --dry-run -oyaml > allowROle.yaml
k create clusterrole cluster_role_name --verb delete --resources pod

# k create rolebinding -n namespace_name --clusterrole clusterrole_name --user user_name --dry-run -yaml > abc.yaml


k create rolebinding -n namespace_name --clusterrole clusterrole_name --user user_name
# now we don't need create every role. we just now bind the cluster role . so decrease 1 step :D . every time we will just write the command (up given) just change the namespace.

# another way
# first create cluster role (see top)
k create clusterrolebinding cluster_role_binding_name --clusterrole clusterrole_name --user user_name # for sa --service-account default:default
# first_default = default namespace
# sec default = service acc name

# now we don't need to do anything. It can access from every namespace . :P . all step now ommited :D
```

## Aggregated cluster role -- comming soon..

If you want to create cluster role out of n number of cluster role. means let's say you create a Clusuter role by manisfest.
```yaml
    apiVersion: rback.authorization.k8s.io/v1
    kind: ClusterRole
    metadata:
        name: monitoring
    aggregationRule: # we will create a aggregate rule. agg rule is specify label
        clusterRoleSelectors:
            rbac.example.com/aggregate-to-monitoring: "true" # that is label. any cluster role with this label (next text (2))
        rules: [] #(2) those cluster role updated automatically here.
```

### cluster role with label `rbac.example.com/aggregate-to-monitoring: "true"`
```yaml
    apiVersion: rbac.authorization.k8s.io/v1
    kind: CLusterRole
    metadata:
        labels: 
            rbac.example.com/aggregate-to-monitoring: "true"
        name: test
    rules:
    - apiGroups:
      - ""
      resources:
      - services
      verbs:
      - list
      - watch
```

If we apply then , aggregated roles: [] will be updated

Have to learn more about it.

## now one more thing

there are something that are not resources ex: logs `k logs`
```yaml
    # for this logs add pods/log (here pod/logs is just example)
        --resources pods/log
```
