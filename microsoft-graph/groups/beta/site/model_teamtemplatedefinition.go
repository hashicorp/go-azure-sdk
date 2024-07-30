package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamTemplateDefinition struct {
	Audience             *TeamTemplateDefinitionAudience `json:"audience,omitempty"`
	Categories           *[]string                       `json:"categories,omitempty"`
	Description          *string                         `json:"description,omitempty"`
	DisplayName          *string                         `json:"displayName,omitempty"`
	IconUrl              *string                         `json:"iconUrl,omitempty"`
	Id                   *string                         `json:"id,omitempty"`
	LanguageTag          *string                         `json:"languageTag,omitempty"`
	LastModifiedBy       *IdentitySet                    `json:"lastModifiedBy,omitempty"`
	LastModifiedDateTime *string                         `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                         `json:"@odata.type,omitempty"`
	ParentTemplateId     *string                         `json:"parentTemplateId,omitempty"`
	PublisherName        *string                         `json:"publisherName,omitempty"`
	ShortDescription     *string                         `json:"shortDescription,omitempty"`
	TeamDefinition       *Team                           `json:"teamDefinition,omitempty"`
}
