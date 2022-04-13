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

// VAVerifyReconciler will verify the VerticaArchive in the VBU CR exists
type VAVerifyReconciler struct {
	VBURec *VerticaBackupReconciler
	Vbu    *vapi.VerticaBackup
	Varc   *vapi.VerticaArchive
}

func MakeVAVerifyReconciler(r *VerticaBackupReconciler, vbu *vapi.VerticaBackup) controllers.ReconcileActor {
	return &VAVerifyReconciler{VBURec: r, Vbu: vbu, Varc: &vapi.VerticaArchive{}}
}

// Reconcile will verify the VerticaArchive in the VBU CR exists
func (s *VAVerifyReconciler) Reconcile(ctx context.Context, req *ctrl.Request) (ctrl.Result, error) {
	varcNm := controllers.GetNamespacedName(s.Vbu.Namespace, s.Vbu.Spec.Archive)
	// This reconciler is to get afeedback if the VerticaArchive that is referenced in the vbu doesn't exist.
	// This will print out an event if the VerticaArchive cannot be found.
	return fetchVdbOrVa(ctx, s.VBURec, s.Vbu, varcNm, s.Varc)
}
