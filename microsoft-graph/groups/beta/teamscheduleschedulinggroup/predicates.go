package teamscheduleschedulinggroup

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type SchedulingGroupOperationPredicate struct {
}

func (p SchedulingGroupOperationPredicate) Matches(input beta.SchedulingGroup) bool {

	return true
}
