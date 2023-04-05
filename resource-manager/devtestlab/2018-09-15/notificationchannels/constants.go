package notificationchannels

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotificationChannelEventType string

const (
	NotificationChannelEventTypeAutoShutdown NotificationChannelEventType = "AutoShutdown"
	NotificationChannelEventTypeCost         NotificationChannelEventType = "Cost"
)

func PossibleValuesForNotificationChannelEventType() []string {
	return []string{
		string(NotificationChannelEventTypeAutoShutdown),
		string(NotificationChannelEventTypeCost),
	}
}
