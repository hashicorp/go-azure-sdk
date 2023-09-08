package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ArchivedPrintJobOperationPredicate struct {
	AcquiredByPrinter  *bool
	AcquiredDateTime   *string
	CompletionDateTime *string
	CopiesPrinted      *int64
	CreatedDateTime    *string
	Id                 *string
	ODataType          *string
	PrinterId          *string
	PrinterName        *string
}

func (p ArchivedPrintJobOperationPredicate) Matches(input ArchivedPrintJob) bool {

	if p.AcquiredByPrinter != nil && (input.AcquiredByPrinter == nil || *p.AcquiredByPrinter != *input.AcquiredByPrinter) {
		return false
	}

	if p.AcquiredDateTime != nil && (input.AcquiredDateTime == nil || *p.AcquiredDateTime != *input.AcquiredDateTime) {
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

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PrinterId != nil && (input.PrinterId == nil || *p.PrinterId != *input.PrinterId) {
		return false
	}

	if p.PrinterName != nil && (input.PrinterName == nil || *p.PrinterName != *input.PrinterName) {
		return false
	}

	return true
}
