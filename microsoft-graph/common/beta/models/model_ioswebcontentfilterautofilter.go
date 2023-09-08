package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosWebContentFilterAutoFilter struct {
	AllowedUrls *[]string `json:"allowedUrls,omitempty"`
	BlockedUrls *[]string `json:"blockedUrls,omitempty"`
	ODataType   *string   `json:"@odata.type,omitempty"`
}
