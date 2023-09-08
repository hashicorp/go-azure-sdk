package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserLastSignInRecommendationInsightSettingOperationPredicate struct {
	ODataType                      *string
	RecommendationLookBackDuration *string
}

func (p UserLastSignInRecommendationInsightSettingOperationPredicate) Matches(input UserLastSignInRecommendationInsightSetting) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RecommendationLookBackDuration != nil && (input.RecommendationLookBackDuration == nil || *p.RecommendationLookBackDuration != *input.RecommendationLookBackDuration) {
		return false
	}

	return true
}
