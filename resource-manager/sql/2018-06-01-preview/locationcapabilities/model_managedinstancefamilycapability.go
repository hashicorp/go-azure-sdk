package locationcapabilities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedInstanceFamilyCapability struct {
	Name                  *string                            `json:"name,omitempty"`
	Reason                *string                            `json:"reason,omitempty"`
	Sku                   *string                            `json:"sku,omitempty"`
	Status                *CapabilityStatus                  `json:"status,omitempty"`
	SupportedLicenseTypes *[]LicenseTypeCapability           `json:"supportedLicenseTypes,omitempty"`
	SupportedVcoresValues *[]ManagedInstanceVcoresCapability `json:"supportedVcoresValues,omitempty"`
}
