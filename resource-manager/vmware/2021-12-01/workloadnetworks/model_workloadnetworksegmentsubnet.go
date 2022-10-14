package workloadnetworks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkloadNetworkSegmentSubnet struct {
	DhcpRanges     *[]string `json:"dhcpRanges,omitempty"`
	GatewayAddress *string   `json:"gatewayAddress,omitempty"`
}
