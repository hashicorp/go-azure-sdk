package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessageEvent struct {
	DateTime    *string                `json:"dateTime,omitempty"`
	Description *string                `json:"description,omitempty"`
	EventType   *MessageEventEventType `json:"eventType,omitempty"`
	Id          *string                `json:"id,omitempty"`
	ODataType   *string                `json:"@odata.type,omitempty"`
}
