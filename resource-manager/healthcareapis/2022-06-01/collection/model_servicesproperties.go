package collection

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServicesProperties struct {
	AccessPolicies              *[]ServiceAccessPolicyEntry             `json:"accessPolicies,omitempty"`
	AcrConfiguration            *ServiceAcrConfigurationInfo            `json:"acrConfiguration"`
	AuthenticationConfiguration *ServiceAuthenticationConfigurationInfo `json:"authenticationConfiguration"`
	CorsConfiguration           *ServiceCorsConfigurationInfo           `json:"corsConfiguration"`
	CosmosDbConfiguration       *ServiceCosmosDbConfigurationInfo       `json:"cosmosDbConfiguration"`
	ExportConfiguration         *ServiceExportConfigurationInfo         `json:"exportConfiguration"`
	ImportConfiguration         *ServiceImportConfigurationInfo         `json:"importConfiguration"`
	PrivateEndpointConnections  *[]PrivateEndpointConnection            `json:"privateEndpointConnections,omitempty"`
	ProvisioningState           *ProvisioningState                      `json:"provisioningState,omitempty"`
	PublicNetworkAccess         *PublicNetworkAccess                    `json:"publicNetworkAccess,omitempty"`
}
