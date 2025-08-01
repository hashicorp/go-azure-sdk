package computerps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineNetworkInterfaceIPConfiguration struct {
	Name       string                                                   `json:"name"`
	Properties *VirtualMachineNetworkInterfaceIPConfigurationProperties `json:"properties,omitempty"`
}
