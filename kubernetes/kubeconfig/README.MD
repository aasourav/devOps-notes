User
----------------------------> [k8s]
APP (service acc)

when kubectl talk with kubsernetes. kubectl uses special file name kubeconfig (locate in ~/.kube/config)

Auth in kubernetes we heavily use TLS based authentication. If a user wants to authenticate themselves against kubernetes cluster. they will have to specify client certificate and key .
one those certificate and key is correct the client would be authenticated against kubernetes cluster

what inside kube config file.

# Kubernetes Kubeconfig File Documentation

A `kubeconfig` file is a YAML file used by `kubectl` and other Kubernetes tools to configure access to clusters. It can contain multiple contexts, clusters, and user configurations, allowing you to switch between different clusters and namespaces.

## Structure of a Kubeconfig File

A `kubeconfig` file generally contains the following sections:

- `apiVersion`: The version of the kubeconfig file format. Currently, it is `v1`.
- `kind`: The type of resource. For kubeconfig files, it is `Config`.
- `clusters`: A list of clusters that the user can connect to.
- `contexts`: A list of contexts, each specifying a cluster, a user, and a namespace. it's create map between cluster and user.
- `current-context`: The context currently being used.
- `preferences`: General preferences. This is usually empty.
- `users`: A list of users that can access the clusters.

### Example Kubeconfig File

```yaml
apiVersion: v1
kind: Config
clusters:
  - cluster:
      certificate-authority-data: LS0tLS1CRUdJTiBDRVJU...
      server: https://my-cluster-server
    name: my-cluster
contexts:
  - context:
      cluster: my-cluster
      namespace: my-namespace
      user: my-user
    name: my-context
current-context: my-context
preferences: {}
users:
  - name: my-user
    user:
      client-certificate-data: LS0tLS1CRUdJTiBDRVJU...
      client-key-data: LS0tLS1CRUdJTiBSU0EgUElW...
```

```sh
#create user
kubectl config set-credentials user_name  --client-certificate .crt_path --client-key .key_path

#create context
kubectl config set-context context-name --cluster cluster-name --user user-name-which-we-want-to-map-with

# to see which context is using (who is talking with k8s)
kubectl config current-context

```

#### How `kubectl` figures out the kubeconfig file (how it's know where kubeconfig is located)?

kubectl by default looks for three things (to get kubeconfig data) in this particular order.

1. it's look for --kubeconfig flag in kubectl command
2. it looks for KUBECONFIG env
3. it looks for ~/.kbue/config

if we want to merge multiple kube config file. we have to add it in env variable (in KUBECONFIG).

```sh
export KUBECONFIG=~/.kube/config:another_path1:another_path2

# now if we run
kubectx #a package to switch context easily
# then we can see all the contexts what was in config files that  path was added in KUBECONFIG env
```

Note: User authentication (Users in config) mechanism might not similar ( some cloud provider provide different way )

```sh
    kubectl get pods -n kube-system
```

in output we can see a pod name
`kube-apiserver-kind-control-plane`
if we describe it.

we can see

```sh
    ......
    ......
    Command:
      kube-apiserver
      --advertise-address=172.25.0.4
      --allow-privileged=true
      --authorization-mode=Node,RBAC
      --client-ca-file=/etc/kubernetes/pki/ca.crt
      ......
      ......
```

here `--client-ca-file=/etc/kubernetes/pki/ca.crt` by using the `ca.crt` certificate to verify that the client's certificate is valid and was signed by the trusted CA. Means

Let's say I have kubernetes API server
it has `ca.crt` and `ca.key`

If I want a client to be authenticated against this API server. I would create

1. Private key file for the client (end user of k8s)
2. Using private key I will crete `csr` (certificate signing request). this `csr` file would be sign using kubernetes api server's `ca.crt` (what we have seen using command `--client-ca-file=/etc/kubernetes/pki/ca.crt`)
3. after signing we will be get `file_name.crt` file. once we have to `file_name.crt` file , we are going to use this file and the key file (see number 1) be authenticate as a valid user to the kube api server

```sh
    # create private key
    openssl genrsa -out name_of_private.key 2048

    # create csr file
    openssl req -new -key name_of_private.key -out name_of.csr -sub "/CN=common_name/O=organization_name"

```

NB: in the openssl req command you provided, /CN= (Common Name) and /O= (Organization) are mandatory fields when generating a Certificate Signing Request (CSR). These fields are typically required by certificate authorities (CAs) to validate and issue a certificate. The CN field usually represents the domain name associated with the certificate, and the O field represents the organization name. It's important to fill these fields accurately according to your requirements and the CA's guidelines.

Once we have `.csr` there are two ways to generate `.crt`.

Way 1. `.csr` file -> use k8s `.crt` and `.key` to sign certificate `.crt` ( this is manual way)

Way 2 : we can create certificate signing request (it's a k8s obj) resource in kubernetes cluster. once it's approved we will get the `.crt` in this (certificate signing request) manifest.

### Way 1 (Hands one)

we have already created `.csr` (certificate signign request) .

exec docker container (here for kind cluster. we know kind is running on docker).

```sh
    docker container ls
```

now we will look for the container name `kindset/node:v1.20.2`
copy the container id. and then exec it.

```sh
    docker exec container_id_or_name -it bash
    cd /etc/kubernetes/pki/
    ls -l
```

now we can see the `.crt` file and `.key` file. here.
now we will create 1. key 2. csr
both of this i write on top of this doc.

now sign the certificate

```sh
    openssl x500 -req -n csr_file. -CA /etc/kubernetes/pki/ca.crt -CAkey /etc/kubernetes/pki/ca.key -CAcreateserial -out output.crt -day 20
```

now we should be see the `.crt` file Here (output.crt). now copy the `.key` and `.crt` and use that to access the api server (add this in kubeconfig)

Now this user is authorized (after add this in config file and map it to the context). but not authenticated (have to add role for this user (RBAC)).

```sh
    # this role_name will can only see namespace list
    # don't write this in newly created. bcz this resoure does not have the access. so first add role from diffrent user.

    kubectl create role role_name --resource pods --verb list,get,create...

    # now bind the role (role binding)
    kubectl create rolebinding role_binding_name --role role_name --user user_name

```

### Way 2 (using k8s obj) Hands on

for this request we don't need to specify `-sub "/CN=common_name/O=organization_name"` to create `.csr ` file.

CN -> common_name , and O -> organization/grooup specify the user

But it will ask those information when you keep blank them. give random name. give common name : what you will do in yaml

here we define group `O=organization_name"` and user `CN=common_name/` in `.csr` file. but in manifest (given below) we don't need to specify those in creating `.csr` file (what we added in the yaml) yaml.

```yaml
# this specify the user
metadata:
  name: my-csr

# this specify the groups
 groups:
    -
```

```yaml
apiVersion: certificates.k8s.io/v1
kind: CertificateSigningRequest
metadata:
  name: my-csr
spec:
  groups: # groups where `my-csr` belongs to.
    -
  request: $(cat my-csr.csr | base64 | tr -d '\n')
  #before signer can sign the certificate it should be approved
  # if not approved signer will be signe
  signerName: kubernetes.io/kube-apiserver-client # who will sign
  usages:
    - client auth
```

`k create -f csr.yaml` -> apply the yaml file
`kubectl get csr` -> to see the certificate signing requests

it still in pending state. now we have to approve the certificate by authorized user (admin, or who have the power).

`k certificate approve name` . name -> what we set the name in yaml file

now it will be approved . for prove you can run `kubectl get csr`

now to see the certificate run

`k get csr name -oyaml`
now we can see under `status`, name `certificate`

copy this and undecode it and create a `.csr` file and paste it .
then you know what to do.

## Adding users to a manages kubernetes cluster (ex: EKS)

create key, csr file as mentioned previous. in csr we will add CN and O

```yaml
apiVersion: certificates.k8s.io/v1
kind: CertificateSigningRequest
metadata:
  name: my-csr
spec:
  groups: # groups where `my-csr` belongs to.
    -
  request: $(cat my-csr.csr | base64 | tr -d '\n')
  #before signer can sign the certificate it should be approved
  # if not approved signer will be signed
  signerName: kubernetes.io/kube-apiserver-client # who will sign
  usages:
    - client auth
```

rest of the things is similiar to WAY-2

now create user and context as like before

#### create user

kubectl config set-credentials user_name --client-certificate .crt_path --client-key .key_path --embed-certs-true

#### create context

kubectl config set-context context-name --cluster cluster-name --user user-name-which-we-want-to-map-with

all are same .

more about [Certificate signing request](https://kubernetes.io/docs/reference/access-authn-authz/certificate-signing-requests/)
