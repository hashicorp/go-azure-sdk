package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttachmentSession struct {
	Content            *string   `json:"content,omitempty"`
	ExpirationDateTime *string   `json:"expirationDateTime,omitempty"`
	Id                 *string   `json:"id,omitempty"`
	NextExpectedRanges *[]string `json:"nextExpectedRanges,omitempty"`
	ODataType          *string   `json:"@odata.type,omitempty"`
}
