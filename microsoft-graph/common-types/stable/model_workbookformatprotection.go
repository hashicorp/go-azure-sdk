package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookFormatProtection struct {
	FormulaHidden *bool   `json:"formulaHidden,omitempty"`
	Id            *string `json:"id,omitempty"`
	Locked        *bool   `json:"locked,omitempty"`
	ODataType     *string `json:"@odata.type,omitempty"`
}
