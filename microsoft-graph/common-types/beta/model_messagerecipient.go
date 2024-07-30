package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessageRecipient struct {
	DeliveryStatus *MessageRecipientDeliveryStatus `json:"deliveryStatus,omitempty"`
	Events         *[]MessageEvent                 `json:"events,omitempty"`
	Id             *string                         `json:"id,omitempty"`
	ODataType      *string                         `json:"@odata.type,omitempty"`
	RecipientEmail *string                         `json:"recipientEmail,omitempty"`
}
