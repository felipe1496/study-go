package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Database é a nossa CRD para gerenciar instâncias de banco de dados.
type Database struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DatabaseSpec   `json:"spec"`
	Status DatabaseStatus `json:"status,omitempty"`
}

// DatabaseSpec define o estado desejado do banco.
type DatabaseSpec struct {
	Engine  string `json:"engine"` // Ex: "postgres", "mysql"
	Version string `json:"version"`
	Storage string `json:"storage"`
}

// DatabaseStatus define o estado observado.
type DatabaseStatus struct {
	Phase string `json:"phase"` // Ex: "Ready", "Provisioning"
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// DatabaseList é uma lista de objetos Database (necessário para o client-gen).
type DatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Database `json:"items"`
}
