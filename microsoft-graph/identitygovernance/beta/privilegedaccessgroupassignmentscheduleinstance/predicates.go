package privilegedaccessgroupassignmentscheduleinstance

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type PrivilegedAccessGroupAssignmentScheduleInstanceOperationPredicate struct {
}

func (p PrivilegedAccessGroupAssignmentScheduleInstanceOperationPredicate) Matches(input beta.PrivilegedAccessGroupAssignmentScheduleInstance) bool {

	return true
}
