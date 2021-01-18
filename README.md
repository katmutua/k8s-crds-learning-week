# k8s-crds-learning-week

Objectives 

Day 1: Build a basic CRD using KubeBuilder
Day 2: BYO CRD controller exercise 1
Day 3: BYO CRD controller exercise 2


Minikube 

Create a deployment
```
 kubectl create deployment hello-node --image=k8s.gcr.io/echoserver:1.4
```
Expose pod to the public internet
```
  kubectl expose deployment hello-node --type=LoadBalancer --port=8080
```
List addonos 
```
minikube addons list
```

Enable addons 
```
minikube addons enable metrics-server
```
Disable addons 
```
minikube addons disable metrics-server
```
Other 
```
kubectl get pod,svc -n kube-system
```
