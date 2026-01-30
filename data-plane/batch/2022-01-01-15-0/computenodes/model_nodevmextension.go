package computenodes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NodeVMExtension struct {
	InstanceView      *VMExtensionInstanceView `json:"instanceView,omitempty"`
	ProvisioningState *string                  `json:"provisioningState,omitempty"`
	VmExtension       *VmExtension             `json:"vmExtension,omitempty"`
}
