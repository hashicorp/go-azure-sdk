package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessageInfo struct {
	Body            *ItemBody                   `json:"body,omitempty"`
	CreatedDateTime *string                     `json:"createdDateTime,omitempty"`
	EventDetail     *EventMessageDetail         `json:"eventDetail,omitempty"`
	From            *ChatMessageFromIdentitySet `json:"from,omitempty"`
	Id              *string                     `json:"id,omitempty"`
	IsDeleted       *bool                       `json:"isDeleted,omitempty"`
	MessageType     *ChatMessageInfoMessageType `json:"messageType,omitempty"`
	ODataType       *string                     `json:"@odata.type,omitempty"`
}
