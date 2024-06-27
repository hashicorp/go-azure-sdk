package invoice

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DocumentDownloadRequest struct {
	DocumentName *string `json:"documentName,omitempty"`
	InvoiceName  *string `json:"invoiceName,omitempty"`
}
