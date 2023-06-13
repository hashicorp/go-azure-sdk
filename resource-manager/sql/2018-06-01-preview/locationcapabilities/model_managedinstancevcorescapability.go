package locationcapabilities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedInstanceVcoresCapability struct {
	IncludedMaxSize       *MaxSizeCapability        `json:"includedMaxSize,omitempty"`
	InstancePoolSupported *bool                     `json:"instancePoolSupported,omitempty"`
	Name                  *string                   `json:"name,omitempty"`
	Reason                *string                   `json:"reason,omitempty"`
	StandaloneSupported   *bool                     `json:"standaloneSupported,omitempty"`
	Status                *CapabilityStatus         `json:"status,omitempty"`
	SupportedStorageSizes *[]MaxSizeRangeCapability `json:"supportedStorageSizes,omitempty"`
	Value                 *int64                    `json:"value,omitempty"`
}
