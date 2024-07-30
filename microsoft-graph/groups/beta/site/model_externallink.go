package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalLink struct {
	Href      *string `json:"href,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
}
