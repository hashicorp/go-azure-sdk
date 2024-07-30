package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ParticipantLeftNotification struct {
	Call          *Call   `json:"call,omitempty"`
	Id            *string `json:"id,omitempty"`
	ODataType     *string `json:"@odata.type,omitempty"`
	ParticipantId *string `json:"participantId,omitempty"`
}
