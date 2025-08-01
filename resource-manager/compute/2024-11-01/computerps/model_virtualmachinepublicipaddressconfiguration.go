package computerps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachinePublicIPAddressConfiguration struct {
	Name       string                                                `json:"name"`
	Properties *VirtualMachinePublicIPAddressConfigurationProperties `json:"properties,omitempty"`
	Sku        *PublicIPAddressSku                                   `json:"sku,omitempty"`
}
