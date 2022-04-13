/*
 (c) Copyright [2021-2022] Micro Focus or one of its affiliates.
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

//nolint:dupl
package vr

import (
	"context"

	vapi "github.com/vertica/vertica-kubernetes/api/v1beta1"
	"github.com/vertica/vertica-kubernetes/pkg/controllers"
	ctrl "sigs.k8s.io/controller-runtime"
)

// VDBVerifyReconciler will verify the VerticaDB in the VR CR exists
type VDBVerifyReconciler struct {
	VRRec *VerticaRestoreReconciler
	Vr    *vapi.VerticaRestore
	Vdb   *vapi.VerticaDB
}

func MakeVDBVerifyReconciler(r *VerticaRestoreReconciler, vr *vapi.VerticaRestore) controllers.ReconcileActor {
	return &VDBVerifyReconciler{VRRec: r, Vr: vr, Vdb: &vapi.VerticaDB{}}
}

// Reconcile will verify the VerticaDB in the VR CR exists
func (s *VDBVerifyReconciler) Reconcile(ctx context.Context, req *ctrl.Request) (ctrl.Result, error) {
	vdbNm := controllers.GetNamespacedName(s.Vr.Namespace, s.Vr.Spec.VerticaDBName)
	// This reconciler is to get a feedback if the VerticaDB that is referenced in the vr doesn't exist.
	// This will print out an event if the VerticaDB cannot be found.
	return fetchVdbOrVa(ctx, s.VRRec, s.Vr, vdbNm, s.Vdb)
}
