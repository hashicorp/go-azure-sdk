package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PresenceStatusMessage struct {
	ExpiryDateTime    *DateTimeTimeZone `json:"expiryDateTime,omitempty"`
	Message           *ItemBody         `json:"message,omitempty"`
	ODataType         *string           `json:"@odata.type,omitempty"`
	PublishedDateTime *string           `json:"publishedDateTime,omitempty"`
}
