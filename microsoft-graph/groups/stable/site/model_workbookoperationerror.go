package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookOperationError struct {
	Code       *string                 `json:"code,omitempty"`
	InnerError *WorkbookOperationError `json:"innerError,omitempty"`
	Message    *string                 `json:"message,omitempty"`
	ODataType  *string                 `json:"@odata.type,omitempty"`
}
