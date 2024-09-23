package userexperienceanalyticsdevicemetrichistory

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type UserExperienceAnalyticsMetricHistoryOperationPredicate struct {
}

func (p UserExperienceAnalyticsMetricHistoryOperationPredicate) Matches(input beta.UserExperienceAnalyticsMetricHistory) bool {

	return true
}
