package libraries

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LibraryInfo struct {
	ContainerName      *string `json:"containerName,omitempty"`
	CreatorId          *string `json:"creatorId,omitempty"`
	Name               *string `json:"name,omitempty"`
	Path               *string `json:"path,omitempty"`
	ProvisioningStatus *string `json:"provisioningStatus,omitempty"`
	Type               *string `json:"type,omitempty"`
	UploadedTimestamp  *string `json:"uploadedTimestamp,omitempty"`
}
