package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Team struct {
	AllChannels       *[]Channel                         `json:"allChannels,omitempty"`
	Channels          *[]Channel                         `json:"channels,omitempty"`
	Classification    *string                            `json:"classification,omitempty"`
	CreatedDateTime   *string                            `json:"createdDateTime,omitempty"`
	Description       *string                            `json:"description,omitempty"`
	DisplayName       *string                            `json:"displayName,omitempty"`
	FunSettings       *TeamFunSettings                   `json:"funSettings,omitempty"`
	Group             *Group                             `json:"group,omitempty"`
	GuestSettings     *TeamGuestSettings                 `json:"guestSettings,omitempty"`
	Id                *string                            `json:"id,omitempty"`
	IncomingChannels  *[]Channel                         `json:"incomingChannels,omitempty"`
	InstalledApps     *[]TeamsAppInstallation            `json:"installedApps,omitempty"`
	InternalId        *string                            `json:"internalId,omitempty"`
	IsArchived        *bool                              `json:"isArchived,omitempty"`
	MemberSettings    *TeamMemberSettings                `json:"memberSettings,omitempty"`
	Members           *[]ConversationMember              `json:"members,omitempty"`
	MessagingSettings *TeamMessagingSettings             `json:"messagingSettings,omitempty"`
	ODataType         *string                            `json:"@odata.type,omitempty"`
	Operations        *[]TeamsAsyncOperation             `json:"operations,omitempty"`
	PermissionGrants  *[]ResourceSpecificPermissionGrant `json:"permissionGrants,omitempty"`
	Photo             *ProfilePhoto                      `json:"photo,omitempty"`
	PrimaryChannel    *Channel                           `json:"primaryChannel,omitempty"`
	Schedule          *Schedule                          `json:"schedule,omitempty"`
	Specialization    *TeamSpecialization                `json:"specialization,omitempty"`
	Summary           *TeamSummary                       `json:"summary,omitempty"`
	Tags              *[]TeamworkTag                     `json:"tags,omitempty"`
	Template          *TeamsTemplate                     `json:"template,omitempty"`
	TenantId          *string                            `json:"tenantId,omitempty"`
	Visibility        *TeamVisibility                    `json:"visibility,omitempty"`
	WebUrl            *string                            `json:"webUrl,omitempty"`
}
