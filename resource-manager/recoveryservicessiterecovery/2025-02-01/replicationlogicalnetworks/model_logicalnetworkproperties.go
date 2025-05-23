package replicationlogicalnetworks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LogicalNetworkProperties struct {
	FriendlyName                    *string `json:"friendlyName,omitempty"`
	LogicalNetworkDefinitionsStatus *string `json:"logicalNetworkDefinitionsStatus,omitempty"`
	LogicalNetworkUsage             *string `json:"logicalNetworkUsage,omitempty"`
	NetworkVirtualizationStatus     *string `json:"networkVirtualizationStatus,omitempty"`
}
