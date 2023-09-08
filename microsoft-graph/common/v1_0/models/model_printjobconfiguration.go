package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintJobConfiguration struct {
	Collate         *bool                                 `json:"collate,omitempty"`
	ColorMode       *PrintJobConfigurationColorMode       `json:"colorMode,omitempty"`
	Copies          *int64                                `json:"copies,omitempty"`
	Dpi             *int64                                `json:"dpi,omitempty"`
	DuplexMode      *PrintJobConfigurationDuplexMode      `json:"duplexMode,omitempty"`
	FeedOrientation *PrintJobConfigurationFeedOrientation `json:"feedOrientation,omitempty"`
	Finishings      *[]PrintJobConfigurationFinishings    `json:"finishings,omitempty"`
	FitPdfToPage    *bool                                 `json:"fitPdfToPage,omitempty"`
	InputBin        *string                               `json:"inputBin,omitempty"`
	Margin          *PrintMargin                          `json:"margin,omitempty"`
	MediaSize       *string                               `json:"mediaSize,omitempty"`
	MediaType       *string                               `json:"mediaType,omitempty"`
	MultipageLayout *PrintJobConfigurationMultipageLayout `json:"multipageLayout,omitempty"`
	ODataType       *string                               `json:"@odata.type,omitempty"`
	Orientation     *PrintJobConfigurationOrientation     `json:"orientation,omitempty"`
	OutputBin       *string                               `json:"outputBin,omitempty"`
	PageRanges      *[]IntegerRange                       `json:"pageRanges,omitempty"`
	PagesPerSheet   *int64                                `json:"pagesPerSheet,omitempty"`
	Quality         *PrintJobConfigurationQuality         `json:"quality,omitempty"`
	Scaling         *PrintJobConfigurationScaling         `json:"scaling,omitempty"`
}
