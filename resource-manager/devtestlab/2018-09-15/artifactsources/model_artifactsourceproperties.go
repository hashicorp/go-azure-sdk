package artifactsources

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ArtifactSourceProperties struct {
	ArmTemplateFolderPath *string            `json:"armTemplateFolderPath,omitempty"`
	BranchRef             *string            `json:"branchRef,omitempty"`
	CreatedDate           *string            `json:"createdDate,omitempty"`
	DisplayName           *string            `json:"displayName,omitempty"`
	FolderPath            *string            `json:"folderPath,omitempty"`
	ProvisioningState     *string            `json:"provisioningState,omitempty"`
	SecurityToken         *string            `json:"securityToken,omitempty"`
	SourceType            *SourceControlType `json:"sourceType,omitempty"`
	Status                *EnableStatus      `json:"status,omitempty"`
	UniqueIdentifier      *string            `json:"uniqueIdentifier,omitempty"`
	Uri                   *string            `json:"uri,omitempty"`
}
