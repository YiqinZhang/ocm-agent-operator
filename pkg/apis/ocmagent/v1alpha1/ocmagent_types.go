package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OcmAgentSpec defines the desired state of OcmAgent
type OcmAgentSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// OcmBaseUrl defines the OCM api endpoint for OCM agent to access
	OcmBaseUrl string `json:"ocmBaseUrl"`

	// Services defines the supported OCM services, eg, service_log, cluster_management
	Services []string `json:"services"`

	// OcmAgentImage defines the image which will be used by the OCM Agent
	OcmAgentImage string `json:"ocmAgentImage"`

	// TokenSecret points to the secret name which stores the access token to OCM server
	TokenSecret string `json:"tokenSecret"`

	// Replicas defines the replica count for the OCM Agent service
	Replicas int32 `json:"replicas"`

	// OcmAgentConfig defines the configmap name which generated by operator and consumed by OCM Agent service
	OcmAgentConfig string `json:"ocmAgentConfig"`
}

// OcmAgentStatus defines the observed state of OcmAgent
type OcmAgentStatus struct {
	// INSERT ADDITIONAL STATUS FIELDS - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// ServiceStatus indicates the status of OCM Agent service
	ServiceStatus string `json:"serviceStatus"`

	AvailableReplicas int32 `json:"availableReplicas"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OcmAgent is the Schema for the ocmagents API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=ocmagents,scope=Namespaced
type OcmAgent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OcmAgentSpec   `json:"spec,omitempty"`
	Status OcmAgentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// OcmAgentList contains a list of OcmAgent
type OcmAgentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OcmAgent `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OcmAgent{}, &OcmAgentList{})
}
