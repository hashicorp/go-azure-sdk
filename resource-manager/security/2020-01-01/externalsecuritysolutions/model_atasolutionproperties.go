package externalsecuritysolutions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AtaSolutionProperties struct {
	DeviceType        *string             `json:"deviceType,omitempty"`
	DeviceVendor      *string             `json:"deviceVendor,omitempty"`
	LastEventReceived *string             `json:"lastEventReceived,omitempty"`
	Workspace         *ConnectedWorkspace `json:"workspace,omitempty"`
}
