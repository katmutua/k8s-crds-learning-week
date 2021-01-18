### Example
source: https://github.com/tektoncd/pipeline/blob/master/docs/tutorial.md

1. apply the hello-world yml to create a task
2. view details with
```
tkn task describe echo-hello-world
```
3. instantiate the task in a task run by applying the taskrun.yml
4. check the task run details with
```
tkn taskrun describe echo-hello-world-task-run
```
5. to see more detail on your task run
```
tkn taskrun logs echo-hello-world-task-run
```
- To pass in inputs or outputs use one or more PipelineResources to define the artifacts you want to pass in and
  out of your Task
- given an example of a git resource specifies a git repository with a specific revision from which the Task will
  pull the source cod
6. create a pipeline resource
```
kubectl apply -f pipeline_resource.yml
```
7. create the image resource that specifies the repo to which the image built by the task will be pushed
```
kubectl apply -f image_resource.yml
```
8. configure task credentials before running the task
```
kubectl create secret docker-registry regcred \
                    --docker-server=<your-registry-server> \
                    --docker-username=<your-name> \
                    --docker-password=<your-pword> \
                    --docker-email=<your-email>
```
9. you must also specify a service account that uses this secret to execute our task run
```
 kubectl apply -f git_tkn_service_account.yml
```
10. apply your task run

```
apiVersion: tekton.dev/v1beta1
kind: TaskRun
metadata:
  name: build-docker-image-from-git-source-task-run
spec:
  serviceAccountName: tutorial-service
  taskRef:
    name: build-docker-image-from-git-source
  params:
    - name: pathToDockerFile
      value: Dockerfile
    - name: pathToContext
      value: $(resources.inputs.docker-source.path)/examples/microservices/leeroy-web #configure: may change according to your source
  resources:
    inputs:
      - name: docker-source
        resourceRef:
          name: skaffold-git
    outputs:
      - name: builtImage
        resourceRef:
          name: skaffold-image-leeroy-web
```
11. to examine the resources created so far you can run
```
kubectl get tekton-pipelines
```

the TaskRun binds the inputs and outputs to already defined PipelineResources, sets values for variable substitution
parameters, and executes the Steps in the Task.

- creating a pipeline
- running the pipeline
- configuring pipeline execution commands
