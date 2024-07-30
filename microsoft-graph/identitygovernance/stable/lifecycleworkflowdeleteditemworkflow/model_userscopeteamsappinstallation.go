package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserScopeTeamsAppInstallation struct {
	Chat                   *Chat                  `json:"chat,omitempty"`
	ConsentedPermissionSet *TeamsAppPermissionSet `json:"consentedPermissionSet,omitempty"`
	Id                     *string                `json:"id,omitempty"`
	ODataType              *string                `json:"@odata.type,omitempty"`
	TeamsApp               *TeamsApp              `json:"teamsApp,omitempty"`
	TeamsAppDefinition     *TeamsAppDefinition    `json:"teamsAppDefinition,omitempty"`
}
