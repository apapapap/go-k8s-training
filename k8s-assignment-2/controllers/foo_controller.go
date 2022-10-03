/*
Copyright 2022.

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

package controllers

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	k8sassignment2v1 "my.domain/k8s-assignment-2/api/v1"
)

// FooReconciler reconciles a Foo object
type FooReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=k8s-assignment-2.my.domain,resources=foos,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=k8s-assignment-2.my.domain,resources=foos/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=k8s-assignment-2.my.domain,resources=foos/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Foo object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.13.0/pkg/reconcile
func (r *FooReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)
	log.Info("reconciling foo custom resource")

	// Get the Foo resource that triggered the reconciliation request
	var foo k8sassignment2v1.Foo
	if err := r.Get(ctx, req.NamespacedName, &foo); err != nil {
		log.Error(err, "unable to fetch Foo")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Get pods with the same name as Foo's friend
	var podList corev1.PodList
	var friendFound bool
	if err := r.List(ctx, &podList); err != nil {
		log.Error(err, "unable to list pods")
	} else {
		for _, item := range podList.Items {
			if item.GetName() == foo.Spec.Name {
				log.Info("pod linked to a foo custom resource found", "name", item.GetName())
				friendFound = true
			}
		}
	}

	// Update Foo' happy status
	foo.Status.Happy = friendFound
	if err := r.Status().Update(ctx, &foo); err != nil {
		log.Error(err, "unable to update foo's happy status", "status", friendFound)
		return ctrl.Result{}, err
	}
	log.Info("foo's happy status updated", "status", friendFound)

	log.Info("foo custom resource reconciled")
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *FooReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&k8sassignment2v1.Foo{}).
		Complete(r)
}
