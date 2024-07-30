package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessBranchConnectivityConfiguration struct {
	BranchId   *string                                       `json:"branchId,omitempty"`
	BranchName *string                                       `json:"branchName,omitempty"`
	Links      *[]NetworkaccessConnectivityConfigurationLink `json:"links,omitempty"`
	ODataType  *string                                       `json:"@odata.type,omitempty"`
}
