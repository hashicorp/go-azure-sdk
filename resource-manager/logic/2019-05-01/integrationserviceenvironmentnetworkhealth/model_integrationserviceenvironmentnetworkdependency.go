package integrationserviceenvironmentnetworkhealth

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationServiceEnvironmentNetworkDependency struct {
	Category    *IntegrationServiceEnvironmentNetworkDependencyCategoryType `json:"category,omitempty"`
	DisplayName *string                                                     `json:"displayName,omitempty"`
	Endpoints   *[]IntegrationServiceEnvironmentNetworkEndpoint             `json:"endpoints,omitempty"`
}
