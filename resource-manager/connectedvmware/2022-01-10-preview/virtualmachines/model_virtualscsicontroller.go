package virtualmachines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualSCSIController struct {
	BusNumber          *int64              `json:"busNumber,omitempty"`
	ControllerKey      *int64              `json:"controllerKey,omitempty"`
	ScsiCtlrUnitNumber *int64              `json:"scsiCtlrUnitNumber,omitempty"`
	Sharing            *VirtualSCSISharing `json:"sharing,omitempty"`
	Type               *SCSIControllerType `json:"type,omitempty"`
}
