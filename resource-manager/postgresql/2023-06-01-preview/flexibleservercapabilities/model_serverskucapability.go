package flexibleservercapabilities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerSkuCapability struct {
	Name                      *string           `json:"name,omitempty"`
	Reason                    *string           `json:"reason,omitempty"`
	Status                    *CapabilityStatus `json:"status,omitempty"`
	SupportedHaMode           *[]HaMode         `json:"supportedHaMode,omitempty"`
	SupportedIops             *int64            `json:"supportedIops,omitempty"`
	SupportedMemoryPerVcoreMb *int64            `json:"supportedMemoryPerVcoreMb,omitempty"`
	SupportedZones            *[]string         `json:"supportedZones,omitempty"`
	VCores                    *int64            `json:"vCores,omitempty"`
}
