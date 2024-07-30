package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookTableRow struct {
	Id        *string `json:"id,omitempty"`
	Index     *int64  `json:"index,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	Values    *Json   `json:"values,omitempty"`
}
