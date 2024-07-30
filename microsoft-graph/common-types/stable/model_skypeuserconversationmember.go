package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SkypeUserConversationMember struct {
	DisplayName                 *string   `json:"displayName,omitempty"`
	Id                          *string   `json:"id,omitempty"`
	ODataType                   *string   `json:"@odata.type,omitempty"`
	Roles                       *[]string `json:"roles,omitempty"`
	SkypeId                     *string   `json:"skypeId,omitempty"`
	VisibleHistoryStartDateTime *string   `json:"visibleHistoryStartDateTime,omitempty"`
}
