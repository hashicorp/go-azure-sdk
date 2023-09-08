package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NumberColumn struct {
	DecimalPlaces *string `json:"decimalPlaces,omitempty"`
	DisplayAs     *string `json:"displayAs,omitempty"`
	ODataType     *string `json:"@odata.type,omitempty"`
}
