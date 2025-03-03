package standbyvirtualmachines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StandbyVirtualMachineResourceProperties struct {
	ProvisioningState        *ProvisioningState `json:"provisioningState,omitempty"`
	VirtualMachineResourceId string             `json:"virtualMachineResourceId"`
}
