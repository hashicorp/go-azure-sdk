package comanageddevice

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SendComanagedDeviceCustomNotificationToCompanyPortalRequest struct {
	NotificationBody  *string `json:"notificationBody,omitempty"`
	NotificationTitle *string `json:"notificationTitle,omitempty"`
}
