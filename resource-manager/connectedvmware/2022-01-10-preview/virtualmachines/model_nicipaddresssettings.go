package virtualmachines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NicIPAddressSettings struct {
	AllocationMethod *string `json:"allocationMethod,omitempty"`
	IPAddress        *string `json:"ipAddress,omitempty"`
	SubnetMask       *string `json:"subnetMask,omitempty"`
}
