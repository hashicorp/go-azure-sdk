package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedWithChannelTeamInfo struct {
	AllowedMembers *[]ConversationMember `json:"allowedMembers,omitempty"`
	DisplayName    *string               `json:"displayName,omitempty"`
	Id             *string               `json:"id,omitempty"`
	IsHostTeam     *bool                 `json:"isHostTeam,omitempty"`
	ODataType      *string               `json:"@odata.type,omitempty"`
	Team           *Team                 `json:"team,omitempty"`
	TenantId       *string               `json:"tenantId,omitempty"`
}
