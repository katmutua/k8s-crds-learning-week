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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	StatePending  = "PENDING"
	StateFinished = "FINISHED"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CheckWeatherSpec defines the desired state of CheckWeather
type CheckWeatherSpec struct {
	City string `json:"city,omitempty"`
}

// CheckWeatherStatus defines the observed state of CheckWeather
type CheckWeatherStatus struct {
	State       string `json:"state,omitempty"`
	Temperature int32  `json:"temperature,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// CheckWeather is the Schema for the checkweathers API
type CheckWeather struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CheckWeatherSpec   `json:"spec,omitempty"`
	Status            CheckWeatherStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CheckWeatherList contains a list of CheckWeather
type CheckWeatherList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CheckWeather `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CheckWeather{}, &CheckWeatherList{})
}
