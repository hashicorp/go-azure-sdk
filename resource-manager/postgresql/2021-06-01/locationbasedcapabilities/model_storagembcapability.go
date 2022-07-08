package locationbasedcapabilities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StorageMBCapability struct {
	Name          *string `json:"name,omitempty"`
	Status        *string `json:"status,omitempty"`
	StorageSizeMB *int64  `json:"storageSizeMB,omitempty"`
	SupportedIops *int64  `json:"supportedIops,omitempty"`
}
