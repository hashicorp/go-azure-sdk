package directoryroleeligibilityscheduleinstance

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type UnifiedRoleEligibilityScheduleInstanceOperationPredicate struct {
}

func (p UnifiedRoleEligibilityScheduleInstanceOperationPredicate) Matches(input stable.UnifiedRoleEligibilityScheduleInstance) bool {

	return true
}
