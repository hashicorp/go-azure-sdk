package invitation

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
