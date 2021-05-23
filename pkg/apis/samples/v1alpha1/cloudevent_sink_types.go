/*
Copyright 2020 waveywaves

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
	"knative.dev/pkg/apis"
	duckv1 "knative.dev/pkg/apis/duck/v1"
	"knative.dev/pkg/kmeta"
)

// CloudeventSink is a Knative abstraction that encapsulates the interface by which Knative
// components express a desire to have a particular image cached.
//
// +genclient
// +genreconciler
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type CloudeventSink struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// Spec holds the desired state of the CloudeventSink (from the client).
	// +optional
	Spec CloudeventSinkSpec `json:"spec,omitempty"`

	// Status communicates the observed state of the CloudeventSink (from the controller).
	// +optional
	Status CloudeventSinkStatus `json:"status,omitempty"`
}

var (
	// Check that AddressableService can be validated and defaulted.
	_ apis.Validatable   = (*CloudeventSink)(nil)
	_ apis.Defaultable   = (*CloudeventSink)(nil)
	_ kmeta.OwnerRefable = (*CloudeventSink)(nil)
	// Check that the type conforms to the duck Knative Resource shape.
	_ duckv1.KRShaped = (*CloudeventSink)(nil)
)

// CloudeventSinkSpec holds the desired state of the CloudeventSink (from the client).
type CloudeventSinkSpec struct {
	SinkType string `json:"type,omitempty"`
}

const (
	// SimpleDeploymentConditionReady is set when the revision is starting to materialize
	// runtime resources, and becomes true when those resources are ready.
	SimpleDeploymentConditionReady = apis.ConditionReady
)

// CloudeventSinkStatus communicates the observed state of the CloudeventSink (from the controller).
type CloudeventSinkStatus struct {
	duckv1.Status `json:",inline"`

	ReadyReplicas int32 `json:"readyReplicas"`
}

// CloudeventSinkList is a list of AddressableService resources
//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type CloudeventSinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []CloudeventSink `json:"items"`
}

// GetStatus retrieves the status of the resource. Implements the KRShaped interface.
func (d *CloudeventSink) GetStatus() *duckv1.Status {
	return &d.Status.Status
}
