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
	"k8s.io/apimachinery/pkg/types"
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
// +kubebuilder:rbac:groups=weather-app.example.com,resources=checkweathers,verbs=get;list;watch
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
	temperatureThreshold := ww.Spec.TempThreshHold
	log.Info("current max comfort temperature set at", "temperatureThreshold", temperatureThreshold)

	log.Info("fetching CheckWeather resource")
	var cw weatherappv1alpha1.CheckWeather

	// build the NamespacedName from the value specified from the the WeatherWarning
	cv := types.NamespacedName{Name: ww.Spec.TargetCheckName, Namespace: ww.Spec.TargetCheckNamespace}
	if err := r.Get(ctx, cv, &cw); err != nil {
		log.Error(err, "unable to fetch CheckWeather Object")
		return ctrl.Result{}, err
	}

	if cw.Status.State == "" {
		cw.Status.State = weatherappv1alpha1.StatePending
	}
	temp := cw.Status.Temperature
	switch cw.Status.State {
	case weatherappv1alpha1.StateFinished:
		log.Info("current temperature", "temp", temp)
		//handle case where the status has not changed
	default:
		//handle case
	}
	//Don't requeue, future changes will trigger Reconcile
	return ctrl.Result{}, nil
}

func (r *WeatherWarningReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&weatherappv1alpha1.WeatherWarning{}).
		Complete(r)
}
