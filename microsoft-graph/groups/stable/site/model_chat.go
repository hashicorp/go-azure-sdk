package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Chat struct {
	ChatType            *ChatChatType                      `json:"chatType,omitempty"`
	CreatedDateTime     *string                            `json:"createdDateTime,omitempty"`
	Id                  *string                            `json:"id,omitempty"`
	InstalledApps       *[]TeamsAppInstallation            `json:"installedApps,omitempty"`
	LastMessagePreview  *ChatMessageInfo                   `json:"lastMessagePreview,omitempty"`
	LastUpdatedDateTime *string                            `json:"lastUpdatedDateTime,omitempty"`
	Members             *[]ConversationMember              `json:"members,omitempty"`
	Messages            *[]ChatMessage                     `json:"messages,omitempty"`
	ODataType           *string                            `json:"@odata.type,omitempty"`
	OnlineMeetingInfo   *TeamworkOnlineMeetingInfo         `json:"onlineMeetingInfo,omitempty"`
	PermissionGrants    *[]ResourceSpecificPermissionGrant `json:"permissionGrants,omitempty"`
	PinnedMessages      *[]PinnedChatMessageInfo           `json:"pinnedMessages,omitempty"`
	Tabs                *[]TeamsTab                        `json:"tabs,omitempty"`
	TenantId            *string                            `json:"tenantId,omitempty"`
	Topic               *string                            `json:"topic,omitempty"`
	Viewpoint           *ChatViewpoint                     `json:"viewpoint,omitempty"`
	WebUrl              *string                            `json:"webUrl,omitempty"`
}
