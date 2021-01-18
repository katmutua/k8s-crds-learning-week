### CRDS
A CRD has 3 main parts
apiGroup  
Kind
Spec

Created custom resources are stored in the k8s cluster inside etcd next to regular resources.
To create a custom resource we do so in the k8s way (declarative using a yml file that describes
object rather than imperative with a script filled with commands to execute).

Example:
Say we want to create a CRD and controller that allow users to receive weather reports every d

```
# my-weather-reminder.yml
apiVersion: weather.com/v1alpha1
kind: Reminder
metadata:
   name: reminder-toronto
spec:
   city: Toronto, ON
   time: '0700'
   email: foo@example.com
```
##### Spec vs Status
status holds the current state of the system, is continously updated by a controller and can be read by the user.
spec is the property of a custom resource that holds the user intent i.e. all the fields that the user configured
that which represent the future state they want.

it's perfectly normal that there exists some time delay between before the user expressing their intent and the controller
(and other processes) completing the necessary actions to realize that intent.

This asynchronous, non-blocking way of working
is how k8s resources are expected to be implemented.

### Controller
Program that runs in the background paying attention to the creation/update/deletion of custom resources
of a specific CRD and takes action based on what it sees.

A controller can access the entire k8s API and manipulate any kind of resources.

How does a controller loop look like?
Therefore the term reconciliation which is essentially the controller loop looks like this

List the custom resources the controller is responsible for ---> Check what the spec (user intention) of those resources. --->
Check the status (current state) of those resources. -----> Take the appropriate action to move the current state closer to the desired state.
 ----> Update the current state --> Repeat.

A controller loop for our weather reminder will therefore looks like this

Look at all the weather reminder resources -> Check the target time of each reminder.
---> For each resource, if the current status is ready, and the target time matches the current time, then move on to the next step.
Otherwise, don't do anything yet with this reminder. ---> Schedule a container to fetch the weather and send an email --->
Update the status of the resource to be successful ----> Wait for a bit and start again from the top.
