package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAppDefinition struct {
	Authorization        *TeamsAppAuthorization             `json:"authorization,omitempty"`
	Bot                  *TeamworkBot                       `json:"bot,omitempty"`
	CreatedBy            *IdentitySet                       `json:"createdBy,omitempty"`
	Description          *string                            `json:"description,omitempty"`
	DisplayName          *string                            `json:"displayName,omitempty"`
	Id                   *string                            `json:"id,omitempty"`
	LastModifiedDateTime *string                            `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                            `json:"@odata.type,omitempty"`
	PublishingState      *TeamsAppDefinitionPublishingState `json:"publishingState,omitempty"`
	ShortDescription     *string                            `json:"shortDescription,omitempty"`
	TeamsAppId           *string                            `json:"teamsAppId,omitempty"`
	Version              *string                            `json:"version,omitempty"`
}
