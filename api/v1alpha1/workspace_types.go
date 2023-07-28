/*
Copyright 2023.

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
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ResourceSpec struct {
	// The number of required GPU nodes.
	Count *int `json:"count,omitempty"`

	// The required instance type of the GPU node.
	InstanceType string `json:"instanceType,omitempty"`

	// The required label for the GPU node.
	LabelSelector *metav1.LabelSelector `json:"labelSelector,omitempty"`

	// The existing GPU nodes with the required labels and the required instancetype.
	// This field is used when the number of qualified existing nodes is larger than the required count.
	// Users need to ensure supported VHD images are installed in the VMs.
	Nodes []string `json:"nodes,omitempty"`
}

type PresetModelName string

const (
	PresetSetModelllama2A            PresetModelName = "llama2-7b"
	PresetSetModelllama2B            PresetModelName = "llama2-13b"
	PresetSetModelllama2C            PresetModelName = "llama2-70b"
	PresetSetModelStableDiffusionXXX PresetModelName = "stablediffusion-xxx"
)

type PresetModelSpec struct {
	// Name of a supported preset model, e.g., llama2-7b.
	Name PresetModelName `json:"name,omitempty"`
	// The custom volume that will be mounted to the pod running preset models.
	// Later, we may limit to AzureFile and configmap in API.
	Volume []v1.Volume `json:"volume,omitempty"`
}

type InferenceSpec struct {
	// The preset model to be deployed.
	Preset PresetModelSpec `json:"preset,omitempty"`
	// The Pod template used by the Deployment. Users can use custom image and Pod spec.
	// Leave this filed unset if preset model is used.
	Template v1.PodTemplateSpec `json:"template,omitempty"`
}

type TrainingSpec struct {
	// Job pytorchJob `json:"job,omitempty"`
}

// WorkspaceStatus defines the observed state of Workspace
type WorkspaceStatus struct {
	WorkerNodes    []string `json:"workernodes,omitempty"`
	ResourceStatus string   `json:"resourceStatus,omitempty"`
	// Determined by the status of the created deployment1
	WorkloadStatus string `json:"workloadStatus,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Workspace is the Schema for the workspaces API
type Workspace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Resource  ResourceSpec  `json:"resource,omitempty"`
	Inference InferenceSpec `json:"inference,omitempty"`
	Training  TrainingSpec  `json:"training,omitempty"`

	Status WorkspaceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// WorkspaceList contains a list of Workspace
type WorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workspace `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Workspace{}, &WorkspaceList{})
}
