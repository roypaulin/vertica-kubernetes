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
package vb

import (
	"context"

	vapi "github.com/vertica/vertica-kubernetes/api/v1beta1"
	"github.com/vertica/vertica-kubernetes/pkg/controllers"
	ctrl "sigs.k8s.io/controller-runtime"
)

// VDBVerifyReconciler will verify the VerticaDB in the VBU CR exists
type VDBVerifyReconciler struct {
	VBURec *VerticaBackupReconciler
	Vbu    *vapi.VerticaBackup
	Vdb    *vapi.VerticaDB
}

func MakeVDBVerifyReconciler(r *VerticaBackupReconciler, vbu *vapi.VerticaBackup) controllers.ReconcileActor {
	return &VDBVerifyReconciler{VBURec: r, Vbu: vbu, Vdb: &vapi.VerticaDB{}}
}

// Reconcile will verify the VerticaDB in the VBU CR exists
func (s *VDBVerifyReconciler) Reconcile(ctx context.Context, req *ctrl.Request) (ctrl.Result, error) {
	vdbNm := controllers.GetNamespacedName(s.Vbu.Namespace, s.Vbu.Spec.VerticaDBName)
	// This reconciler is to get a feedback if the VerticaDB that is referenced in the vbu doesn't exist.
	// This will print out an event if the VerticaDB cannot be found.
	return fetchVdbOrVa(ctx, s.VBURec, s.Vbu, vdbNm, s.Vdb)
}
