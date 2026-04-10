package networksecuritygroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommonIPamPoolPrefixAllocation struct {
	AllocatedAddressPrefixes *[]string                           `json:"allocatedAddressPrefixes,omitempty"`
	NumberOfIPAddresses      *string                             `json:"numberOfIpAddresses,omitempty"`
	Pool                     *CommonIPamPoolPrefixAllocationPool `json:"pool,omitempty"`
}
