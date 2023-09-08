package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosAvailableUpdateVersion struct {
	ExpirationDateTime *string   `json:"expirationDateTime,omitempty"`
	ODataType          *string   `json:"@odata.type,omitempty"`
	PostingDateTime    *string   `json:"postingDateTime,omitempty"`
	ProductVersion     *string   `json:"productVersion,omitempty"`
	SupportedDevices   *[]string `json:"supportedDevices,omitempty"`
}
