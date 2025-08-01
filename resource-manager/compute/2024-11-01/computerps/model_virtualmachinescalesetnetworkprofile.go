package computerps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineScaleSetNetworkProfile struct {
	HealthProbe                    *ApiEntityReference                           `json:"healthProbe,omitempty"`
	NetworkApiVersion              *NetworkApiVersion                            `json:"networkApiVersion,omitempty"`
	NetworkInterfaceConfigurations *[]VirtualMachineScaleSetNetworkConfiguration `json:"networkInterfaceConfigurations,omitempty"`
}
