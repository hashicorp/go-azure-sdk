package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesPublishingProfile struct {
	AgentGroups                     *[]OnPremisesAgentGroup          `json:"agentGroups,omitempty"`
	Agents                          *[]OnPremisesAgent               `json:"agents,omitempty"`
	ConnectorGroups                 *[]ConnectorGroup                `json:"connectorGroups,omitempty"`
	Connectors                      *[]Connector                     `json:"connectors,omitempty"`
	HybridAgentUpdaterConfiguration *HybridAgentUpdaterConfiguration `json:"hybridAgentUpdaterConfiguration,omitempty"`
	Id                              *string                          `json:"id,omitempty"`
	IsDefaultAccessEnabled          *bool                            `json:"isDefaultAccessEnabled,omitempty"`
	IsEnabled                       *bool                            `json:"isEnabled,omitempty"`
	ODataType                       *string                          `json:"@odata.type,omitempty"`
	PublishedResources              *[]PublishedResource             `json:"publishedResources,omitempty"`
}
