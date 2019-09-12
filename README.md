//TODO Improve the README

# ksearch
A kubectl plugin that will help us list all (literally all) the resources in a namespace and the resources can be searched using names as well.

For now we have two main functionalitites in the pugin

# List The resources 
Listing the resources using standard `kubectl get` doesnt give us all the resources that are there in the cluster, for example ingresses and configmaps. Listing the resources using `kubectl search` will list all the resources of a namespace. Example command would be
```
kubectl search <res-name>
```

# Filter the resources 
We can filter the resoureces using metacharacter, for example the below command 
```
kubectl search res-name*
```
will list all the resources if their name starts with `res-name`. We can provide an options `namespace` parameter.
