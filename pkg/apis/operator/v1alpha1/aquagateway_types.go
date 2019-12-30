package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AquaGatewaySpec defines the desired state of AquaGateway
type AquaGatewaySpec struct {
	Infrastructure *AquaInfrastructure `json:"infra"`
	Common         *AquaCommon         `json:"common"`

	GatewayService *AquaService             `json:"deploy,required"`
	ExternalDb     *AquaDatabaseInformation `json:"externalDb,omitempty"`
}

// AquaGatewayStatus defines the observed state of AquaGateway
type AquaGatewayStatus struct {
	Nodes []string            `json:"nodes"`
	State AquaDeploymentState `json:"state"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AquaGateway is the Schema for the aquagateways API
// +k8s:openapi-gen=true
type AquaGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AquaGatewaySpec   `json:"spec,omitempty"`
	Status AquaGatewayStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AquaGatewayList contains a list of AquaGateway
type AquaGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AquaGateway `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AquaGateway{}, &AquaGatewayList{})
}
