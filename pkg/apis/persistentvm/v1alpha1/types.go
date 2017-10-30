package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=pvm

// PVM persistent VM custom resource
type PVM struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PVMSpec   `json:"spec,omitempty"`
	Status PVMStatus `json:"status"`
}

// PVMSpec is the data inside the custom resource
type PVMSpec struct {
	Device string `json:"Device"`
	Name   string `json:"Name"`
}

// PVMStatus contains the runtime status of the VMS
type PVMStatus struct {
	RunningVMS int32 `json:"runningVMS"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=pvms

// PVMList is a list of VMConfigs.
type PVMList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []PVM `json:"items"`
}
