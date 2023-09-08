package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesAgent struct {
	AgentGroups              *[]OnPremisesAgentGroup                    `json:"agentGroups,omitempty"`
	ExternalIp               *string                                    `json:"externalIp,omitempty"`
	Id                       *string                                    `json:"id,omitempty"`
	MachineName              *string                                    `json:"machineName,omitempty"`
	ODataType                *string                                    `json:"@odata.type,omitempty"`
	Status                   *OnPremisesAgentStatus                     `json:"status,omitempty"`
	SupportedPublishingTypes *[]OnPremisesAgentSupportedPublishingTypes `json:"supportedPublishingTypes,omitempty"`
}
