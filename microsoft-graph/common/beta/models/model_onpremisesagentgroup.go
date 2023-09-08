package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesAgentGroup struct {
	Agents             *[]OnPremisesAgent                  `json:"agents,omitempty"`
	DisplayName        *string                             `json:"displayName,omitempty"`
	Id                 *string                             `json:"id,omitempty"`
	IsDefault          *bool                               `json:"isDefault,omitempty"`
	ODataType          *string                             `json:"@odata.type,omitempty"`
	PublishedResources *[]PublishedResource                `json:"publishedResources,omitempty"`
	PublishingType     *OnPremisesAgentGroupPublishingType `json:"publishingType,omitempty"`
}
