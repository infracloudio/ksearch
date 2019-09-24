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

# Specify the kinds that you want list
There are chances that you want some extra resources listed along with the resources that gets listed when we `kubectl get` but not all. `kubectl search` lists a lot of resources that you might not want to see, now you can alos specify the kinds that you want to see along with the basic resources that get displayed when we use `kubectl get`.
For example to get all the basic resources (that get displayed when we use `kubectl get`) and some extra resources (`configmaps` and `secret`) from `kube-system` namespace we can use 
```
kubectl search -n kube-system -kinds configmaps,secret
```
We can alwasy use `-name` flag to filter the only resources that match the provided value. For example below command 
```
kubectl search -n kube-system -kinds configmap,secret,serviceaccount -name nginx
```
will list all the basic resources (that we get from `kubectl get`), configmaps, secrets and serviceaccounts from the `kube-system` namespace that have `nginx` in their name.

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

# Installation 
//TODO, add to kubectl krew plugin repo

# Demo
[![asciicast](https://asciinema.org/a/quPHY6X6eVhkNtJ1Q0c0Z6PxC.svg)](https://asciinema.org/a/quPHY6X6eVhkNtJ1Q0c0Z6PxC)