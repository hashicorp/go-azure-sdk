package provisionednetworks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisionedNetworkProperties struct {
	AddressPrefix     *string                              `json:"addressPrefix,omitempty"`
	NetworkType       *ProvisionedNetworkTypes             `json:"networkType,omitempty"`
	ProvisioningState *ProvisionedNetworkProvisioningState `json:"provisioningState,omitempty"`
}
