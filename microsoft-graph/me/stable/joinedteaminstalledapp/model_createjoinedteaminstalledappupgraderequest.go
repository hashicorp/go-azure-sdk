package joinedteaminstalledapp

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateJoinedTeamInstalledAppUpgradeRequest struct {
	ConsentedPermissionSet *TeamsAppPermissionSet `json:"consentedPermissionSet,omitempty"`
}
