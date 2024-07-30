package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InviteParticipantsOperation struct {
	ClientContext *string                            `json:"clientContext,omitempty"`
	Id            *string                            `json:"id,omitempty"`
	ODataType     *string                            `json:"@odata.type,omitempty"`
	Participants  *[]InvitationParticipantInfo       `json:"participants,omitempty"`
	ResultInfo    *ResultInfo                        `json:"resultInfo,omitempty"`
	Status        *InviteParticipantsOperationStatus `json:"status,omitempty"`
}
