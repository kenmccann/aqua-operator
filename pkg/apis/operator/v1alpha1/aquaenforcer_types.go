package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AquaEnforcerSpec defines the desired state of AquaEnforcer
type AquaEnforcerSpec struct {
	Infrastructure *AquaInfrastructure `json:"infra"`
	Common         *AquaCommon         `json:"common"`

	EnforcerService *AquaService            `json:"deploy,required"`
	Gateway         *AquaGatewayInformation `json:"gateway,required"`
	Token           string                  `json:"token,required"`
	Secret          *AquaSecret             `json:"secret,required"`
}

// AquaEnforcerStatus defines the observed state of AquaEnforcer
type AquaEnforcerStatus struct {
	State AquaDeploymentState `json:"state"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AquaEnforcer is the Schema for the aquaenforcers API
// +k8s:openapi-gen=true
type AquaEnforcer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AquaEnforcerSpec   `json:"spec,omitempty"`
	Status AquaEnforcerStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AquaEnforcerList contains a list of AquaEnforcer
type AquaEnforcerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AquaEnforcer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AquaEnforcer{}, &AquaEnforcerList{})
}
