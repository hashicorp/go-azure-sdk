package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NumberColumn struct {
	DecimalPlaces *string  `json:"decimalPlaces,omitempty"`
	DisplayAs     *string  `json:"displayAs,omitempty"`
	Maximum       *float64 `json:"maximum,omitempty"`
	Minimum       *float64 `json:"minimum,omitempty"`
	ODataType     *string  `json:"@odata.type,omitempty"`
}
