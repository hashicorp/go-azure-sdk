package lifecycleworkflowworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookNamedItem struct {
	Comment   *string            `json:"comment,omitempty"`
	Id        *string            `json:"id,omitempty"`
	Name      *string            `json:"name,omitempty"`
	ODataType *string            `json:"@odata.type,omitempty"`
	Scope     *string            `json:"scope,omitempty"`
	Type      *string            `json:"type,omitempty"`
	Value     *Json              `json:"value,omitempty"`
	Visible   *bool              `json:"visible,omitempty"`
	Worksheet *WorkbookWorksheet `json:"worksheet,omitempty"`
}
