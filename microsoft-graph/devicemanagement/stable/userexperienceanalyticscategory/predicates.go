package userexperienceanalyticscategory

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type UserExperienceAnalyticsCategoryOperationPredicate struct {
}

func (p UserExperienceAnalyticsCategoryOperationPredicate) Matches(input stable.UserExperienceAnalyticsCategory) bool {

	return true
}
