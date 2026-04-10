package virtualnetworktap

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommonRouteTablePropertiesFormat struct {
	DisableBgpRoutePropagation *bool              `json:"disableBgpRoutePropagation,omitempty"`
	ProvisioningState          *ProvisioningState `json:"provisioningState,omitempty"`
	ResourceGuid               *string            `json:"resourceGuid,omitempty"`
	Routes                     *[]CommonRoute     `json:"routes,omitempty"`
	Subnets                    *[]CommonSubnet    `json:"subnets,omitempty"`
}
