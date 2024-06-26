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

package v1beta1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/k8s/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type IndexAlgorithmConfig struct {
	/* Configuration options for using brute force search, which simply implements the
	standard linear search in the database for each query. */
	// +optional
	BruteForceConfig *IndexBruteForceConfig `json:"bruteForceConfig,omitempty"`

	/* Configuration options for using the tree-AH algorithm (Shallow tree + Asymmetric Hashing).
	Please refer to this paper for more details: https://arxiv.org/abs/1908.10396. */
	// +optional
	TreeAhConfig *IndexTreeAhConfig `json:"treeAhConfig,omitempty"`
}

type IndexBruteForceConfig struct {
}

type IndexConfig struct {
	/* The configuration with regard to the algorithms used for efficient search. */
	// +optional
	AlgorithmConfig *IndexAlgorithmConfig `json:"algorithmConfig,omitempty"`

	/* The default number of neighbors to find via approximate search before exact reordering is
	performed. Exact reordering is a procedure where results returned by an
	approximate search algorithm are reordered via a more expensive distance computation.
	Required if tree-AH algorithm is used. */
	// +optional
	ApproximateNeighborsCount *int64 `json:"approximateNeighborsCount,omitempty"`

	/* The number of dimensions of the input vectors. */
	Dimensions int64 `json:"dimensions"`

	/* The distance measure used in nearest neighbor search. The value must be one of the followings:
	* SQUARED_L2_DISTANCE: Euclidean (L_2) Distance
	* L1_DISTANCE: Manhattan (L_1) Distance
	* COSINE_DISTANCE: Cosine Distance. Defined as 1 - cosine similarity.
	* DOT_PRODUCT_DISTANCE: Dot Product Distance. Defined as a negative of the dot product. */
	// +optional
	DistanceMeasureType *string `json:"distanceMeasureType,omitempty"`

	/* Type of normalization to be carried out on each vector. The value must be one of the followings:
	* UNIT_L2_NORM: Unit L2 normalization type
	* NONE: No normalization type is specified. */
	// +optional
	FeatureNormType *string `json:"featureNormType,omitempty"`

	/* Immutable. Index data is split into equal parts to be processed. These are called "shards".
	The shard size must be specified when creating an index. The value must be one of the followings:
	* SHARD_SIZE_SMALL: Small (2GB)
	* SHARD_SIZE_MEDIUM: Medium (20GB)
	* SHARD_SIZE_LARGE: Large (50GB). */
	// +optional
	ShardSize *string `json:"shardSize,omitempty"`
}

type IndexMetadata struct {
	/* Immutable. The configuration of the Matching Engine Index. */
	// +optional
	Config *IndexConfig `json:"config,omitempty"`

	/* Allows creating or replacing the contents of the Matching Engine Index.
	When being updated, the existing content of the Index will be replaced by the data
	from the latest contentsDeltaUri.
	The string must be a valid Cloud Storage directory path. If this
	field is set when calling IndexService.UpdateIndex, then no other
	Index field can be also updated as part of the same call.
	The expected structure and format of the files this URI points to is
	described at https://cloud.google.com/vertex-ai/docs/matching-engine/using-matching-engine#input-data-format. */
	// +optional
	ContentsDeltaUri *string `json:"contentsDeltaUri,omitempty"`
}

type IndexTreeAhConfig struct {
	/* Number of embeddings on each leaf node. The default value is 1000 if not set. */
	// +optional
	LeafNodeEmbeddingCount *int64 `json:"leafNodeEmbeddingCount,omitempty"`

	/* The default percentage of leaf nodes that any query may be searched. Must be in
	range 1-100, inclusive. The default value is 10 (means 10%) if not set. */
	// +optional
	LeafNodesToSearchPercent *int64 `json:"leafNodesToSearchPercent,omitempty"`
}

type VertexAIIndexSpec struct {
	/* The description of the Index. */
	// +optional
	Description *string `json:"description,omitempty"`

	/* The display name of the Index. The name can be up to 128 characters long and can consist of any UTF-8 characters. */
	DisplayName string `json:"displayName"`

	/* Immutable. The update method to use with this Index. The value must be the followings. If not set, BATCH_UPDATE will be used by default.
	* BATCH_UPDATE: user can call indexes.patch with files on Cloud Storage of datapoints to update.
	* STREAM_UPDATE: user can call indexes.upsertDatapoints/DeleteDatapoints to update the Index and the updates will be applied in corresponding DeployedIndexes in nearly real-time. */
	// +optional
	IndexUpdateMethod *string `json:"indexUpdateMethod,omitempty"`

	/* An additional information about the Index. */
	// +optional
	Metadata *IndexMetadata `json:"metadata,omitempty"`

	/* The project that this resource belongs to. */
	ProjectRef v1alpha1.ResourceRef `json:"projectRef"`

	/* Immutable. The region of the index. eg us-central1. */
	Region string `json:"region"`

	/* Immutable. Optional. The service-generated name of the resource. Used for acquisition only. Leave unset to create a new resource. */
	// +optional
	ResourceID *string `json:"resourceID,omitempty"`
}

type IndexIndexStatsStatus struct {
	/* The number of shards in the Index. */
	// +optional
	ShardsCount *int64 `json:"shardsCount,omitempty"`

	/* The number of vectors in the Index. */
	// +optional
	VectorsCount *string `json:"vectorsCount,omitempty"`
}

type IndexObservedStateStatus struct {
	/* The timestamp of when the Index was created in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits. */
	// +optional
	CreateTime *string `json:"createTime,omitempty"`

	/* Stats of the index resource. */
	// +optional
	IndexStats []IndexIndexStatsStatus `json:"indexStats,omitempty"`

	/* Points to a YAML file stored on Google Cloud Storage describing additional information about the Index, that is specific to it. Unset if the Index does not have any additional information. */
	// +optional
	MetadataSchemaUri *string `json:"metadataSchemaUri,omitempty"`

	/* The resource name of the Index. */
	// +optional
	Name *string `json:"name,omitempty"`
}

type VertexAIIndexStatus struct {
	/* Conditions represent the latest available observations of the
	   VertexAIIndex's current state. */
	Conditions []v1alpha1.Condition `json:"conditions,omitempty"`
	/* ObservedGeneration is the generation of the resource that was most recently observed by the Config Connector controller. If this is equal to metadata.generation, then that means that the current reported status reflects the most recent desired state of the resource. */
	// +optional
	ObservedGeneration *int64 `json:"observedGeneration,omitempty"`

	/* The observed state of the underlying GCP resource. */
	// +optional
	ObservedState *IndexObservedStateStatus `json:"observedState,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:resource:categories=gcp,shortName=gcpvertexaiindex;gcpvertexaiindexes
// +kubebuilder:subresource:status
// +kubebuilder:metadata:labels="cnrm.cloud.google.com/managed-by-kcc=true";"cnrm.cloud.google.com/stability-level=stable";"cnrm.cloud.google.com/system=true";"cnrm.cloud.google.com/tf2crd=true"
// +kubebuilder:printcolumn:name="Age",JSONPath=".metadata.creationTimestamp",type="date"
// +kubebuilder:printcolumn:name="Ready",JSONPath=".status.conditions[?(@.type=='Ready')].status",type="string",description="When 'True', the most recent reconcile of the resource succeeded"
// +kubebuilder:printcolumn:name="Status",JSONPath=".status.conditions[?(@.type=='Ready')].reason",type="string",description="The reason for the value in 'Ready'"
// +kubebuilder:printcolumn:name="Status Age",JSONPath=".status.conditions[?(@.type=='Ready')].lastTransitionTime",type="date",description="The last transition time for the value in 'Status'"

// VertexAIIndex is the Schema for the vertexai API
// +k8s:openapi-gen=true
type VertexAIIndex struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VertexAIIndexSpec   `json:"spec,omitempty"`
	Status VertexAIIndexStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VertexAIIndexList contains a list of VertexAIIndex
type VertexAIIndexList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VertexAIIndex `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VertexAIIndex{}, &VertexAIIndexList{})
}
