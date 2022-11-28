package namespacesnetworksecurityperimeterconfigurations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkSecurityPerimeterConfigurationProperties struct {
	NetworkSecurityPerimeter *NetworkSecurityPerimeter                                           `json:"networkSecurityPerimeter"`
	Profile                  *NetworkSecurityPerimeterConfigurationPropertiesProfile             `json:"profile"`
	ProvisioningIssues       *[]ProvisioningIssue                                                `json:"provisioningIssues,omitempty"`
	ProvisioningState        *NetworkSecurityPerimeterConfigurationProvisioningState             `json:"provisioningState,omitempty"`
	ResourceAssociation      *NetworkSecurityPerimeterConfigurationPropertiesResourceAssociation `json:"resourceAssociation"`
}
