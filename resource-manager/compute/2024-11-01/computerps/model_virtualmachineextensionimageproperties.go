package computerps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineExtensionImageProperties struct {
	ComputeRole                string `json:"computeRole"`
	HandlerSchema              string `json:"handlerSchema"`
	OperatingSystem            string `json:"operatingSystem"`
	SupportsMultipleExtensions *bool  `json:"supportsMultipleExtensions,omitempty"`
	VMScaleSetEnabled          *bool  `json:"vmScaleSetEnabled,omitempty"`
}
