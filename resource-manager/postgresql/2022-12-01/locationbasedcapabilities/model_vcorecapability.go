package locationbasedcapabilities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VcoreCapability struct {
	Name                      *string `json:"name,omitempty"`
	Status                    *string `json:"status,omitempty"`
	SupportedIops             *int64  `json:"supportedIops,omitempty"`
	SupportedMemoryPerVcoreMB *int64  `json:"supportedMemoryPerVcoreMB,omitempty"`
	VCores                    *int64  `json:"vCores,omitempty"`
}
