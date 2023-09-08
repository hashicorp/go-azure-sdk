package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BaseEndUserNotification struct {
	DefaultLanguage     *string              `json:"defaultLanguage,omitempty"`
	EndUserNotification *EndUserNotification `json:"endUserNotification,omitempty"`
	ODataType           *string              `json:"@odata.type,omitempty"`
}
