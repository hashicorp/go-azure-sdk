package activesessionhostconfiguration

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityInfoProperties struct {
	SecureBootEnabled *bool                       `json:"secureBootEnabled,omitempty"`
	Type              *VirtualMachineSecurityType `json:"type,omitempty"`
	VTpmEnabled       *bool                       `json:"vTpmEnabled,omitempty"`
}
