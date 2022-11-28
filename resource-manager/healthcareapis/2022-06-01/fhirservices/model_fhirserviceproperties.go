package fhirservices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FhirServiceProperties struct {
	AccessPolicies                     *[]FhirServiceAccessPolicyEntry         `json:"accessPolicies,omitempty"`
	AcrConfiguration                   *FhirServiceAcrConfiguration            `json:"acrConfiguration"`
	AuthenticationConfiguration        *FhirServiceAuthenticationConfiguration `json:"authenticationConfiguration"`
	CorsConfiguration                  *FhirServiceCorsConfiguration           `json:"corsConfiguration"`
	EventState                         *ServiceEventState                      `json:"eventState,omitempty"`
	ExportConfiguration                *FhirServiceExportConfiguration         `json:"exportConfiguration"`
	ImportConfiguration                *FhirServiceImportConfiguration         `json:"importConfiguration"`
	PrivateEndpointConnections         *[]PrivateEndpointConnection            `json:"privateEndpointConnections,omitempty"`
	ProvisioningState                  *ProvisioningState                      `json:"provisioningState,omitempty"`
	PublicNetworkAccess                *PublicNetworkAccess                    `json:"publicNetworkAccess,omitempty"`
	ResourceVersionPolicyConfiguration *ResourceVersionPolicyConfiguration     `json:"resourceVersionPolicyConfiguration"`
}
