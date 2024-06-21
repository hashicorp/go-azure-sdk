package notificationchannels

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotificationChannelProperties struct {
	CreatedDate        *string  `json:"createdDate,omitempty"`
	Description        *string  `json:"description,omitempty"`
	EmailRecipient     *string  `json:"emailRecipient,omitempty"`
	Events             *[]Event `json:"events,omitempty"`
	NotificationLocale *string  `json:"notificationLocale,omitempty"`
	ProvisioningState  *string  `json:"provisioningState,omitempty"`
	UniqueIdentifier   *string  `json:"uniqueIdentifier,omitempty"`
	WebHookUrl         *string  `json:"webHookUrl,omitempty"`
}
