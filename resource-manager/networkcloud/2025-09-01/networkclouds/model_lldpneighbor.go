package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LldpNeighbor struct {
	PortDescription   *string `json:"portDescription,omitempty"`
	PortName          *string `json:"portName,omitempty"`
	SystemDescription *string `json:"systemDescription,omitempty"`
	SystemName        *string `json:"systemName,omitempty"`
}
