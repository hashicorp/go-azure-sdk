package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintDocumentUploadProperties struct {
	ContentType  *string `json:"contentType,omitempty"`
	DocumentName *string `json:"documentName,omitempty"`
	ODataType    *string `json:"@odata.type,omitempty"`
	Size         *int64  `json:"size,omitempty"`
}
