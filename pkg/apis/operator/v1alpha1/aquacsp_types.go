package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AquaCspSpec defines the desired state of AquaCsp

type AquaCspSpec struct {
	Infrastructure *AquaInfrastructure `json:"infra,omitempty"`
	Common         *AquaCommon         `json:"common,omitempty"`

	RegistryData *AquaDockerRegistry      `json:"registry,omitempty"`
	ExternalDb   *AquaDatabaseInformation `json:"externalDb,omitempty"`

	DbService      *AquaService         `json:"database,omitempty"`
	GatewayService *AquaService         `json:"gateway,required"`
	ServerService  *AquaService         `json:"server,required"`
	ScannerService *AquaService         `json:"scanner,omitempty"`
	Scale          *AquaScannerCliScale `json:"scale,omitempty"`

	LicenseToken  string `json:"licenseToken,omitempty"`
	AdminPassword string `json:"adminPassword,omitempty"`
}

// AquaCspStatus defines the observed state of AquaCsp
type AquaCspStatus struct {
	Phase string              `json:"phase"`
	State AquaDeploymentState `json:"state"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AquaCsp is the Schema for the aquacsps API
// +k8s:openapi-gen=true
type AquaCsp struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AquaCspSpec   `json:"spec,omitempty"`
	Status AquaCspStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AquaCspList contains a list of AquaCsp
type AquaCspList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AquaCsp `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AquaCsp{}, &AquaCspList{})
}
