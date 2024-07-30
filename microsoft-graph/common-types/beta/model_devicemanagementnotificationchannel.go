package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementNotificationChannel struct {
	NotificationChannelType *DeviceManagementNotificationChannelNotificationChannelType `json:"notificationChannelType,omitempty"`
	NotificationReceivers   *[]DeviceManagementNotificationReceiver                     `json:"notificationReceivers,omitempty"`
	ODataType               *string                                                     `json:"@odata.type,omitempty"`
}
