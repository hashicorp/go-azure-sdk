package devicecompliancepolicy

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type DeviceCompliancePolicyOperationPredicate struct {
}

func (p DeviceCompliancePolicyOperationPredicate) Matches(input stable.DeviceCompliancePolicy) bool {

	return true
}

type DeviceCompliancePolicyAssignmentOperationPredicate struct {
}

func (p DeviceCompliancePolicyAssignmentOperationPredicate) Matches(input stable.DeviceCompliancePolicyAssignment) bool {

	return true
}
