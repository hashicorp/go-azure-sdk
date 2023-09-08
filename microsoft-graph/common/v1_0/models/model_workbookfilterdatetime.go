package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookFilterDatetime struct {
	Date        *string `json:"date,omitempty"`
	ODataType   *string `json:"@odata.type,omitempty"`
	Specificity *string `json:"specificity,omitempty"`
}
