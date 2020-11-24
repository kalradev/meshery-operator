/*
Copyright 2020 Layer5, Inc.

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

// MeshSyncSpec defines the desired state of MeshSync
type MeshSyncSpec struct {
	Cluster   string `json:"cluster,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Size      int32  `json:"size,omitempty"`
}

// MeshSyncStatus defines the observed state of MeshSync
type MeshSyncStatus struct {
	Conditions []Condition `json:"conditions,omitempty"`
}

// MeshSync is the Schema for the meshsyncs API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type MeshSync struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MeshSyncSpec   `json:"spec,omitempty"`
	Status MeshSyncStatus `json:"status,omitempty"`
}

// MeshSyncList contains a list of MeshSync
// +kubebuilder:object:root=true
type MeshSyncList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MeshSync `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MeshSync{}, &MeshSyncList{})
}
