package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDocumentConfiguration struct {
	Collate         *bool                                        `json:"collate,omitempty"`
	ColorMode       *PrinterDocumentConfigurationColorMode       `json:"colorMode,omitempty"`
	Copies          *int64                                       `json:"copies,omitempty"`
	Dpi             *int64                                       `json:"dpi,omitempty"`
	DuplexMode      *PrinterDocumentConfigurationDuplexMode      `json:"duplexMode,omitempty"`
	FeedDirection   *PrinterDocumentConfigurationFeedDirection   `json:"feedDirection,omitempty"`
	FeedOrientation *PrinterDocumentConfigurationFeedOrientation `json:"feedOrientation,omitempty"`
	Finishings      *PrinterDocumentConfigurationFinishings      `json:"finishings,omitempty"`
	FitPdfToPage    *bool                                        `json:"fitPdfToPage,omitempty"`
	InputBin        *string                                      `json:"inputBin,omitempty"`
	Margin          *PrintMargin                                 `json:"margin,omitempty"`
	MediaSize       *string                                      `json:"mediaSize,omitempty"`
	MediaType       *string                                      `json:"mediaType,omitempty"`
	MultipageLayout *PrinterDocumentConfigurationMultipageLayout `json:"multipageLayout,omitempty"`
	ODataType       *string                                      `json:"@odata.type,omitempty"`
	Orientation     *PrinterDocumentConfigurationOrientation     `json:"orientation,omitempty"`
	OutputBin       *string                                      `json:"outputBin,omitempty"`
	PageRanges      *[]IntegerRange                              `json:"pageRanges,omitempty"`
	PagesPerSheet   *int64                                       `json:"pagesPerSheet,omitempty"`
	Quality         *PrinterDocumentConfigurationQuality         `json:"quality,omitempty"`
	Scaling         *PrinterDocumentConfigurationScaling         `json:"scaling,omitempty"`
}
