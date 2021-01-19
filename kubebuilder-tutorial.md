### KubeBuilder Tutorial

1. create a new project 
   ```
     mkdir $GOPATH/kubebuilder-tutorial
     cd $GOPATH/kubebuilder-tutorial
   ```
 2. create the scaffold for the project 
    ```
     kubebuilder init --domain tutorial.kubebuilder.io
    ```
    
 Reading through the code
 - each main.go file starts with basic imports 
   the core container library - `controller-runtime library`
   see `"k8s.io/apimachinery/pkg/runtime"`

- every controller needs a scheme and which provides a mapping between Kinds and corresponging Go types
    ```
    var (
        scheme   = runtime.NewScheme()
        setupLog = ctrl.Log.WithName("setup")
    )
    
    func init() {
    
        // +kubebuilder:scaffold:scheme
    }
    ``` 
 