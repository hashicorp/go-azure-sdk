package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingPolicyUpdatedEventMessageDetail struct {
	Initiator          *IdentitySet `json:"initiator,omitempty"`
	MeetingChatEnabled *bool        `json:"meetingChatEnabled,omitempty"`
	MeetingChatId      *string      `json:"meetingChatId,omitempty"`
	ODataType          *string      `json:"@odata.type,omitempty"`
}
