package notificationchannels

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

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

func (o *NotificationChannelProperties) GetCreatedDateAsTime() (*time.Time, error) {
	if o.CreatedDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedDate, "2006-01-02T15:04:05Z07:00")
}
