package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AadUserConversationMember struct {
	DisplayName                 *string   `json:"displayName,omitempty"`
	Email                       *string   `json:"email,omitempty"`
	Id                          *string   `json:"id,omitempty"`
	ODataType                   *string   `json:"@odata.type,omitempty"`
	Roles                       *[]string `json:"roles,omitempty"`
	TenantId                    *string   `json:"tenantId,omitempty"`
	User                        *User     `json:"user,omitempty"`
	UserId                      *string   `json:"userId,omitempty"`
	VisibleHistoryStartDateTime *string   `json:"visibleHistoryStartDateTime,omitempty"`
}
