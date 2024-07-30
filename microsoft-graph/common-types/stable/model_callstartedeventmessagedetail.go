package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallStartedEventMessageDetail struct {
	CallEventType *CallStartedEventMessageDetailCallEventType `json:"callEventType,omitempty"`
	CallId        *string                                     `json:"callId,omitempty"`
	Initiator     *IdentitySet                                `json:"initiator,omitempty"`
	ODataType     *string                                     `json:"@odata.type,omitempty"`
}
