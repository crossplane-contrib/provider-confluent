// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

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

type ClusterConfigCredentialsInitParameters struct {
}

type ClusterConfigCredentialsObservation struct {
}

type ClusterConfigCredentialsParameters struct {

	// The Kafka API Key.
	// The Cluster API Key for your Confluent Cloud cluster.
	// +kubebuilder:validation:Required
	KeySecretRef v1.SecretKeySelector `json:"keySecretRef" tf:"-"`

	// The Kafka API Secret.
	// The Cluster API Secret for your Confluent Cloud cluster.
	// +kubebuilder:validation:Required
	SecretSecretRef v1.SecretKeySelector `json:"secretSecretRef" tf:"-"`
}

type ClusterConfigInitParameters struct {

	// The custom cluster settings to set:
	// The custom cluster settings to set (e.g., `"num.partitions" = "8"`).
	Config map[string]*string `json:"config,omitempty" tf:"config,omitempty"`

	// supports the following:
	// The Cluster API Credentials.
	Credentials []ClusterConfigCredentialsInitParameters `json:"credentials,omitempty" tf:"credentials,omitempty"`

	// supports the following:
	KafkaCluster []ClusterConfigKafkaClusterInitParameters `json:"kafkaCluster,omitempty" tf:"kafka_cluster,omitempty"`

	// The REST endpoint of the Dedicated Kafka cluster, for example, https://pkc-00000.us-central1.gcp.confluent.cloud:443).
	// The REST endpoint of the Kafka cluster (e.g., `https://pkc-00000.us-central1.gcp.confluent.cloud:443`).
	RestEndpoint *string `json:"restEndpoint,omitempty" tf:"rest_endpoint,omitempty"`
}

type ClusterConfigKafkaClusterInitParameters struct {

	// The ID of the Dedicated Kafka cluster, for example, lkc-abc123.
	// The Kafka cluster ID (e.g., `lkc-12345`).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ClusterConfigKafkaClusterObservation struct {

	// The ID of the Dedicated Kafka cluster, for example, lkc-abc123.
	// The Kafka cluster ID (e.g., `lkc-12345`).
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type ClusterConfigKafkaClusterParameters struct {

	// The ID of the Dedicated Kafka cluster, for example, lkc-abc123.
	// The Kafka cluster ID (e.g., `lkc-12345`).
	// +kubebuilder:validation:Optional
	ID *string `json:"id" tf:"id,omitempty"`
}

type ClusterConfigObservation struct {

	// The custom cluster settings to set:
	// The custom cluster settings to set (e.g., `"num.partitions" = "8"`).
	Config map[string]*string `json:"config,omitempty" tf:"config,omitempty"`

	// supports the following:
	// The Cluster API Credentials.
	Credentials []ClusterConfigCredentialsParameters `json:"credentials,omitempty" tf:"credentials,omitempty"`

	// The ID of the Dedicated Kafka cluster, for example, lkc-abc123.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// supports the following:
	KafkaCluster []ClusterConfigKafkaClusterObservation `json:"kafkaCluster,omitempty" tf:"kafka_cluster,omitempty"`

	// The REST endpoint of the Dedicated Kafka cluster, for example, https://pkc-00000.us-central1.gcp.confluent.cloud:443).
	// The REST endpoint of the Kafka cluster (e.g., `https://pkc-00000.us-central1.gcp.confluent.cloud:443`).
	RestEndpoint *string `json:"restEndpoint,omitempty" tf:"rest_endpoint,omitempty"`
}

type ClusterConfigParameters struct {

	// The custom cluster settings to set:
	// The custom cluster settings to set (e.g., `"num.partitions" = "8"`).
	// +kubebuilder:validation:Optional
	Config map[string]*string `json:"config,omitempty" tf:"config,omitempty"`

	// supports the following:
	// The Cluster API Credentials.
	// +kubebuilder:validation:Optional
	Credentials []ClusterConfigCredentialsParameters `json:"credentials,omitempty" tf:"credentials,omitempty"`

	// supports the following:
	// +kubebuilder:validation:Optional
	KafkaCluster []ClusterConfigKafkaClusterParameters `json:"kafkaCluster,omitempty" tf:"kafka_cluster,omitempty"`

	// The REST endpoint of the Dedicated Kafka cluster, for example, https://pkc-00000.us-central1.gcp.confluent.cloud:443).
	// The REST endpoint of the Kafka cluster (e.g., `https://pkc-00000.us-central1.gcp.confluent.cloud:443`).
	// +kubebuilder:validation:Optional
	RestEndpoint *string `json:"restEndpoint,omitempty" tf:"rest_endpoint,omitempty"`
}

// ClusterConfigSpec defines the desired state of ClusterConfig
type ClusterConfigSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ClusterConfigParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider ClusterConfigInitParameters `json:"initProvider,omitempty"`
}

// ClusterConfigStatus defines the observed state of ClusterConfig.
type ClusterConfigStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ClusterConfigObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterConfig is the Schema for the ClusterConfigs API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,confluent}
type ClusterConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.config) || (has(self.initProvider) && has(self.initProvider.config))",message="spec.forProvider.config is a required parameter"
	Spec   ClusterConfigSpec   `json:"spec"`
	Status ClusterConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterConfigList contains a list of ClusterConfigs
type ClusterConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterConfig `json:"items"`
}

// Repository type metadata.
var (
	ClusterConfig_Kind             = "ClusterConfig"
	ClusterConfig_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ClusterConfig_Kind}.String()
	ClusterConfig_KindAPIVersion   = ClusterConfig_Kind + "." + CRDGroupVersion.String()
	ClusterConfig_GroupVersionKind = CRDGroupVersion.WithKind(ClusterConfig_Kind)
)

func init() {
	SchemeBuilder.Register(&ClusterConfig{}, &ClusterConfigList{})
}
