package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookFilterCriteria struct {
	Color           *string       `json:"color,omitempty"`
	Criterion1      *string       `json:"criterion1,omitempty"`
	Criterion2      *string       `json:"criterion2,omitempty"`
	DynamicCriteria *string       `json:"dynamicCriteria,omitempty"`
	FilterOn        *string       `json:"filterOn,omitempty"`
	Icon            *WorkbookIcon `json:"icon,omitempty"`
	ODataType       *string       `json:"@odata.type,omitempty"`
	Operator        *string       `json:"operator,omitempty"`
	Values          *Json         `json:"values,omitempty"`
}
