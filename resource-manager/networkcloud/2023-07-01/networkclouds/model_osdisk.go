package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OsDisk struct {
	CreateOption *OsDiskCreateOption `json:"createOption,omitempty"`
	DeleteOption *OsDiskDeleteOption `json:"deleteOption,omitempty"`
	DiskSizeGB   int64               `json:"diskSizeGB"`
}
