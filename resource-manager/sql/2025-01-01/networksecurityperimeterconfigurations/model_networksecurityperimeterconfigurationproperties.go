package networksecurityperimeterconfigurations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkSecurityPerimeterConfigurationProperties struct {
	NetworkSecurityPerimeter *NSPConfigPerimeter     `json:"networkSecurityPerimeter,omitempty"`
	Profile                  *NSPConfigProfile       `json:"profile,omitempty"`
	ProvisioningIssues       *[]NSPProvisioningIssue `json:"provisioningIssues,omitempty"`
	ProvisioningState        *string                 `json:"provisioningState,omitempty"`
	ResourceAssociation      *NSPConfigAssociation   `json:"resourceAssociation,omitempty"`
}
