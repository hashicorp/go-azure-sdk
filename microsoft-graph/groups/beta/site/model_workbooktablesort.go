package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookTableSort struct {
	Fields    *[]WorkbookSortField `json:"fields,omitempty"`
	Id        *string              `json:"id,omitempty"`
	MatchCase *bool                `json:"matchCase,omitempty"`
	Method    *string              `json:"method,omitempty"`
	ODataType *string              `json:"@odata.type,omitempty"`
}
