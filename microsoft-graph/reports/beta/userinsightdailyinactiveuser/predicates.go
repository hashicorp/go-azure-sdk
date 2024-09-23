package userinsightdailyinactiveuser

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type DailyInactiveUsersMetricOperationPredicate struct {
}

func (p DailyInactiveUsersMetricOperationPredicate) Matches(input beta.DailyInactiveUsersMetric) bool {

	return true
}
