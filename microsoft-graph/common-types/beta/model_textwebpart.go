package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TextWebPart struct {
	Id        *string `json:"id,omitempty"`
	InnerHtml *string `json:"innerHtml,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
}
