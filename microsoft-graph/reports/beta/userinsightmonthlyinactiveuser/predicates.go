package userinsightmonthlyinactiveuser

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type MonthlyInactiveUsersMetricOperationPredicate struct {
}

func (p MonthlyInactiveUsersMetricOperationPredicate) Matches(input beta.MonthlyInactiveUsersMetric) bool {

	return true
}
