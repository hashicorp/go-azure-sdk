package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocalizedNotificationMessage struct {
	Id                   *string `json:"id,omitempty"`
	IsDefault            *bool   `json:"isDefault,omitempty"`
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`
	Locale               *string `json:"locale,omitempty"`
	MessageTemplate      *string `json:"messageTemplate,omitempty"`
	ODataType            *string `json:"@odata.type,omitempty"`
	Subject              *string `json:"subject,omitempty"`
}
