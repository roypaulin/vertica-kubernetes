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

package archive

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	vapi "github.com/vertica/vertica-kubernetes/api/v1beta1"
	"github.com/vertica/vertica-kubernetes/pkg/controllers"
	verrors "github.com/vertica/vertica-kubernetes/pkg/errors"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// VerticaArchiveReconciler reconciles a VerticaArchive object
type VerticaArchiveReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Log    logr.Logger
	EVRec  record.EventRecorder
}

//+kubebuilder:rbac:groups=vertica.com,namespace=WATCH_NAMESPACE,resources=verticaarchives,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=vertica.com,namespace=WATCH_NAMESPACE,resources=verticaarhives/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=vertica.com,namespace=WATCH_NAMESPACE,resources=verticaarchives/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.11.0/pkg/reconcile
func (r *VerticaArchiveReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := r.Log.WithValues("verticaarchive", req.NamespacedName)
	log.Info("starting reconcile of VerticaArchive")

	var res ctrl.Result
	varc := &vapi.VerticaArchive{}
	err := r.Get(ctx, req.NamespacedName, varc)
	if err != nil {
		if errors.IsNotFound(err) {
			log.Info("VerticaArchive resource not found.  Ignoring since object must be deleted")
			return ctrl.Result{}, nil
		}
		log.Error(err, "failed to get VerticaArchive")
		return ctrl.Result{}, err
	}

	// The actors that will be applied, in sequence, to reconcile a varc.
	actors := []controllers.ReconcileActor{}

	// Iterate over each actor
	for _, act := range actors {
		log.Info("starting actor", "name", fmt.Sprintf("%T", act))
		res, err = act.Reconcile(ctx, &req)
		// Error or a request to requeue will stop the reconciliation.
		if verrors.IsReconcileAborted(res, err) {
			log.Info("aborting reconcile of VerticaArchive", "result", res, "err", err)
			return res, err
		}
	}

	log.Info("ending reconcile of VerticaArchive", "result", res, "err", err)
	return res, err
}

// SetupWithManager sets up the controller with the Manager.
func (r *VerticaArchiveReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&vapi.VerticaArchive{}).
		Complete(r)
}
