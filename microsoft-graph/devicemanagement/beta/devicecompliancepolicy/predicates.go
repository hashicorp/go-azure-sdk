package devicecompliancepolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type DeviceCompliancePolicyOperationPredicate struct {
}

func (p DeviceCompliancePolicyOperationPredicate) Matches(input beta.DeviceCompliancePolicy) bool {

	return true
}

type DeviceCompliancePolicyAssignmentOperationPredicate struct {
}

func (p DeviceCompliancePolicyAssignmentOperationPredicate) Matches(input beta.DeviceCompliancePolicyAssignment) bool {

	return true
}

type HasPayloadLinkResultItemOperationPredicate struct {
}

func (p HasPayloadLinkResultItemOperationPredicate) Matches(input beta.HasPayloadLinkResultItem) bool {

	return true
}
