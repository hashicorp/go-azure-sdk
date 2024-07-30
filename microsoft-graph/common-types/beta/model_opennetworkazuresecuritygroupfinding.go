package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenNetworkAzureSecurityGroupFinding struct {
	CreatedDateTime *string                      `json:"createdDateTime,omitempty"`
	Id              *string                      `json:"id,omitempty"`
	InboundPorts    *InboundPorts                `json:"inboundPorts,omitempty"`
	ODataType       *string                      `json:"@odata.type,omitempty"`
	SecurityGroup   *AuthorizationSystemResource `json:"securityGroup,omitempty"`
	VirtualMachines *[]VirtualMachineDetails     `json:"virtualMachines,omitempty"`
}
