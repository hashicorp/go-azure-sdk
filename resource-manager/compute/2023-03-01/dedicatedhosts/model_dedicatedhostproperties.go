package dedicatedhosts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DedicatedHostProperties struct {
	AutoReplaceOnFailure *bool                      `json:"autoReplaceOnFailure,omitempty"`
	HostId               *string                    `json:"hostId,omitempty"`
	InstanceView         *DedicatedHostInstanceView `json:"instanceView,omitempty"`
	LicenseType          *DedicatedHostLicenseTypes `json:"licenseType,omitempty"`
	PlatformFaultDomain  *int64                     `json:"platformFaultDomain,omitempty"`
	ProvisioningState    *string                    `json:"provisioningState,omitempty"`
	ProvisioningTime     *string                    `json:"provisioningTime,omitempty"`
	TimeCreated          *string                    `json:"timeCreated,omitempty"`
	VirtualMachines      *[]SubResourceReadOnly     `json:"virtualMachines,omitempty"`
}
