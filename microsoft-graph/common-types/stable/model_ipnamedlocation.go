package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IpNamedLocation struct {
	CreatedDateTime  *string    `json:"createdDateTime,omitempty"`
	DisplayName      *string    `json:"displayName,omitempty"`
	Id               *string    `json:"id,omitempty"`
	IpRanges         *[]IpRange `json:"ipRanges,omitempty"`
	IsTrusted        *bool      `json:"isTrusted,omitempty"`
	ModifiedDateTime *string    `json:"modifiedDateTime,omitempty"`
	ODataType        *string    `json:"@odata.type,omitempty"`
}
