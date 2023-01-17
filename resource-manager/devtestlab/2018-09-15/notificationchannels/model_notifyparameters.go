package notificationchannels

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotifyParameters struct {
	EventName   *NotificationChannelEventType `json:"eventName,omitempty"`
	JsonPayload *string                       `json:"jsonPayload,omitempty"`
}
