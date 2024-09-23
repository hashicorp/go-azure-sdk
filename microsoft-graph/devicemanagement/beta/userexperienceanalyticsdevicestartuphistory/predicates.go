package userexperienceanalyticsdevicestartuphistory

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type UserExperienceAnalyticsDeviceStartupHistoryOperationPredicate struct {
}

func (p UserExperienceAnalyticsDeviceStartupHistoryOperationPredicate) Matches(input beta.UserExperienceAnalyticsDeviceStartupHistory) bool {

	return true
}
