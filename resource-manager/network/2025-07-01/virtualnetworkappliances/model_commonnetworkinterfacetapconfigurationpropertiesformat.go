package virtualnetworkappliances

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommonNetworkInterfaceTapConfigurationPropertiesFormat struct {
	ProvisioningState *ProvisioningState       `json:"provisioningState,omitempty"`
	VirtualNetworkTap *CommonVirtualNetworkTap `json:"virtualNetworkTap,omitempty"`
}
