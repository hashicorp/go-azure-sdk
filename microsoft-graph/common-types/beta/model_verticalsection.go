package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VerticalSection struct {
	Emphasis  *VerticalSectionEmphasis `json:"emphasis,omitempty"`
	Id        *string                  `json:"id,omitempty"`
	ODataType *string                  `json:"@odata.type,omitempty"`
	Webparts  *[]WebPart               `json:"webparts,omitempty"`
}
