# ksearch
A kubectl plugin that will help us list all (literally all) the resources in a namespace and the resources can be searched using names as well.
Right now the resources that are in apps/v1 and core/v1 resouce groups are being printed.

For now we have two main functionalitites in the pugin

# List The resources 
Listing the resources using standard `kubectl get` doesnt give us all the resources that are there in the cluster, for example ingresses and configmaps. Listing the resources using `kubectl search` will list all the resources of a namespace. Example command would be
```
kubectl search 
```
and it will list all the resources in all the namespaces. To list all the resources from a specific namespace you can use below command
```
kubectl search -n <ns-name>
```

# Filter the resources 
Now that we have all the resources from all the namespaces or from a specific namespace, we can filter the resoureces by giving a substring of the resource name, for example the below command 
```
kubectl search -name <res-name>
```
will list all the resources from all the namespaces if their name contains the string `res-name`. We can provide an options `-n` parameter to do the same in for a specific namespave.

# Getting Help
You can run below command to get the list of supported flags.
```
kubectl search -h
```

# Why ksearch
There are some situations when you would to know whether the configmaps/secrets that are being referred by the are there in the cluster or not, or the sevice has the respective endpoints created or not. If you use `kubectl get` utility you will have to get the secrets and configmaps separately, but using `ksearch` will list all the resources at once in one place. Similarly lets say you want to get all the resource that are deploys as part of `nginx` deployment, they will most probably have name `nginx` in them. You can just search for all those resources using the below command
```
kubectl search -name  <res-name>
```
if you know the namespace you can append that using `-n` flag.

# Demo
[![asciicast](https://asciinema.org/a/RDcSSrmq6m0hhsaxgIOO6nQ6D.svg)](https://asciinema.org/a/RDcSSrmq6m0hhsaxgIOO6nQ6D)