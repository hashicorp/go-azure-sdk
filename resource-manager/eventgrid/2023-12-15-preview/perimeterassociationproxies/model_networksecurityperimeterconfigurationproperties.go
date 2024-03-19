package perimeterassociationproxies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkSecurityPerimeterConfigurationProperties struct {
	NetworkSecurityPerimeter *NetworkSecurityPerimeterInfo                    `json:"networkSecurityPerimeter,omitempty"`
	Profile                  *NetworkSecurityPerimeterConfigurationProfile    `json:"profile,omitempty"`
	ProvisioningIssues       *[]NetworkSecurityPerimeterConfigurationIssues   `json:"provisioningIssues,omitempty"`
	ProvisioningState        *NetworkSecurityPerimeterConfigProvisioningState `json:"provisioningState,omitempty"`
	ResourceAssociation      *ResourceAssociation                             `json:"resourceAssociation,omitempty"`
}
