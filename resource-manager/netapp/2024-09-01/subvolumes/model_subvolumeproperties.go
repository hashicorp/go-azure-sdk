package subvolumes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubvolumeProperties struct {
	ParentPath        *string `json:"parentPath,omitempty"`
	Path              *string `json:"path,omitempty"`
	ProvisioningState *string `json:"provisioningState,omitempty"`
	Size              *int64  `json:"size,omitempty"`
}
