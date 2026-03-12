package kubeenvironments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ArcConfiguration struct {
	ArtifactStorageAccessMode    *string                `json:"artifactStorageAccessMode,omitempty"`
	ArtifactStorageClassName     *string                `json:"artifactStorageClassName,omitempty"`
	ArtifactStorageMountPath     *string                `json:"artifactStorageMountPath,omitempty"`
	ArtifactStorageNodeName      *string                `json:"artifactStorageNodeName,omitempty"`
	ArtifactsStorageType         *StorageType           `json:"artifactsStorageType,omitempty"`
	FrontEndServiceConfiguration *FrontEndConfiguration `json:"frontEndServiceConfiguration,omitempty"`
	KubeConfig                   *string                `json:"kubeConfig,omitempty"`
}
