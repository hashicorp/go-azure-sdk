package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterCapabilities struct {
	BottomMargins                   *[]int64                                            `json:"bottomMargins,omitempty"`
	Collation                       *bool                                               `json:"collation,omitempty"`
	ColorModes                      *PrinterCapabilitiesColorModes                      `json:"colorModes,omitempty"`
	ContentTypes                    *[]string                                           `json:"contentTypes,omitempty"`
	CopiesPerJob                    *IntegerRange                                       `json:"copiesPerJob,omitempty"`
	Dpis                            *[]int64                                            `json:"dpis,omitempty"`
	DuplexModes                     *PrinterCapabilitiesDuplexModes                     `json:"duplexModes,omitempty"`
	FeedDirections                  *PrinterCapabilitiesFeedDirections                  `json:"feedDirections,omitempty"`
	FeedOrientations                *PrinterCapabilitiesFeedOrientations                `json:"feedOrientations,omitempty"`
	Finishings                      *PrinterCapabilitiesFinishings                      `json:"finishings,omitempty"`
	InputBins                       *[]string                                           `json:"inputBins,omitempty"`
	IsColorPrintingSupported        *bool                                               `json:"isColorPrintingSupported,omitempty"`
	IsPageRangeSupported            *bool                                               `json:"isPageRangeSupported,omitempty"`
	LeftMargins                     *[]int64                                            `json:"leftMargins,omitempty"`
	MediaColors                     *[]string                                           `json:"mediaColors,omitempty"`
	MediaSizes                      *[]string                                           `json:"mediaSizes,omitempty"`
	MediaTypes                      *[]string                                           `json:"mediaTypes,omitempty"`
	MultipageLayouts                *PrinterCapabilitiesMultipageLayouts                `json:"multipageLayouts,omitempty"`
	ODataType                       *string                                             `json:"@odata.type,omitempty"`
	Orientations                    *PrinterCapabilitiesOrientations                    `json:"orientations,omitempty"`
	OutputBins                      *[]string                                           `json:"outputBins,omitempty"`
	PagesPerSheet                   *[]int64                                            `json:"pagesPerSheet,omitempty"`
	Qualities                       *PrinterCapabilitiesQualities                       `json:"qualities,omitempty"`
	RightMargins                    *[]int64                                            `json:"rightMargins,omitempty"`
	Scalings                        *PrinterCapabilitiesScalings                        `json:"scalings,omitempty"`
	SupportedColorConfigurations    *PrinterCapabilitiesSupportedColorConfigurations    `json:"supportedColorConfigurations,omitempty"`
	SupportedCopiesPerJob           *IntegerRange                                       `json:"supportedCopiesPerJob,omitempty"`
	SupportedDocumentMimeTypes      *[]string                                           `json:"supportedDocumentMimeTypes,omitempty"`
	SupportedDuplexConfigurations   *PrinterCapabilitiesSupportedDuplexConfigurations   `json:"supportedDuplexConfigurations,omitempty"`
	SupportedFinishings             *PrinterCapabilitiesSupportedFinishings             `json:"supportedFinishings,omitempty"`
	SupportedMediaColors            *[]string                                           `json:"supportedMediaColors,omitempty"`
	SupportedMediaSizes             *[]string                                           `json:"supportedMediaSizes,omitempty"`
	SupportedMediaTypes             *PrinterCapabilitiesSupportedMediaTypes             `json:"supportedMediaTypes,omitempty"`
	SupportedOrientations           *PrinterCapabilitiesSupportedOrientations           `json:"supportedOrientations,omitempty"`
	SupportedOutputBins             *[]string                                           `json:"supportedOutputBins,omitempty"`
	SupportedPagesPerSheet          *IntegerRange                                       `json:"supportedPagesPerSheet,omitempty"`
	SupportedPresentationDirections *PrinterCapabilitiesSupportedPresentationDirections `json:"supportedPresentationDirections,omitempty"`
	SupportedPrintQualities         *PrinterCapabilitiesSupportedPrintQualities         `json:"supportedPrintQualities,omitempty"`
	SupportsFitPdfToPage            *bool                                               `json:"supportsFitPdfToPage,omitempty"`
	TopMargins                      *[]int64                                            `json:"topMargins,omitempty"`
}
