/*
Copyright 2023 Rakshit Gondwal.

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

package v1alpha1

import (
	"github.com/pkg/errors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
)

// log is for logging in this package.
var scalerlog = logf.Log.WithName("scaler-resource")

func (r *Scaler) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// TODO(user): EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
//+kubebuilder:webhook:path=/validate-api-rakshitgondwal-online-v1alpha1-scaler,mutating=false,failurePolicy=fail,sideEffects=None,groups=api.rakshitgondwal.online,resources=scalers,verbs=create;update,versions=v1alpha1,name=vscaler.kb.io,admissionReviewVersions=v1

var _ webhook.Validator = &Scaler{}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *Scaler) ValidateCreate() error {
	scalerlog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.
	return nil
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *Scaler) ValidateUpdate(old runtime.Object) error {
	scalerlog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return nil
}

// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *Scaler) ValidateDelete() error {
	scalerlog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}

func (r *Scaler) validateScaler() error {
	var allErrs field.ErrorList
	if err := r.validateScalerReplicas(); err != nil {
		allErrs = append(allErrs, err)
	}
	if len(allErrs) == 0 {
		return nil
	}

	return apierrors.NewInvalid(
		schema.GroupKind{Group: "batch.tutorial.kubebuilder.io", Kind: "Scaler"},
		r.Name, allErrs)
}

func (r *Scaler) validateScalerReplicas() *field.Error {
	if r.Spec.Replicas != 5 {
		return field.Invalid(
			field.NewPath("spec").Child("replicas"),
			r.Spec.Replicas,
			errors.New("Invalid number of replicas").Error(),
		)
	}
	return nil
}
