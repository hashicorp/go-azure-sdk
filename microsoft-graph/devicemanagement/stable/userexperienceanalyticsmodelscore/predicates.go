package userexperienceanalyticsmodelscore

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"

type UserExperienceAnalyticsModelScoresOperationPredicate struct {
}

func (p UserExperienceAnalyticsModelScoresOperationPredicate) Matches(input stable.UserExperienceAnalyticsModelScores) bool {

	return true
}
