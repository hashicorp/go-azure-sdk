package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ArchivedPrintJobOperationPredicate struct {
	AcquiredByPrinter      *bool
	AcquiredDateTime       *string
	BlackAndWhitePageCount *int64
	ColorPageCount         *int64
	CompletionDateTime     *string
	CopiesPrinted          *int64
	CreatedDateTime        *string
	DuplexPageCount        *int64
	Id                     *string
	ODataType              *string
	PageCount              *int64
	PrinterId              *string
	PrinterName            *string
	SimplexPageCount       *int64
}

func (p ArchivedPrintJobOperationPredicate) Matches(input ArchivedPrintJob) bool {

	if p.AcquiredByPrinter != nil && (input.AcquiredByPrinter == nil || *p.AcquiredByPrinter != *input.AcquiredByPrinter) {
		return false
	}

	if p.AcquiredDateTime != nil && (input.AcquiredDateTime == nil || *p.AcquiredDateTime != *input.AcquiredDateTime) {
		return false
	}

	if p.BlackAndWhitePageCount != nil && (input.BlackAndWhitePageCount == nil || *p.BlackAndWhitePageCount != *input.BlackAndWhitePageCount) {
		return false
	}

	if p.ColorPageCount != nil && (input.ColorPageCount == nil || *p.ColorPageCount != *input.ColorPageCount) {
		return false
	}

	if p.CompletionDateTime != nil && (input.CompletionDateTime == nil || *p.CompletionDateTime != *input.CompletionDateTime) {
		return false
	}

	if p.CopiesPrinted != nil && (input.CopiesPrinted == nil || *p.CopiesPrinted != *input.CopiesPrinted) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DuplexPageCount != nil && (input.DuplexPageCount == nil || *p.DuplexPageCount != *input.DuplexPageCount) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PageCount != nil && (input.PageCount == nil || *p.PageCount != *input.PageCount) {
		return false
	}

	if p.PrinterId != nil && (input.PrinterId == nil || *p.PrinterId != *input.PrinterId) {
		return false
	}

	if p.PrinterName != nil && (input.PrinterName == nil || *p.PrinterName != *input.PrinterName) {
		return false
	}

	if p.SimplexPageCount != nil && (input.SimplexPageCount == nil || *p.SimplexPageCount != *input.SimplexPageCount) {
		return false
	}

	return true
}
