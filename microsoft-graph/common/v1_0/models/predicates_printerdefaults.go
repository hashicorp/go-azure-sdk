package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDefaultsOperationPredicate struct {
	ContentType   *string
	CopiesPerJob  *int64
	Dpi           *int64
	FitPdfToPage  *bool
	InputBin      *string
	MediaColor    *string
	MediaSize     *string
	MediaType     *string
	ODataType     *string
	OutputBin     *string
	PagesPerSheet *int64
}

func (p PrinterDefaultsOperationPredicate) Matches(input PrinterDefaults) bool {

	if p.ContentType != nil && (input.ContentType == nil || *p.ContentType != *input.ContentType) {
		return false
	}

	if p.CopiesPerJob != nil && (input.CopiesPerJob == nil || *p.CopiesPerJob != *input.CopiesPerJob) {
		return false
	}

	if p.Dpi != nil && (input.Dpi == nil || *p.Dpi != *input.Dpi) {
		return false
	}

	if p.FitPdfToPage != nil && (input.FitPdfToPage == nil || *p.FitPdfToPage != *input.FitPdfToPage) {
		return false
	}

	if p.InputBin != nil && (input.InputBin == nil || *p.InputBin != *input.InputBin) {
		return false
	}

	if p.MediaColor != nil && (input.MediaColor == nil || *p.MediaColor != *input.MediaColor) {
		return false
	}

	if p.MediaSize != nil && (input.MediaSize == nil || *p.MediaSize != *input.MediaSize) {
		return false
	}

	if p.MediaType != nil && (input.MediaType == nil || *p.MediaType != *input.MediaType) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OutputBin != nil && (input.OutputBin == nil || *p.OutputBin != *input.OutputBin) {
		return false
	}

	if p.PagesPerSheet != nil && (input.PagesPerSheet == nil || *p.PagesPerSheet != *input.PagesPerSheet) {
		return false
	}

	return true
}
