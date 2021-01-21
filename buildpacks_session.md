### Buildpacks Expert Session


meeeting was recorded 

usage of custom resource in kpack (read the repp)
ClusterStack Resource 
ClusterStore Resource 
  - repository of avaialable buildpacks 


  add after any get kubectl command to get the yml file -o yaml 


The Builder resource references a ClusterStore and a ClusterStack resource 

triggering a build in kpack eg after a clusterstack change 


notary setting is added to image spec not pod spec :') 
#TODO Reading: 
Say we need to update our image 
This flow - Stack resource changes - Build changes  - leImage changes and tiggeres a build 