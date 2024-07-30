package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserLastSignInRecommendationInsightSetting struct {
	ODataType                      *string                                                `json:"@odata.type,omitempty"`
	RecommendationLookBackDuration *string                                                `json:"recommendationLookBackDuration,omitempty"`
	SignInScope                    *UserLastSignInRecommendationInsightSettingSignInScope `json:"signInScope,omitempty"`
}
