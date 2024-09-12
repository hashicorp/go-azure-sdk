package entitlementmanagementroleassignmentschedulerequest

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type UnifiedRoleAssignmentScheduleRequestOperationPredicate struct {
}

func (p UnifiedRoleAssignmentScheduleRequestOperationPredicate) Matches(input stable.UnifiedRoleAssignmentScheduleRequest) bool {

	return true
}
