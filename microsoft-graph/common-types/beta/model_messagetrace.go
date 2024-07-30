package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessageTrace struct {
	DestinationIPAddress *string             `json:"destinationIPAddress,omitempty"`
	Id                   *string             `json:"id,omitempty"`
	MessageId            *string             `json:"messageId,omitempty"`
	ODataType            *string             `json:"@odata.type,omitempty"`
	ReceivedDateTime     *string             `json:"receivedDateTime,omitempty"`
	Recipients           *[]MessageRecipient `json:"recipients,omitempty"`
	SenderEmail          *string             `json:"senderEmail,omitempty"`
	Size                 *int64              `json:"size,omitempty"`
	SourceIPAddress      *string             `json:"sourceIPAddress,omitempty"`
	Subject              *string             `json:"subject,omitempty"`
}
