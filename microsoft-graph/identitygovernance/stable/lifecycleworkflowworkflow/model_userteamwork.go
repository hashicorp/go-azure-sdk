package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserTeamwork struct {
	AssociatedTeams *[]AssociatedTeamInfo            `json:"associatedTeams,omitempty"`
	Id              *string                          `json:"id,omitempty"`
	InstalledApps   *[]UserScopeTeamsAppInstallation `json:"installedApps,omitempty"`
	ODataType       *string                          `json:"@odata.type,omitempty"`
}
