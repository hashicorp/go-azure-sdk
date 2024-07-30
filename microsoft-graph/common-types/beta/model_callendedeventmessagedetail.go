package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallEndedEventMessageDetail struct {
	CallDuration     *string                                   `json:"callDuration,omitempty"`
	CallEventType    *CallEndedEventMessageDetailCallEventType `json:"callEventType,omitempty"`
	CallId           *string                                   `json:"callId,omitempty"`
	CallParticipants *[]CallParticipantInfo                    `json:"callParticipants,omitempty"`
	Initiator        *IdentitySet                              `json:"initiator,omitempty"`
	ODataType        *string                                   `json:"@odata.type,omitempty"`
}
