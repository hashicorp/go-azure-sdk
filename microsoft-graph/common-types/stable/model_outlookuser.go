package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutlookUser struct {
	Id               *string            `json:"id,omitempty"`
	MasterCategories *[]OutlookCategory `json:"masterCategories,omitempty"`
	ODataType        *string            `json:"@odata.type,omitempty"`
}
