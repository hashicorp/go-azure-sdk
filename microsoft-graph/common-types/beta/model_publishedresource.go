package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PublishedResource struct {
	AgentGroups    *[]OnPremisesAgentGroup          `json:"agentGroups,omitempty"`
	DisplayName    *string                          `json:"displayName,omitempty"`
	Id             *string                          `json:"id,omitempty"`
	ODataType      *string                          `json:"@odata.type,omitempty"`
	PublishingType *PublishedResourcePublishingType `json:"publishingType,omitempty"`
	ResourceName   *string                          `json:"resourceName,omitempty"`
}
