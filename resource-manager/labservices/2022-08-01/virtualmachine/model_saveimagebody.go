package virtualmachine

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SaveImageBody struct {
	LabVirtualMachineId *string `json:"labVirtualMachineId,omitempty"`
	Name                *string `json:"name,omitempty"`
}
