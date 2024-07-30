package entitlementmanagementroleassignmentscheduleinstance

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type UnifiedRoleAssignmentScheduleInstanceOperationPredicate struct {
}

func (p UnifiedRoleAssignmentScheduleInstanceOperationPredicate) Matches(input stable.UnifiedRoleAssignmentScheduleInstance) bool {

	return true
}
