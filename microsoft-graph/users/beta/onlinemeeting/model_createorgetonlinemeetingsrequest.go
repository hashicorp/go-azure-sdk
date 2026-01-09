package onlinemeeting

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOrGetOnlineMeetingsRequest struct {
	ChatInfo      *beta.ChatInfo            `json:"chatInfo,omitempty"`
	EndDateTime   nullable.Type[string]     `json:"endDateTime,omitempty"`
	ExternalId    *string                   `json:"externalId,omitempty"`
	Participants  *beta.MeetingParticipants `json:"participants,omitempty"`
	StartDateTime nullable.Type[string]     `json:"startDateTime,omitempty"`
	Subject       nullable.Type[string]     `json:"subject,omitempty"`
}
