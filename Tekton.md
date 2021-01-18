### Tekton

Tekton is a k8s native CI/CD tool/ the Kubernetes-native continuous integration (CI) and
continuous delivery (CD) engine
Tekton’s basic building blocks are called
 - Steps -  individual command to execute
 - Tasks - represents a discrete, self-contained part of a process therefore comprises one or more Steps
Tekton is implemented as a set of Kubernetes CustomResourceDefinitions (CRDs),

Tekton introduces several new CRDs including Task, Pipeline,TaskRun, and PipelineRun.

A PipelineRun represents a single running instance of a Pipeline

PipelineRun is responsible for creating a Pod for each of its Tasks and as many containers within each Pod as
it has Steps. < see pipeline run diagram >


Tekton runs scripts and other commands in containers.
Steps map to containers.
Tasks map to Pods.
PipelineRuns are Pipeline instances, such as a single Travis build or Jenkins job.
PipelineRuns manage their associated Pods (by creating TaskRuns).
Pods often share a temporary filesystem backed by a PersistentVolume.

Tekton vs other CI/CD tools? Tekton is currently optimized for building and deploying cloud-native, microservice-based
applications compared to the more general-purpose nature of the others

Tekton’s use of custom resources means that a Pipeline must exist in a namespace before it can be used

### Setup
1. install tekton
```kubectl apply --filename https://storage.googleapis.com/tekton-releases/pipeline/latest/release.yaml
```
2. Install tekton cli

```brew install tektoncd-cli
```
3. Configure storage on a GCS bucket in your cluster
```
## example
apiVersion: v1
kind: Secret
metadata:
  name: tekton-test-storage
  namespace: tekton-pipelines
type: kubernetes.io/opaque
stringData:
  gcs-config: |
    {
      < config_json_for_authentication >
    }
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: config-artifact-bucket
  namespace: tekton-pipelines
data:
  location: gs://mybucket
  bucket.service.account.secret.name: tekton-storage
  bucket.service.account.secret.key: gcs-config
  bucket.service.account.field.name: GOOGLE_APPLICATION_CREDENTIALS
```

Setup for local
https://github.com/tektoncd/pipeline/blob/master/docs/developers/local-setup.md
