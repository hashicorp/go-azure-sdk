package fhirservices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FhirServiceProperties struct {
	AccessPolicies              *[]FhirServiceAccessPolicyEntry         `json:"accessPolicies,omitempty"`
	AcrConfiguration            *FhirServiceAcrConfiguration            `json:"acrConfiguration,omitempty"`
	AuthenticationConfiguration *FhirServiceAuthenticationConfiguration `json:"authenticationConfiguration,omitempty"`
	CorsConfiguration           *FhirServiceCorsConfiguration           `json:"corsConfiguration,omitempty"`
	ExportConfiguration         *FhirServiceExportConfiguration         `json:"exportConfiguration,omitempty"`
	ProvisioningState           *ProvisioningState                      `json:"provisioningState,omitempty"`
}
