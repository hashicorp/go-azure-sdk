package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharingLink struct {
	Application      *Identity `json:"application,omitempty"`
	ConfiguratorUrl  *string   `json:"configuratorUrl,omitempty"`
	ODataType        *string   `json:"@odata.type,omitempty"`
	PreventsDownload *bool     `json:"preventsDownload,omitempty"`
	Scope            *string   `json:"scope,omitempty"`
	Type             *string   `json:"type,omitempty"`
	WebHtml          *string   `json:"webHtml,omitempty"`
	WebUrl           *string   `json:"webUrl,omitempty"`
}
