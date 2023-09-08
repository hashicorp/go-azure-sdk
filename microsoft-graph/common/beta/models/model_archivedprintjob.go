package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ArchivedPrintJob struct {
	AcquiredByPrinter      *bool                            `json:"acquiredByPrinter,omitempty"`
	AcquiredDateTime       *string                          `json:"acquiredDateTime,omitempty"`
	BlackAndWhitePageCount *int64                           `json:"blackAndWhitePageCount,omitempty"`
	ColorPageCount         *int64                           `json:"colorPageCount,omitempty"`
	CompletionDateTime     *string                          `json:"completionDateTime,omitempty"`
	CopiesPrinted          *int64                           `json:"copiesPrinted,omitempty"`
	CreatedBy              *UserIdentity                    `json:"createdBy,omitempty"`
	CreatedDateTime        *string                          `json:"createdDateTime,omitempty"`
	DuplexPageCount        *int64                           `json:"duplexPageCount,omitempty"`
	Id                     *string                          `json:"id,omitempty"`
	ODataType              *string                          `json:"@odata.type,omitempty"`
	PageCount              *int64                           `json:"pageCount,omitempty"`
	PrinterId              *string                          `json:"printerId,omitempty"`
	PrinterName            *string                          `json:"printerName,omitempty"`
	ProcessingState        *ArchivedPrintJobProcessingState `json:"processingState,omitempty"`
	SimplexPageCount       *int64                           `json:"simplexPageCount,omitempty"`
}
