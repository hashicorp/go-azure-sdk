package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookTableColumn struct {
	Filter    *WorkbookFilter `json:"filter,omitempty"`
	Id        *string         `json:"id,omitempty"`
	Index     *int64          `json:"index,omitempty"`
	Name      *string         `json:"name,omitempty"`
	ODataType *string         `json:"@odata.type,omitempty"`
	Values    *Json           `json:"values,omitempty"`
}
