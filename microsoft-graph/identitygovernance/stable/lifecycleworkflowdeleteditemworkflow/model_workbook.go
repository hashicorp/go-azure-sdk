package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Workbook struct {
	Application *WorkbookApplication `json:"application,omitempty"`
	Comments    *[]WorkbookComment   `json:"comments,omitempty"`
	Functions   *WorkbookFunctions   `json:"functions,omitempty"`
	Id          *string              `json:"id,omitempty"`
	Names       *[]WorkbookNamedItem `json:"names,omitempty"`
	ODataType   *string              `json:"@odata.type,omitempty"`
	Operations  *[]WorkbookOperation `json:"operations,omitempty"`
	Tables      *[]WorkbookTable     `json:"tables,omitempty"`
	Worksheets  *[]WorkbookWorksheet `json:"worksheets,omitempty"`
}
