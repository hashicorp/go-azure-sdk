package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookSortField struct {
	Ascending  *bool         `json:"ascending,omitempty"`
	Color      *string       `json:"color,omitempty"`
	DataOption *string       `json:"dataOption,omitempty"`
	Icon       *WorkbookIcon `json:"icon,omitempty"`
	Key        *int64        `json:"key,omitempty"`
	ODataType  *string       `json:"@odata.type,omitempty"`
	SortOn     *string       `json:"sortOn,omitempty"`
}
