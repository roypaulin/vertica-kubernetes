/*
 (c) Copyright [2021-2023] Open Text.
 Licensed under the Apache License, Version 2.0 (the "License");
 You may not use this file except in compliance with the License.
 You may obtain a copy of the License at

 http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package meta

import "strconv"

const (
	// Annotations that we set in each of the pod.  These are set by the
	// AnnotateAndLabelPodReconciler.  They are available in the pod with the
	// downwardAPI so they can be picked up by the Vertica data collector (DC).
	KubernetesVersionAnnotation   = "kubernetes.io/version"   // Version of the k8s server
	KubernetesGitCommitAnnotation = "kubernetes.io/gitcommit" // Git commit of the k8s server
	KubernetesBuildDateAnnotation = "kubernetes.io/buildDate" // Build date of the k8s server

	// If this annotation is on any CR, the operator will skip processing. This can
	// be used to avoid getting in an infinity error-retry loop. Or, if you know
	// no additional work will ever exist for an object. Just set this to a
	// true|ON|1 value.
	PauseOperatorAnnotation = "vertica.com/pause"

	// This is a feature flag for using vertica without admintools. Set this
	// annotation in the VerticaDB that you want to use the new vclusterOps
	// library for any vertica admin task. The value of this annotation is
	// treated as a boolean.
	VClusterOpsAnnotation      = "vertica.com/vcluster-ops"
	VClusterOpsAnnotationTrue  = "true"
	VClusterOpsAnnotationFalse = "false"

	// This is a feature flag for accessing the secrets configured in Google Secret Manager.
	// The value of this annotation is treated as a boolean.
	GcpGsmAnnotation = "vertica.com/use-gcp-secret-manager"

	// Two annotations that are set by the operator when creating objects.
	OperatorDeploymentMethodAnnotation = "vertica.com/operator-deployment-method"
	OperatorVersionAnnotation          = "vertica.com/operator-version"

	// This is an indicator added to the VerticaDB annotation that the CR was
	// converted from a different version of the API. The API it was converted
	// from is the value of the annotation.
	APIConversionAnnotation = "vertica.com/converted-from-api-version"

	// Ignore the cluster lease when doing a revive or start_db.  Use this with
	// caution, as ignoring the cluster lease when another system is using the
	// same communal storage will cause corruption.
	IgnoreClusterLeaseAnnotation = "vertica.com/ignore-cluster-lease"

	// When set to False, this parameter will ensure that when changing the
	// vertica version that we follow the upgrade path.  The Vertica upgrade
	// path means you cannot downgrade a Vertica release, nor can you skip any
	// released Vertica versions when upgrading.
	IgnoreUpgradePathAnnotation = "vertica.com/ignore-upgrade-path"

	// The timeout, in seconds, to use when the operator restarts a node or the
	// entire cluster.  If omitted, we use the default timeout of 20 minutes.
	RestartTimeoutAnnotation = "vertica.com/restart-timeout"

	// Annotations that we add by parsing vertica --version output
	VersionAnnotation   = "vertica.com/version"
	BuildDateAnnotation = "vertica.com/buildDate"
	BuildRefAnnotation  = "vertica.com/buildRef"
	// Annotation for the database's revive_instance_id
	ReviveInstanceIDAnnotation = "vertica.com/revive-instance-id"
)

// IsPauseAnnotationSet will check the annotations for a special value that will
// pause the operator for the CR.
func IsPauseAnnotationSet(annotations map[string]string) bool {
	return lookupBoolAnnotation(annotations, PauseOperatorAnnotation)
}

// UseVClusterOps returns true if all admin commands should use the vclusterOps
// library rather than admintools.
func UseVClusterOps(annotations map[string]string) bool {
	return lookupBoolAnnotation(annotations, VClusterOpsAnnotation)
}

// UseGCPSecretManager returns true if access to the communal secret should go through
// Google's secret manager rather the fetching the secret from k8s meta-data.
func UseGCPSecretManager(annotations map[string]string) bool {
	return lookupBoolAnnotation(annotations, GcpGsmAnnotation)
}

// IgnoreClusterLease returns true if revive/start should ignore the cluster lease
func IgnoreClusterLease(annotations map[string]string) bool {
	return lookupBoolAnnotation(annotations, IgnoreClusterLeaseAnnotation)
}

// IgnoreUpgradePath returns true if the upgrade path can be ignored when
// changing images.
func IgnoreUpgradePath(annotations map[string]string) bool {
	return lookupBoolAnnotation(annotations, IgnoreUpgradePathAnnotation)
}

// GetRestartTimeout returns the timeout to use for restart node or start db. If
// 0 is returned, this means to use the default.
func GetRestartTimeout(annotations map[string]string) int {
	return lookupIntAnnotation(annotations, RestartTimeoutAnnotation)
}

// lookupBoolAnnotation is a helper function to lookup a specific annotation and
// treat it as if it were a boolean.
func lookupBoolAnnotation(annotations map[string]string, annotation string) bool {
	defaultValue := false
	if val, ok := annotations[annotation]; ok {
		varAsBool, err := strconv.ParseBool(val)
		if err != nil {
			return defaultValue
		}
		return varAsBool
	}
	return defaultValue
}

func lookupIntAnnotation(annotations map[string]string, annotation string) int {
	const defaultValue = 0
	if val, ok := annotations[annotation]; ok {
		varAsInt, err := strconv.ParseInt(val, 10, 0)
		if err != nil {
			return defaultValue
		}
		return int(varAsInt)
	}
	return defaultValue
}
