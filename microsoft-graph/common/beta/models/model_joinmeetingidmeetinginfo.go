package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JoinMeetingIdMeetingInfo struct {
	AllowConversationWithoutHost *bool   `json:"allowConversationWithoutHost,omitempty"`
	JoinMeetingId                *string `json:"joinMeetingId,omitempty"`
	ODataType                    *string `json:"@odata.type,omitempty"`
	Passcode                     *string `json:"passcode,omitempty"`
}
