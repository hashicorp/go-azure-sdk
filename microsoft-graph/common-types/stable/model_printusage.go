package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintUsage struct {
	BlackAndWhitePageCount         *int64  `json:"blackAndWhitePageCount,omitempty"`
	ColorPageCount                 *int64  `json:"colorPageCount,omitempty"`
	CompletedBlackAndWhiteJobCount *int64  `json:"completedBlackAndWhiteJobCount,omitempty"`
	CompletedColorJobCount         *int64  `json:"completedColorJobCount,omitempty"`
	CompletedJobCount              *int64  `json:"completedJobCount,omitempty"`
	DoubleSidedSheetCount          *int64  `json:"doubleSidedSheetCount,omitempty"`
	Id                             *string `json:"id,omitempty"`
	IncompleteJobCount             *int64  `json:"incompleteJobCount,omitempty"`
	MediaSheetCount                *int64  `json:"mediaSheetCount,omitempty"`
	ODataType                      *string `json:"@odata.type,omitempty"`
	PageCount                      *int64  `json:"pageCount,omitempty"`
	SingleSidedSheetCount          *int64  `json:"singleSidedSheetCount,omitempty"`
	UsageDate                      *string `json:"usageDate,omitempty"`
}
