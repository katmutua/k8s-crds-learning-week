### EXERCISE 1 BYO Controller checklist
1. Add a new field to the CRD's spec to accept a city location. To do that we need to modify checkweather_types.
```
type CheckWeatherSpec struct {
  City string `json:"city,omitempty"`
}
```
2. add constants to keep track of current state of our request
```
const (
  StatePending  = "PENDING"
  StateFinished = "FINISHED"
)
```

3. add temperature and state as part of the CheckWeatherStatus struct
```
type CheckWeatherStatus struct {
  State       string `json:"state,omitempty"`
  Temperature int32  `json:"temperature"`
}

// +kubebuilder:subresource:status
```
ensure to have the marker in the comment which is used to generate some code

4. update the controller code to implement the reconcile logic (see Reconcile func in `checkweather_controller.go`)
our reconcile loop will be
  - identify the custom resource (CheckWeather resource)
  - check if its state is READY
  - if the state is FINISHED don't do anything and don't enqueue
  - if the state is PENDING then read the target city from the Spec and check the weather and
    write the current temperature to the status
    Then set the state to FINISHED
  - if an error occurs then log the message and re-queue

5. remove previously created resource
```
kubectl delete checkweather checkweather-sample
```
6. update the resource config to include the new field added (City)
```
apiVersion: weather-app.example.com/v1alpha1
kind: CheckWeather
metadata:
  name: checkweather-sample
spec:
  city: "Mexico City"
```
7. rebuild and re-install the crd

```
$ make install
$ make run
```

8. create your resource to test it out
```
$ kubectl apply -f config/samples/weather-app_v1alpha1_checkweather.yaml
```

## Main Exercise

1. Create another controller and CRD called WeatherWarning
```
kubebuilder create api --group weather-app --version v1alpha1 --kind WeatherWarning
```

2. Set max temperature as part of the weather warning spec and apply the configs

```
kubectl delete  checkweather checkweather-sample && kubectl delete  weatherwarning weatherwarning-sample
kubectl apply -f config/samples/weather-app_v1alpha1_checkweather.yaml && kubectl apply -f config/samples/weather-app_v1alpha1_weatherwarning.yaml
make install
make run
```
3. Update the controller with the if logic to check against the CheckWeather resource
   and log as appropriate.
