package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Notification struct {
	DisplayTimeToLive  *int64                 `json:"displayTimeToLive,omitempty"`
	ExpirationDateTime *string                `json:"expirationDateTime,omitempty"`
	GroupName          *string                `json:"groupName,omitempty"`
	Id                 *string                `json:"id,omitempty"`
	ODataType          *string                `json:"@odata.type,omitempty"`
	Payload            *PayloadTypes          `json:"payload,omitempty"`
	Priority           *NotificationPriority  `json:"priority,omitempty"`
	TargetHostName     *string                `json:"targetHostName,omitempty"`
	TargetPolicy       *TargetPolicyEndpoints `json:"targetPolicy,omitempty"`
}
