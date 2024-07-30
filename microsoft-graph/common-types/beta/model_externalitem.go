package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalItem struct {
	Acl        *[]Acl               `json:"acl,omitempty"`
	Content    *ExternalItemContent `json:"content,omitempty"`
	Id         *string              `json:"id,omitempty"`
	ODataType  *string              `json:"@odata.type,omitempty"`
	Properties *Properties          `json:"properties,omitempty"`
}
