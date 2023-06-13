package locationcapabilities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InstancePoolVcoresCapability struct {
	Name         *string            `json:"name,omitempty"`
	Reason       *string            `json:"reason,omitempty"`
	Status       *CapabilityStatus  `json:"status,omitempty"`
	StorageLimit *MaxSizeCapability `json:"storageLimit,omitempty"`
	Value        *int64             `json:"value,omitempty"`
}
