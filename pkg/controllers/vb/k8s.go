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

package vb

import (
	"context"
	"fmt"

	vapi "github.com/vertica/vertica-kubernetes/api/v1beta1"
	"github.com/vertica/vertica-kubernetes/pkg/events"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// fetchVdbOrVa will fetch the VerticaDB/VerticaArchive that is referenced in a VerticaBackup.
// This will log an event if the VerticaDB/VerticaArchive is not found.
func fetchVdbOrVa(ctx context.Context, vburec *VerticaBackupReconciler,
	vbu *vapi.VerticaBackup, nm types.NamespacedName, obj client.Object) (ctrl.Result, error) {
	err := vburec.Client.Get(ctx, nm, obj)
	if err != nil && errors.IsNotFound(err) {
		event := ""
		objType := ""
		ownedObjName := ""
		switch v := obj.(type) {
		case *vapi.VerticaDB:
			event = events.VerticaDBNotFound
			objType = vapi.VerticaDBKind
			ownedObjName = vbu.Spec.VerticaDBName
		case *vapi.VerticaArchive:
			event = events.VerticaArchiveNotFound
			objType = vapi.VerticaArchiveKind
			ownedObjName = vbu.Spec.Archive
		default:
			event = events.ObjectNotFound
			objType = fmt.Sprintf("%T", v)
			ownedObjName = "Unknown"
		}
		vburec.EVRec.Eventf(vbu, corev1.EventTypeWarning, event,
			"The '%s' named '%s' was not found", objType, ownedObjName)
		return ctrl.Result{Requeue: true}, nil
	}
	return ctrl.Result{}, err
}
