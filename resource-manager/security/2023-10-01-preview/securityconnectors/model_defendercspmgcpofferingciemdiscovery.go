package securityconnectors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefenderCspmGcpOfferingCiemDiscovery struct {
	AzureActiveDirectoryAppName *string `json:"azureActiveDirectoryAppName,omitempty"`
	ServiceAccountEmailAddress  *string `json:"serviceAccountEmailAddress,omitempty"`
	WorkloadIdentityProviderId  *string `json:"workloadIdentityProviderId,omitempty"`
}
