package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndUserNotificationDetail struct {
	EmailContent      *string        `json:"emailContent,omitempty"`
	Id                *string        `json:"id,omitempty"`
	IsDefaultLangauge *bool          `json:"isDefaultLangauge,omitempty"`
	Language          *string        `json:"language,omitempty"`
	Locale            *string        `json:"locale,omitempty"`
	ODataType         *string        `json:"@odata.type,omitempty"`
	SentFrom          *EmailIdentity `json:"sentFrom,omitempty"`
	Subject           *string        `json:"subject,omitempty"`
}
