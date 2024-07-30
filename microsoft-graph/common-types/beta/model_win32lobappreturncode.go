package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppReturnCode struct {
	ODataType  *string                    `json:"@odata.type,omitempty"`
	ReturnCode *int64                     `json:"returnCode,omitempty"`
	Type       *Win32LobAppReturnCodeType `json:"type,omitempty"`
}
