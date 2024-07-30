package devicecompliancepolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type DeviceCompliancePolicyOperationPredicate struct {
}

func (p DeviceCompliancePolicyOperationPredicate) Matches(input stable.DeviceCompliancePolicy) bool {

	return true
}

type DeviceCompliancePolicyAssignmentOperationPredicate struct {
	Id        *string
	ODataType *string
}

func (p DeviceCompliancePolicyAssignmentOperationPredicate) Matches(input stable.DeviceCompliancePolicyAssignment) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
