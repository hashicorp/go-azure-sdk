package userexperienceanalyticsanomalydevice

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type UserExperienceAnalyticsAnomalyDeviceOperationPredicate struct {
}

func (p UserExperienceAnalyticsAnomalyDeviceOperationPredicate) Matches(input beta.UserExperienceAnalyticsAnomalyDevice) bool {

	return true
}
