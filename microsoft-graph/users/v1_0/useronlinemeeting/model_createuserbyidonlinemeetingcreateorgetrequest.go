package useronlinemeeting

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserByIdOnlineMeetingCreateOrGetRequest struct {
	ChatInfo      *ChatInfo            `json:"chatInfo,omitempty"`
	EndDateTime   *string              `json:"endDateTime,omitempty"`
	ExternalId    *string              `json:"externalId,omitempty"`
	Participants  *MeetingParticipants `json:"participants,omitempty"`
	StartDateTime *string              `json:"startDateTime,omitempty"`
	Subject       *string              `json:"subject,omitempty"`
}
