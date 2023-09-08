package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAppDefinition struct {
	AllowedInstallationScopes *TeamsAppDefinitionAllowedInstallationScopes `json:"allowedInstallationScopes,omitempty"`
	Authorization             *TeamsAppAuthorization                       `json:"authorization,omitempty"`
	AzureADAppId              *string                                      `json:"azureADAppId,omitempty"`
	Bot                       *TeamworkBot                                 `json:"bot,omitempty"`
	ColorIcon                 *TeamsAppIcon                                `json:"colorIcon,omitempty"`
	CreatedBy                 *IdentitySet                                 `json:"createdBy,omitempty"`
	Description               *string                                      `json:"description,omitempty"`
	DisplayName               *string                                      `json:"displayName,omitempty"`
	Id                        *string                                      `json:"id,omitempty"`
	LastModifiedDateTime      *string                                      `json:"lastModifiedDateTime,omitempty"`
	ODataType                 *string                                      `json:"@odata.type,omitempty"`
	OutlineIcon               *TeamsAppIcon                                `json:"outlineIcon,omitempty"`
	PublishingState           *TeamsAppDefinitionPublishingState           `json:"publishingState,omitempty"`
	Shortdescription          *string                                      `json:"shortdescription,omitempty"`
	TeamsAppId                *string                                      `json:"teamsAppId,omitempty"`
	Version                   *string                                      `json:"version,omitempty"`
}
