package firewallrules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateFirewallRuleProperties struct {
	EndIpAddress   *string `json:"endIpAddress,omitempty"`
	StartIpAddress *string `json:"startIpAddress,omitempty"`
}
