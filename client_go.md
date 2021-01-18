### client-go

client-go is a tool for accessing the k8s api in golang and can use it to get information about pods,
deployments, and any other types of resources.

1. clone repo
```
  git clone https://github.com/kubernetes/client-go
```
By default, client-go uses the current context in the kubeconfig to access your cluster.

2. exercise looking for a pod in a certain namespace
 - update main.go to search for relevant pod and namespace

3. try a different demo which creates, updates, and deletes a deployment using client-go
- run main go file in create and delete deployment
