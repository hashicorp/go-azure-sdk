package flexibleservercapabilities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StorageMbCapability struct {
	DefaultIopsTier    *string                  `json:"defaultIopsTier,omitempty"`
	Reason             *string                  `json:"reason,omitempty"`
	Status             *CapabilityStatus        `json:"status,omitempty"`
	StorageSizeMb      *int64                   `json:"storageSizeMb,omitempty"`
	SupportedIops      *int64                   `json:"supportedIops,omitempty"`
	SupportedIopsTiers *[]StorageTierCapability `json:"supportedIopsTiers,omitempty"`
}
