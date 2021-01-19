/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	weatherappv1alpha1 "github.com/Manifaust/k8s-custom-resources-learning-aid/examples/weather-app/api/v1alpha1"
)

// WeatherWarningReconciler reconciles a WeatherWarning object
type WeatherWarningReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=weather-app.example.com,resources=weatherwarnings,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=weather-app.example.com,resources=weatherwarnings/status,verbs=get;update;patch

func (r *WeatherWarningReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("namespace", req.Namespace, "weatherwarning", req.Name)
	log.Info("Start of Reconcile")

	log.Info("Getting WeatherWarning resource")
	var ww weatherappv1alpha1.WeatherWarning
	if err := r.Get(ctx, req.NamespacedName, &ww); err != nil {
		log.Error(err, "unable to fetch WeatherWarning")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	maxtemperature := ww.Spec.MaxTemperature
	log.Info("current max comfort temperature set at", "maxtemperature", maxtemperature)

	log.Info("Getting CheckWeather resource")
	var cw weatherappv1alpha1.CheckWeather
	if err := r.Get(ctx, "checkweather", &cw); err != nil {
		log.Error(err, "unable to fetch CheckWeather")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	current_temp := int32(cw.Status.Temperature)
	log.Info("current temperature", "current", current_temp)

	// if current_temp < comfort_temp {
	// 	log.Info("DANGER: Too cold")
	// } else {
	// 	log.Info("nothing to see here")
	// }
	//
	// log.Info("Updating the status of the resource")
	// if err := r.Status().Update(ctx, &ww); err != nil {
	// 	log.Error(err, "Error updating status")
	// 	return ctrl.Result{}, err
	// }

	// Don't requeue, future changes will trigger Reconcile
	return ctrl.Result{}, nil
}

func (r *WeatherWarningReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&weatherappv1alpha1.WeatherWarning{}).
		Complete(r)
}
