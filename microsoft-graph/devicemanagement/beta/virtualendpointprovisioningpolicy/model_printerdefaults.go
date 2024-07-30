package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterDefaults struct {
	ColorMode               *PrinterDefaultsColorMode               `json:"colorMode,omitempty"`
	ContentType             *string                                 `json:"contentType,omitempty"`
	CopiesPerJob            *int64                                  `json:"copiesPerJob,omitempty"`
	DocumentMimeType        *string                                 `json:"documentMimeType,omitempty"`
	Dpi                     *int64                                  `json:"dpi,omitempty"`
	DuplexConfiguration     *PrinterDefaultsDuplexConfiguration     `json:"duplexConfiguration,omitempty"`
	DuplexMode              *PrinterDefaultsDuplexMode              `json:"duplexMode,omitempty"`
	Finishings              *PrinterDefaultsFinishings              `json:"finishings,omitempty"`
	FitPdfToPage            *bool                                   `json:"fitPdfToPage,omitempty"`
	InputBin                *string                                 `json:"inputBin,omitempty"`
	MediaColor              *string                                 `json:"mediaColor,omitempty"`
	MediaSize               *string                                 `json:"mediaSize,omitempty"`
	MediaType               *string                                 `json:"mediaType,omitempty"`
	MultipageLayout         *PrinterDefaultsMultipageLayout         `json:"multipageLayout,omitempty"`
	ODataType               *string                                 `json:"@odata.type,omitempty"`
	Orientation             *PrinterDefaultsOrientation             `json:"orientation,omitempty"`
	OutputBin               *string                                 `json:"outputBin,omitempty"`
	PagesPerSheet           *int64                                  `json:"pagesPerSheet,omitempty"`
	PdfFitToPage            *bool                                   `json:"pdfFitToPage,omitempty"`
	PresentationDirection   *PrinterDefaultsPresentationDirection   `json:"presentationDirection,omitempty"`
	PrintColorConfiguration *PrinterDefaultsPrintColorConfiguration `json:"printColorConfiguration,omitempty"`
	PrintQuality            *PrinterDefaultsPrintQuality            `json:"printQuality,omitempty"`
	Quality                 *PrinterDefaultsQuality                 `json:"quality,omitempty"`
	Scaling                 *PrinterDefaultsScaling                 `json:"scaling,omitempty"`
}
