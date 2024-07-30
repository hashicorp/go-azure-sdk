package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultColumnValue struct {
	Formula   *string `json:"formula,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	Value     *string `json:"value,omitempty"`
}
