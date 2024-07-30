package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImageInfo struct {
	AddImageQuery   *bool   `json:"addImageQuery,omitempty"`
	AlternateText   *string `json:"alternateText,omitempty"`
	AlternativeText *string `json:"alternativeText,omitempty"`
	IconUrl         *string `json:"iconUrl,omitempty"`
	ODataType       *string `json:"@odata.type,omitempty"`
}
