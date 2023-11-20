package namespacesnetworksecurityperimeterconfigurations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkSecurityPerimeterConfigurationProperties struct {
	ApplicableFeatures       *[]string                                                           `json:"applicableFeatures,omitempty"`
	IsBackingResource        *bool                                                               `json:"isBackingResource,omitempty"`
	NetworkSecurityPerimeter *NetworkSecurityPerimeter                                           `json:"networkSecurityPerimeter,omitempty"`
	ParentAssociationName    *string                                                             `json:"parentAssociationName,omitempty"`
	Profile                  *NetworkSecurityPerimeterConfigurationPropertiesProfile             `json:"profile,omitempty"`
	ProvisioningIssues       *[]ProvisioningIssue                                                `json:"provisioningIssues,omitempty"`
	ProvisioningState        *NetworkSecurityPerimeterConfigurationProvisioningState             `json:"provisioningState,omitempty"`
	ResourceAssociation      *NetworkSecurityPerimeterConfigurationPropertiesResourceAssociation `json:"resourceAssociation,omitempty"`
	SourceResourceId         *string                                                             `json:"sourceResourceId,omitempty"`
}
