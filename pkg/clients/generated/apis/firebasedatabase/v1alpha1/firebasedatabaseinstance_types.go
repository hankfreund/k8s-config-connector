// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Config Connector and manual
//     changes will be clobbered when the file is regenerated.
//
// ----------------------------------------------------------------------------

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type FirebaseDatabaseInstanceSpec struct {
	/* The intended database state. */
	// +optional
	DesiredState *string `json:"desiredState,omitempty"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. A reference to the region where the Firebase Realtime database resides.
	Check all [available regions](https://firebase.google.com/docs/projects/locations#rtdb-locations). */
	Region string `json:"region"`

	/* Immutable. Optional. The instanceId of the resource. Used for creation and acquisition. When unset, the value of `metadata.name` is used as the default. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`

	/* Immutable. The database type.
	Each project can create one default Firebase Realtime Database, which cannot be deleted once created.
	Creating user Databases is only available for projects on the Blaze plan.
	Projects can be upgraded using the Cloud Billing API https://cloud.google.com/billing/reference/rest/v1/projects/updateBillingInfo. Default value: "USER_DATABASE" Possible values: ["DEFAULT_DATABASE", "USER_DATABASE"]. */
	// +optional
	Type *string `json:"type,omitempty"`
}

type FirebaseDatabaseInstanceStatus struct {
	/* Conditions represent the latest available observations of the
	   FirebaseDatabaseInstance's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* The database URL in the form of https://{instance-id}.firebaseio.com for us-central1 instances
	or https://{instance-id}.{region}.firebasedatabase.app in other regions. */
	// +optional
	DatabaseUrl *string `json:"databaseUrl,omitempty"`

	/* The fully-qualified resource name of the Firebase Realtime Database, in the
	format: projects/PROJECT_NUMBER/locations/REGION_IDENTIFIER/instances/INSTANCE_ID
	PROJECT_NUMBER: The Firebase project's ['ProjectNumber'](https://firebase.google.com/docs/reference/firebase-management/rest/v1beta1/projects#FirebaseProject.FIELDS.project_number)
	Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510). */
	// +optional
	Name *string `json:"name,omitempty"`

	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int `json:"observedGeneration,omitempty"`

	/* The current database state. Set desired_state to :DISABLED to disable the database and :ACTIVE to reenable the database. */
	// +optional
	State *string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FirebaseDatabaseInstance is the Schema for the firebasedatabase API
// +k8s:openapi-gen=true
type FirebaseDatabaseInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FirebaseDatabaseInstanceSpec   `json:"spec,omitempty"`
	Status FirebaseDatabaseInstanceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FirebaseDatabaseInstanceList contains a list of FirebaseDatabaseInstance
type FirebaseDatabaseInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirebaseDatabaseInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FirebaseDatabaseInstance{}, &FirebaseDatabaseInstanceList{})
}