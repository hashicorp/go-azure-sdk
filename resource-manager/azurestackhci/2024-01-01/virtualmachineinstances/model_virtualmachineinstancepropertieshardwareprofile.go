package virtualmachineinstances

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineInstancePropertiesHardwareProfile struct {
	DynamicMemoryConfig *VirtualMachineInstancePropertiesHardwareProfileDynamicMemoryConfig `json:"dynamicMemoryConfig,omitempty"`
	MemoryMB            *int64                                                              `json:"memoryMB,omitempty"`
	Processors          *int64                                                              `json:"processors,omitempty"`
	VMSize              *VMSizeEnum                                                         `json:"vmSize,omitempty"`
}
