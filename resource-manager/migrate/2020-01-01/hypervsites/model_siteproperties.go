package hypervsites

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SiteProperties struct {
	AgentDetails                    *SiteAgentProperties `json:"agentDetails,omitempty"`
	ApplianceName                   *string              `json:"applianceName,omitempty"`
	DiscoverySolutionId             *string              `json:"discoverySolutionId,omitempty"`
	ServiceEndpoint                 *string              `json:"serviceEndpoint,omitempty"`
	ServicePrincipalIdentityDetails *SiteSpnProperties   `json:"servicePrincipalIdentityDetails,omitempty"`
}
