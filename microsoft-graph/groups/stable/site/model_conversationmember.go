package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConversationMember struct {
	DisplayName                 *string   `json:"displayName,omitempty"`
	Id                          *string   `json:"id,omitempty"`
	ODataType                   *string   `json:"@odata.type,omitempty"`
	Roles                       *[]string `json:"roles,omitempty"`
	VisibleHistoryStartDateTime *string   `json:"visibleHistoryStartDateTime,omitempty"`
}
