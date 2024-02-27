package invitation

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvitationProperties struct {
	InvitationId            *string           `json:"invitationId,omitempty"`
	InvitationStatus        *InvitationStatus `json:"invitationStatus,omitempty"`
	RespondedAt             *string           `json:"respondedAt,omitempty"`
	SentAt                  *string           `json:"sentAt,omitempty"`
	TargetActiveDirectoryId *string           `json:"targetActiveDirectoryId,omitempty"`
	TargetEmail             *string           `json:"targetEmail,omitempty"`
	TargetObjectId          *string           `json:"targetObjectId,omitempty"`
	UserEmail               *string           `json:"userEmail,omitempty"`
	UserName                *string           `json:"userName,omitempty"`
}

func (o *InvitationProperties) GetRespondedAtAsTime() (*time.Time, error) {
	if o.RespondedAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.RespondedAt, "2006-01-02T15:04:05Z07:00")
}

func (o *InvitationProperties) GetSentAtAsTime() (*time.Time, error) {
	if o.SentAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.SentAt, "2006-01-02T15:04:05Z07:00")
}
