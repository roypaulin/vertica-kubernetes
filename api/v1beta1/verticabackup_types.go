/*
Copyright [2021-2022] Micro Focus or one of its affiliates.
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

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VerticaBackupSpec defines the desired state of VerticaBackup
type VerticaBackupSpec struct {
	// Important: Run "make" to regenerate code after modifying this file

	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	// The name of the VerticaDB that has the database we will backup.
	// This parameter is required.
	// The name must be an existing VerticaDB in the same namespace.
	VerticaDBName string `json:"verticaDBName,omitempty"`

	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	// The name of the VerticaArchive to use for this backup. This parameter is required.
	// The archive must already exist in the same namespace.
	Archive string `json:"archive,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default:=true
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors="urn:alm:descriptor:com.tectonic.ui:hidden"
	// When set to true, the data sent to the backup path will be encrypted.
	// This is required for certain backup locations.
	EncryptTransport bool `json:"encryptTransport,omitempty"`

	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// The path  to a SSL certificate bundle. This path is relative inside a VerticaDB pod.
	// These are typically one of the certs that are mounted at /certs/<secretName>/<key>
	CaFile string `json:"caFile,omitempty"`

	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// List of objects to include in the backup.
	IncludeObjects string `json:"includeObjects,omitempty"`

	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// List of objects to exclude in the backup. This can only be used if includeObjects is used,
	// as it remove objects from that were included in that wildcard.
	ExcludeObjects string `json:"excludeObjects,omitempty"`
}

type VerticaBackupPhase string

const (
	// BackupInitialized means that all required resources have been initialized
	BackupInitialized VerticaBackupPhase = "Initialized"
	// BackupRunning means that the backup pod has been created and the backup process initiated on the infinispan server.
	BackupRunning VerticaBackupPhase = "Running"
	// BackupSucceeded means that the backup process has completed.
	BackupSucceeded VerticaBackupPhase = "Succeeded"
	// BackupFailed means that the backup failed.
	BackupFailed VerticaBackupPhase = "Failed"
)

// BackupCondition describes the observed state of a Backup at a certain point.
type VerticaBackupCondition struct {
	// +operator-sdk:csv:customresourcedefinitions:type=status
	// Type is the type of the condition
	Type VerticaBackupPhase `json:"type"`

	// +operator-sdk:csv:customresourcedefinitions:type=status
	// Status is the status of the condition
	// can be True, False or Unknown
	Status corev1.ConditionStatus `json:"status"`

	// +operator-sdk:csv:customresourcedefinitions:type=status
	// Last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
}

// VerticaBackupStatus defines the observed state of Vertica
type VerticaBackupStatus struct {
	// +operator-sdk:csv:customresourcedefinitions:type=status
	// Status message for the current backup.
	BackupStatus string `json:"backupStatus,omitempty"`

	// +operator-sdk:csv:customresourcedefinitions:type=status
	// Reason indicates the reason for any backup related failures.
	Reason string `json:"reason,omitempty"`

	// +operator-sdk:csv:customresourcedefinitions:type=status
	// Current phase of the backup operation
	Phase VerticaBackupPhase `json:"phase,omitempty"`

	// +operator-sdk:csv:customresourcedefinitions:type=status
	// The name of the VerticaArchive to use for this backup
	Archive string `json:"archive,omitempty"`

	// +operator-sdk:csv:customresourcedefinitions:type=status
	// Conditions for VerticaDB
	Conditions []VerticaBackupCondition `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:resource:shortName=vbu
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
//+kubebuilder:printcolumn:name="Phase",type="string",JSONPath=".status.phase"
//+kubebuilder:printcolumn:name="Archive",type="string",JSONPath=".status.archive"
//+operator-sdk:csv:customresourcedefinitions:resources={{VerticaDB,vertica.com/v1beta1,""},{VerticaArchive,vertica.com/v1beta1,""}}

// VerticaBackup is the schema for verticabackups API
type VerticaBackup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VerticaBackupSpec   `json:"spec,omitempty"`
	Status VerticaBackupStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// VerticaArchiveList contains a list of VerticaArchive
type VerticaBackupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VerticaBackup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VerticaBackup{}, &VerticaBackupList{})
}
