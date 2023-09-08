package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessConnectivity struct {
	Branches  *[]NetworkaccessBranchSite `json:"branches,omitempty"`
	Id        *string                    `json:"id,omitempty"`
	ODataType *string                    `json:"@odata.type,omitempty"`
}
