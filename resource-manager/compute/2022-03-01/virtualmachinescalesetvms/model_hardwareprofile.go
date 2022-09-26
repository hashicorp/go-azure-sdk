package virtualmachinescalesetvms

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HardwareProfile struct {
	VmSize           *VirtualMachineSizeTypes `json:"vmSize,omitempty"`
	VmSizeProperties *VmSizeProperties        `json:"vmSizeProperties,omitempty"`
}
