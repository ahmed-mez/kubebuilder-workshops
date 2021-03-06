/*
Copyright 2019 Google LLC

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// GuestBookSpec defines the desired state of GuestBook
type GuestBookSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	Frontend FrontendSpec `json:"frontend"`

	RedisName string `json:"redisName"`
}

type FrontendSpec struct {
	// +optional
	Resources corev1.ResourceRequirements `json:"resources,omitempty"`
	// +optional
	ServingPort int32 `json:"servingPort,omitempty"`

	// +optional
	Replicas *int32 `json:"replicas,omitempty"` // If it's nil, we'll use a default
}

// GuestBookStatus defines the observed state of GuestBook
type GuestBookStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	URL string `json:"url,omitempty"`

	Conditions []StatusCondition `json:"conditions,omitempty"`
}

type ConditionStatus string

var (
	ConditionStatusHealthy   ConditionStatus = "Healthy"
	ConditionStatusUnhealthy ConditionStatus = "Unhealthy"
	ConditionStatusUnknown   ConditionStatus = "Unknown"
)

type StatusCondition struct {
	Type   string          `json:"type"`
	Status ConditionStatus `json:"status"`
	// +optional
	LastProbeTime metav1.Time `json:"lastProbeTime,omitempty"`
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
	// +optional
	Reason string `json:"reason,omitempty"`
	// +optional
	Message string `json:"message,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:priority=0,name=URL,type=string,JSONPath=".status.url",description="GuestBook Frontend URL",format=""
// +kubebuilder:printcolumn:priority=0,name=Deployment,type=string,JSONPath=".status.conditions[?(@.type==\"DeploymentUpToDate\")].status",description="Is the Deployment Up-To-Date",format=""
// +kubebuilder:printcolumn:priority=0,name=Service,type=string,JSONPath=".status.conditions[?(@.type==\"ServiceUpToDate\")].status",description="Is the Service Up-To-Date",format=""

// GuestBook is the Schema for the guestbooks API
type GuestBook struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GuestBookSpec   `json:"spec,omitempty"`
	Status GuestBookStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// GuestBookList contains a list of GuestBook
type GuestBookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GuestBook `json:"items"`
}

func init() {
	SchemeBuilder.Register(&GuestBook{}, &GuestBookList{})
}
