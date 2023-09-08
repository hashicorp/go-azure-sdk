package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessBranchConnectivityConfigurationOperationPredicate struct {
	BranchId   *string
	BranchName *string
	ODataType  *string
}

func (p NetworkaccessBranchConnectivityConfigurationOperationPredicate) Matches(input NetworkaccessBranchConnectivityConfiguration) bool {

	if p.BranchId != nil && (input.BranchId == nil || *p.BranchId != *input.BranchId) {
		return false
	}

	if p.BranchName != nil && (input.BranchName == nil || *p.BranchName != *input.BranchName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
