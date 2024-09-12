package onlinemeeting

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOrGetOnlineMeetingRequest struct {
	ChatInfo      *stable.ChatInfo            `json:"chatInfo,omitempty"`
	EndDateTime   nullable.Type[string]       `json:"endDateTime,omitempty"`
	ExternalId    *string                     `json:"externalId,omitempty"`
	Participants  *stable.MeetingParticipants `json:"participants,omitempty"`
	StartDateTime nullable.Type[string]       `json:"startDateTime,omitempty"`
	Subject       nullable.Type[string]       `json:"subject,omitempty"`
}
