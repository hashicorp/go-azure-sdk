package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCConnectivityResultOperationPredicate struct {
	ODataType       *string
	UpdatedDateTime *string
}

func (p CloudPCConnectivityResultOperationPredicate) Matches(input CloudPCConnectivityResult) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UpdatedDateTime != nil && (input.UpdatedDateTime == nil || *p.UpdatedDateTime != *input.UpdatedDateTime) {
		return false
	}

	return true
}
