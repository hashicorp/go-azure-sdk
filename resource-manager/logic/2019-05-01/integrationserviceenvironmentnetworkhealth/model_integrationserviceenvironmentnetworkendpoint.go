package integrationserviceenvironmentnetworkhealth

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationServiceEnvironmentNetworkEndpoint struct {
	Accessibility *IntegrationServiceEnvironmentNetworkEndPointAccessibilityState `json:"accessibility,omitempty"`
	DomainName    *string                                                         `json:"domainName,omitempty"`
	Ports         *[]string                                                       `json:"ports,omitempty"`
}
