package externalsecuritysolutions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CefSolutionProperties struct {
	Agent             *string             `json:"agent,omitempty"`
	DeviceType        *string             `json:"deviceType,omitempty"`
	DeviceVendor      *string             `json:"deviceVendor,omitempty"`
	Hostname          *string             `json:"hostname,omitempty"`
	LastEventReceived *string             `json:"lastEventReceived,omitempty"`
	Workspace         *ConnectedWorkspace `json:"workspace,omitempty"`
}
