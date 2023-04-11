package applicationwhitelistings

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserRecommendation struct {
	RecommendationAction *RecommendationAction `json:"recommendationAction,omitempty"`
	Username             *string               `json:"username,omitempty"`
}
