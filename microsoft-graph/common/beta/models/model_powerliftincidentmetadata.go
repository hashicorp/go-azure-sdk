package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PowerliftIncidentMetadata struct {
	Application       *string   `json:"application,omitempty"`
	ClientVersion     *string   `json:"clientVersion,omitempty"`
	CreatedAtDateTime *string   `json:"createdAtDateTime,omitempty"`
	EasyId            *string   `json:"easyId,omitempty"`
	FileNames         *[]string `json:"fileNames,omitempty"`
	Locale            *string   `json:"locale,omitempty"`
	ODataType         *string   `json:"@odata.type,omitempty"`
	Platform          *string   `json:"platform,omitempty"`
	PowerliftId       *string   `json:"powerliftId,omitempty"`
}
