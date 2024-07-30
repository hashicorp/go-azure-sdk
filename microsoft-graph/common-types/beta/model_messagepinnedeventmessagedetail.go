package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessagePinnedEventMessageDetail struct {
	EventDateTime *string      `json:"eventDateTime,omitempty"`
	Initiator     *IdentitySet `json:"initiator,omitempty"`
	ODataType     *string      `json:"@odata.type,omitempty"`
}
