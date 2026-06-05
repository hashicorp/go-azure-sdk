package virtualnetworkappliances

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommonNetworkSecurityGroupPropertiesFormat struct {
	DefaultSecurityRules *[]CommonSecurityRule     `json:"defaultSecurityRules,omitempty"`
	FlowLogs             *[]CommonFlowLog          `json:"flowLogs,omitempty"`
	FlushConnection      *bool                     `json:"flushConnection,omitempty"`
	NetworkInterfaces    *[]CommonNetworkInterface `json:"networkInterfaces,omitempty"`
	ProvisioningState    *ProvisioningState        `json:"provisioningState,omitempty"`
	ResourceGuid         *string                   `json:"resourceGuid,omitempty"`
	SecurityRules        *[]CommonSecurityRule     `json:"securityRules,omitempty"`
	Subnets              *[]CommonSubnet           `json:"subnets,omitempty"`
}
