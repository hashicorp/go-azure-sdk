package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CollapseProperty struct {
	Fields    *[]string `json:"fields,omitempty"`
	Limit     *int64    `json:"limit,omitempty"`
	ODataType *string   `json:"@odata.type,omitempty"`
}
