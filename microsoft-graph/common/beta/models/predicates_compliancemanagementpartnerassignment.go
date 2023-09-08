package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComplianceManagementPartnerAssignmentOperationPredicate struct {
	ODataType *string
}

func (p ComplianceManagementPartnerAssignmentOperationPredicate) Matches(input ComplianceManagementPartnerAssignment) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
