package subvolumes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubvolumeModelProperties struct {
	AccessedTimeStamp *string `json:"accessedTimeStamp,omitempty"`
	BytesUsed         *int64  `json:"bytesUsed,omitempty"`
	ChangedTimeStamp  *string `json:"changedTimeStamp,omitempty"`
	CreationTimeStamp *string `json:"creationTimeStamp,omitempty"`
	ModifiedTimeStamp *string `json:"modifiedTimeStamp,omitempty"`
	ParentPath        *string `json:"parentPath,omitempty"`
	Path              *string `json:"path,omitempty"`
	Permissions       *string `json:"permissions,omitempty"`
	ProvisioningState *string `json:"provisioningState,omitempty"`
	Size              *int64  `json:"size,omitempty"`
}
