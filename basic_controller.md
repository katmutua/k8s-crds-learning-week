### Basic controller

To create a controller we can use KubeBuilder
Kubebuilder is a framework for creating k8s controllers

The project is maintained by the k8s SIG API Machinery.

1. install kustomize
2. install kubebuilder and configure path
3. initialize your project
```
$ mkdir my-weather-app-crd
$ cd my-weather-app-crd
$ kubebuilder init --domain example.com --repo github.com/Manifaust/k8s-custom-resources-learning-aid/examples/weather-app
```

- go.mod includes libraries from apimachinery, client-go, and controller-runtime.
These libraries makes it easier to implement controllers and CRDs.

4. use the `kubebuilder create` command to create CRDs and controllers.
```
kubebuilder create api --group weather-app --version v1alpha1 --kind CheckWeather
```
- This command creates a controller checkweather_controller.go and a also defines properties of the spec and status of your CRD

5. to install your crd
```
$ kubectl create namespace weather-testing
$ kubectl config set-context $(kubectl config current-context) --namespace=weather-testing
```
the above command sets the current context's namespace as weather-testing

```
$make install
```
or if you have a custom kubeconfig that's not in the $HOME/kube then you can use the KUBECONFIG environment variable to tell kubectl
where your kubeconfig is.
```
$ KUBECONFIG=<my_kubeconfig_location> make install
$ kubectl get crds
```

6. running your controller
- In a real production cluster, your controller is deployed in a regular k8s deployment. During our testing however, we'll instead just
run the controller locally on our computer as a normal golang app.
```
$ make run
```

7. now that the CRD is created let's create a custom resource

```
kubectl apply -f config/samples/weather-app_v1alpha1_checkweather.yaml
```
