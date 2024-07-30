package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityFormattedContent struct {
	Content   *string                         `json:"content,omitempty"`
	Format    *SecurityFormattedContentFormat `json:"format,omitempty"`
	ODataType *string                         `json:"@odata.type,omitempty"`
}
