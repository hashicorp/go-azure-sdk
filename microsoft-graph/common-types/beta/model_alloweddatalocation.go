package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AllowedDataLocation struct {
	AppId     *string `json:"appId,omitempty"`
	Domain    *string `json:"domain,omitempty"`
	Id        *string `json:"id,omitempty"`
	IsDefault *bool   `json:"isDefault,omitempty"`
	Location  *string `json:"location,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
}
