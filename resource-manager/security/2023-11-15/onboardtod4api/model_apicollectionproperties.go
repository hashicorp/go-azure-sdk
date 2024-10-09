package onboardtod4api

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiCollectionProperties struct {
	BaseURL                                      *string            `json:"baseUrl,omitempty"`
	DiscoveredVia                                *string            `json:"discoveredVia,omitempty"`
	DisplayName                                  *string            `json:"displayName,omitempty"`
	NumberOfApiEndpoints                         *int64             `json:"numberOfApiEndpoints,omitempty"`
	NumberOfApiEndpointsWithSensitiveDataExposed *int64             `json:"numberOfApiEndpointsWithSensitiveDataExposed,omitempty"`
	NumberOfExternalApiEndpoints                 *int64             `json:"numberOfExternalApiEndpoints,omitempty"`
	NumberOfInactiveApiEndpoints                 *int64             `json:"numberOfInactiveApiEndpoints,omitempty"`
	NumberOfUnauthenticatedApiEndpoints          *int64             `json:"numberOfUnauthenticatedApiEndpoints,omitempty"`
	ProvisioningState                            *ProvisioningState `json:"provisioningState,omitempty"`
	SensitivityLabel                             *string            `json:"sensitivityLabel,omitempty"`
}
