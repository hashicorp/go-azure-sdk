package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Teamwork struct {
	DeletedTeams          *[]DeletedTeam          `json:"deletedTeams,omitempty"`
	Id                    *string                 `json:"id,omitempty"`
	ODataType             *string                 `json:"@odata.type,omitempty"`
	TeamsAppSettings      *TeamsAppSettings       `json:"teamsAppSettings,omitempty"`
	WorkforceIntegrations *[]WorkforceIntegration `json:"workforceIntegrations,omitempty"`
}
