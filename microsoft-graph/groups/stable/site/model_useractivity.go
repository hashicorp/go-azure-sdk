package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserActivity struct {
	ActivationUrl        *string                `json:"activationUrl,omitempty"`
	ActivitySourceHost   *string                `json:"activitySourceHost,omitempty"`
	AppActivityId        *string                `json:"appActivityId,omitempty"`
	AppDisplayName       *string                `json:"appDisplayName,omitempty"`
	ContentInfo          *Json                  `json:"contentInfo,omitempty"`
	ContentUrl           *string                `json:"contentUrl,omitempty"`
	CreatedDateTime      *string                `json:"createdDateTime,omitempty"`
	ExpirationDateTime   *string                `json:"expirationDateTime,omitempty"`
	FallbackUrl          *string                `json:"fallbackUrl,omitempty"`
	HistoryItems         *[]ActivityHistoryItem `json:"historyItems,omitempty"`
	Id                   *string                `json:"id,omitempty"`
	LastModifiedDateTime *string                `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                `json:"@odata.type,omitempty"`
	Status               *UserActivityStatus    `json:"status,omitempty"`
	UserTimezone         *string                `json:"userTimezone,omitempty"`
	VisualElements       *VisualInfo            `json:"visualElements,omitempty"`
}
