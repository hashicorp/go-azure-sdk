package replicationmigrationitems

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VMwareCbtUpdateDiskInput struct {
	DiskId         string  `json:"diskId"`
	IsOSDisk       *string `json:"isOSDisk,omitempty"`
	TargetDiskName *string `json:"targetDiskName,omitempty"`
}
