package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewNotificationRecipientItem struct {
	NotificationRecipientScope *AccessReviewNotificationRecipientScope `json:"notificationRecipientScope,omitempty"`
	NotificationTemplateType   *string                                 `json:"notificationTemplateType,omitempty"`
	ODataType                  *string                                 `json:"@odata.type,omitempty"`
}
