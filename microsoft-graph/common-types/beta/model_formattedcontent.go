package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FormattedContent struct {
	Content   *string                 `json:"content,omitempty"`
	Format    *FormattedContentFormat `json:"format,omitempty"`
	ODataType *string                 `json:"@odata.type,omitempty"`
}
