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

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

const VerticaArchiveKind = "VerticaArchive"

// VerticaArchiveSpec defines the desired state of VerticaArchive
type VerticaArchiveSpec struct {
	// Important: Run "make" to regenerate code after modifying this file

	// +operator-sdk:csv:customresourcedefinitions:type=spec
	// The path where the archive is stored. This parameter is required.
	// For S3-compatible or cloud locations, provide the bucket name and backup path.
	// For HDFS locations, provide the appropriate protocol and backup path.
	Path string `json:"path,omitempty"`

	// +kubebuilder:default:=1
	// +kubebuilder:validation:Optional
	// +operator-sdk:csv:customresourcedefinitions:type=spec,xDescriptors="urn:alm:descriptor:com.tectonic.ui:number"
	// Number of earlier backups to retain with the most recent backup.
	// If set to 1 (the default), we maintains two backups: the latest backup and the one before it
	RestorePointLimit int `json:"restorePointLimit"`
}

// VerticaArchiveStatus defines the observed state of VerticaArchive
type VerticaArchiveStatus struct {
	// +operator-sdk:csv:customresourcedefinitions:type=status
	// The number of Restore/Backup using the archive.
	ReferencedCount int `json:"referencedCount"`
}

//+kubebuilder:object:root=true
//+kubebuilder:resource:shortName=varc
//+kubebuilder:subresource:status
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
//+kubebuilder:printcolumn:name="Referenced Count",type="integer",JSONPath=".status.referencedCount"

// VerticaArchive is the schema for verticaarchives API
type VerticaArchive struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VerticaArchiveSpec   `json:"spec,omitempty"`
	Status VerticaArchiveStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// VerticaArchiveList contains a list of VerticaArchive
type VerticaArchiveList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VerticaArchive `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VerticaArchive{}, &VerticaArchiveList{})
}
