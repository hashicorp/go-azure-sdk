package lab

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LabUpdateProperties struct {
	AutoShutdownProfile   *AutoShutdownProfile   `json:"autoShutdownProfile"`
	ConnectionProfile     *ConnectionProfile     `json:"connectionProfile"`
	Description           *string                `json:"description,omitempty"`
	LabPlanId             *string                `json:"labPlanId,omitempty"`
	RosterProfile         *RosterProfile         `json:"rosterProfile"`
	SecurityProfile       *SecurityProfile       `json:"securityProfile"`
	Title                 *string                `json:"title,omitempty"`
	VirtualMachineProfile *VirtualMachineProfile `json:"virtualMachineProfile"`
}
