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

// VerticaRestoreSpec defines the desired state of VerticaRestore
type VerticaRestoreSpec struct {
	// Important: Run "make" to regenerate code after modifying this file

	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	// The name of the VerticaDB we are going to restore to.
	// This parameter is required.
	// The name must be an existing VerticaDB in the same namespace.
	VerticaDBName string `json:"verticaDBName,omitempty"`

	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors="urn:alm:descriptor:com.tectonic.ui:text"
	// The name of the VerticaArchive to restore from. It must reference an existing VerticaArchive,
	// and at least one backup must exist there.
	Archive string `json:"archive,omitempty"`

	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// The timestamp of the image to restore from.  This must be one of the retained backup images.
	// If this is omitted, the last backup is used.
	Timestamp metav1.Timestamp `json:"timestamp,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default:=false
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors="urn:alm:descriptor:com.tectonic.ui:booleanSwitch"
	// When set to true, all foreign key constraints are unconditionally dropped during object-level restore.
	// This allows you to restore database objects independent of their foreign key dependencies.
	// You must set objectRestoreMode to coexist, otherwise this setting is ignored.
	DropForeignConstraints bool `json:"dropForeignConstraints,omitempty"`

	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// List of objects to restore from the backup. This is a wildcard.
	IncludeObjects string `json:"includeObjects,omitempty"`

	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// List of objects to exclude from the backup. This can only be used if includeObjects is used,
	// as it remove objects from that were included in that wildcard.
	ExcludeObjects string `json:"excludeObjects,omitempty"`

	// +kubebuilder:validation:Optional
	// +kubebuilder:default:=createOrReplace
	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// Specifies how to handle objects of the same name when restoring schema or table backups.
	// Valid values are one of the following:
	//	- createOrReplace: creates any objects that do not exist. If an object does exist,
	//		restore overwrites it with the version from the archive.
	//	- create: creates any objects that do not exist and does not replace existing objects.
	//	  	If an object being restored does exist, the restore fails.
	//	- coexist: vbr creates the restored version of each object with a name formatted as follows:
	//		backup_timestamp_objectname

	//		This approach allows existing and restored objects to exist simultaneously.
	//		If the appended information pushes the schema name past the maximum length of 128 characters,
	//		Vertica truncates the name. You can perform a reverse lookup of the original schema name
	//		by querying the system table TRUNCATED_SCHEMATA.
	ObjectRestoreMode string `json:"objectRestoreMode,omitempty"`
}

type VerticaRestorePhase string

const (
	// RestoreInitialized means that all required resources have been initialized.
	RestoreInitialized VerticaRestorePhase = "Initialized"
	// RestoreRunning means that the Restore process has initiated .
	RestoreRunning VerticaRestorePhase = "Running"
	// RestoreSucceeded means that the Restore process has completed.
	RestoreSucceeded VerticaRestorePhase = "Succeeded"
	// RestoreFailed means that the Restore failed.
	RestoreFailed VerticaRestorePhase = "Failed"
)

// VerticaRestoreCondition describes the observed state of a VerticaRestore at a certain point.
type VerticaRestoreCondition struct {
	// +operator-sdk:csv:customresourcedefinitions:type=status
	// Type is the type of the condition
	Type VerticaRestorePhase `json:"type"`

	// +operator-sdk:csv:customresourcedefinitions:type=status
	// Status is the status of the condition
	// can be True, False or Unknown
	Status corev1.ConditionStatus `json:"status"`

	// +operator-sdk:csv:customresourcedefinitions:type=status
	// Last time the condition transitioned from one status to another.
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
}

// VerticaRestoreStatus defines the observed state of VerticaRestore
type VerticaRestoreStatus struct {

	// +operator-sdk:csv:customresourcedefinitions:type=status
	// Current phase of the restore operation
	Phase VerticaRestorePhase `json:"phase"`

	// +operator-sdk:csv:customresourcedefinitions:type=status
	// Reason indicates the reason for any restore related failures.
	Reason string `json:"reason,omitempty"`

	// +operator-sdk:csv:customresourcedefinitions:type=status
	// The name of the VerticaArchive used for this restore
	Archive string `json:"archive,omitempty"`

	// +operator-sdk:csv:customresourcedefinitions:type=status
	// Conditions for VerticaRestore
	Conditions []VerticaRestoreCondition `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:resource:categories=all;vertica,shortName=vr
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
//+kubebuilder:printcolumn:name="Phase",type="string",JSONPath=".status.phase"
//+kubebuilder:printcolumn:name="Archive",type="string",JSONPath=".status.archive"
//+operator-sdk:csv:customresourcedefinitions:resources={{VerticaDB,vertica.com/v1beta1,""},{VerticaArchive,vertica.com/v1beta1,""}}

// VerticaRestore is the schema for verticarestores API
type VerticaRestore struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VerticaRestoreSpec   `json:"spec,omitempty"`
	Status VerticaRestoreStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// VerticaRestoreList contains a list of VerticaRestore
type VerticaRestoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VerticaRestore `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VerticaRestore{}, &VerticaRestoreList{})
}
