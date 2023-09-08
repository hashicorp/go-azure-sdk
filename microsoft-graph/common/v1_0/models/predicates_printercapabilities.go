package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterCapabilitiesOperationPredicate struct {
	Collation                *bool
	IsColorPrintingSupported *bool
	IsPageRangeSupported     *bool
	ODataType                *string
	SupportsFitPdfToPage     *bool
}

func (p PrinterCapabilitiesOperationPredicate) Matches(input PrinterCapabilities) bool {

	if p.Collation != nil && (input.Collation == nil || *p.Collation != *input.Collation) {
		return false
	}

	if p.IsColorPrintingSupported != nil && (input.IsColorPrintingSupported == nil || *p.IsColorPrintingSupported != *input.IsColorPrintingSupported) {
		return false
	}

	if p.IsPageRangeSupported != nil && (input.IsPageRangeSupported == nil || *p.IsPageRangeSupported != *input.IsPageRangeSupported) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SupportsFitPdfToPage != nil && (input.SupportsFitPdfToPage == nil || *p.SupportsFitPdfToPage != *input.SupportsFitPdfToPage) {
		return false
	}

	return true
}
