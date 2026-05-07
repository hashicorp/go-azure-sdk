package openapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IPAddressGroup struct {
	DeliveryRegion *string          `json:"deliveryRegion,omitempty"`
	IPv4Addresses  *[]CidrIPAddress `json:"ipv4Addresses,omitempty"`
	IPv6Addresses  *[]CidrIPAddress `json:"ipv6Addresses,omitempty"`
}
