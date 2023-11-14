/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type RoleBindingInitParameters struct {

	// A Confluent Resource Name(CRN) that specifies the scope and resource patterns necessary for the role to bind.
	// A CRN that specifies the scope and resource patterns necessary for the role to bind.
	CrnPattern *string `json:"crnPattern,omitempty" tf:"crn_pattern,omitempty"`

	// A principal User to bind the role to, for example, "User:u-111aaa" for binding to a user "u-111aaa", or "User:sa-111aaa" for binding to a service account "sa-111aaa".
	// The principal User to bind the role to.
	Principal *string `json:"principal,omitempty" tf:"principal,omitempty"`

	// A name of the role to bind to the principal. See Confluent Cloud RBAC Roles for a full list of supported role names.
	// The name of the role to bind to the principal.
	RoleName *string `json:"roleName,omitempty" tf:"role_name,omitempty"`
}

type RoleBindingObservation struct {

	// A Confluent Resource Name(CRN) that specifies the scope and resource patterns necessary for the role to bind.
	// A CRN that specifies the scope and resource patterns necessary for the role to bind.
	CrnPattern *string `json:"crnPattern,omitempty" tf:"crn_pattern,omitempty"`

	// The ID of the Role Binding (e.g., rb-f3a90de).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A principal User to bind the role to, for example, "User:u-111aaa" for binding to a user "u-111aaa", or "User:sa-111aaa" for binding to a service account "sa-111aaa".
	// The principal User to bind the role to.
	Principal *string `json:"principal,omitempty" tf:"principal,omitempty"`

	// A name of the role to bind to the principal. See Confluent Cloud RBAC Roles for a full list of supported role names.
	// The name of the role to bind to the principal.
	RoleName *string `json:"roleName,omitempty" tf:"role_name,omitempty"`
}

type RoleBindingParameters struct {

	// A Confluent Resource Name(CRN) that specifies the scope and resource patterns necessary for the role to bind.
	// A CRN that specifies the scope and resource patterns necessary for the role to bind.
	// +kubebuilder:validation:Optional
	CrnPattern *string `json:"crnPattern,omitempty" tf:"crn_pattern,omitempty"`

	// A principal User to bind the role to, for example, "User:u-111aaa" for binding to a user "u-111aaa", or "User:sa-111aaa" for binding to a service account "sa-111aaa".
	// The principal User to bind the role to.
	// +kubebuilder:validation:Optional
	Principal *string `json:"principal,omitempty" tf:"principal,omitempty"`

	// A name of the role to bind to the principal. See Confluent Cloud RBAC Roles for a full list of supported role names.
	// The name of the role to bind to the principal.
	// +kubebuilder:validation:Optional
	RoleName *string `json:"roleName,omitempty" tf:"role_name,omitempty"`
}

// RoleBindingSpec defines the desired state of RoleBinding
type RoleBindingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoleBindingParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider RoleBindingInitParameters `json:"initProvider,omitempty"`
}

// RoleBindingStatus defines the observed state of RoleBinding.
type RoleBindingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoleBindingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// RoleBinding is the Schema for the RoleBindings API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,confluent}
type RoleBinding struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.crnPattern) || has(self.initProvider.crnPattern)",message="crnPattern is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.principal) || has(self.initProvider.principal)",message="principal is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.roleName) || has(self.initProvider.roleName)",message="roleName is a required parameter"
	Spec   RoleBindingSpec   `json:"spec"`
	Status RoleBindingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleBindingList contains a list of RoleBindings
type RoleBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoleBinding `json:"items"`
}

// Repository type metadata.
var (
	RoleBinding_Kind             = "RoleBinding"
	RoleBinding_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RoleBinding_Kind}.String()
	RoleBinding_KindAPIVersion   = RoleBinding_Kind + "." + CRDGroupVersion.String()
	RoleBinding_GroupVersionKind = CRDGroupVersion.WithKind(RoleBinding_Kind)
)

func init() {
	SchemeBuilder.Register(&RoleBinding{}, &RoleBindingList{})
}
