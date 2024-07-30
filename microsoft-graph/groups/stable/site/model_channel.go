package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Channel struct {
	CreatedDateTime     *string                      `json:"createdDateTime,omitempty"`
	Description         *string                      `json:"description,omitempty"`
	DisplayName         *string                      `json:"displayName,omitempty"`
	Email               *string                      `json:"email,omitempty"`
	FilesFolder         *DriveItem                   `json:"filesFolder,omitempty"`
	Id                  *string                      `json:"id,omitempty"`
	IsFavoriteByDefault *bool                        `json:"isFavoriteByDefault,omitempty"`
	Members             *[]ConversationMember        `json:"members,omitempty"`
	MembershipType      *ChannelMembershipType       `json:"membershipType,omitempty"`
	Messages            *[]ChatMessage               `json:"messages,omitempty"`
	ODataType           *string                      `json:"@odata.type,omitempty"`
	SharedWithTeams     *[]SharedWithChannelTeamInfo `json:"sharedWithTeams,omitempty"`
	Summary             *ChannelSummary              `json:"summary,omitempty"`
	Tabs                *[]TeamsTab                  `json:"tabs,omitempty"`
	TenantId            *string                      `json:"tenantId,omitempty"`
	WebUrl              *string                      `json:"webUrl,omitempty"`
}
