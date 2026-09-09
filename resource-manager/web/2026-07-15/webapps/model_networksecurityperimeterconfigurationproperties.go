package webapps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkSecurityPerimeterConfigurationProperties struct {
	NetworkSecurityPerimeter *NetworkSecurityPerimeter `json:"networkSecurityPerimeter,omitempty"`
	Profile                  *NspProfile               `json:"profile,omitempty"`
	ProvisioningIssues       *[]NspProvisioningIssue   `json:"provisioningIssues,omitempty"`
	ProvisioningState        *string                   `json:"provisioningState,omitempty"`
	ResourceAssociations     *NspResourceAssociation   `json:"resourceAssociations,omitempty"`
}
