package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintDocument struct {
	Configuration      *PrinterDocumentConfiguration `json:"configuration,omitempty"`
	ContentType        *string                       `json:"contentType,omitempty"`
	DisplayName        *string                       `json:"displayName,omitempty"`
	DownloadedDateTime *string                       `json:"downloadedDateTime,omitempty"`
	Id                 *string                       `json:"id,omitempty"`
	ODataType          *string                       `json:"@odata.type,omitempty"`
	Size               *int64                        `json:"size,omitempty"`
	UploadedDateTime   *string                       `json:"uploadedDateTime,omitempty"`
}
