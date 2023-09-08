package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Teamwork struct {
	DeletedChats          *[]DeletedChat          `json:"deletedChats,omitempty"`
	DeletedTeams          *[]DeletedTeam          `json:"deletedTeams,omitempty"`
	Devices               *[]TeamworkDevice       `json:"devices,omitempty"`
	Id                    *string                 `json:"id,omitempty"`
	ODataType             *string                 `json:"@odata.type,omitempty"`
	TeamTemplates         *[]TeamTemplate         `json:"teamTemplates,omitempty"`
	TeamsAppSettings      *TeamsAppSettings       `json:"teamsAppSettings,omitempty"`
	WorkforceIntegrations *[]WorkforceIntegration `json:"workforceIntegrations,omitempty"`
}
