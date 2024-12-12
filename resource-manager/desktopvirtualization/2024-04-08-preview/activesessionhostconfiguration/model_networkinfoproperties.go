package activesessionhostconfiguration

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkInfoProperties struct {
	SecurityGroupId *string `json:"securityGroupId,omitempty"`
	SubnetId        string  `json:"subnetId"`
}
