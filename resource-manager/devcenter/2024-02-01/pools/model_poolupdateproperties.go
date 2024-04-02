package pools

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PoolUpdateProperties struct {
	DevBoxDefinitionName         *string                        `json:"devBoxDefinitionName,omitempty"`
	DisplayName                  *string                        `json:"displayName,omitempty"`
	LicenseType                  *LicenseType                   `json:"licenseType,omitempty"`
	LocalAdministrator           *LocalAdminStatus              `json:"localAdministrator,omitempty"`
	ManagedVirtualNetworkRegions *[]string                      `json:"managedVirtualNetworkRegions,omitempty"`
	NetworkConnectionName        *string                        `json:"networkConnectionName,omitempty"`
	SingleSignOnStatus           *SingleSignOnStatus            `json:"singleSignOnStatus,omitempty"`
	StopOnDisconnect             *StopOnDisconnectConfiguration `json:"stopOnDisconnect,omitempty"`
	VirtualNetworkType           *VirtualNetworkType            `json:"virtualNetworkType,omitempty"`
}
