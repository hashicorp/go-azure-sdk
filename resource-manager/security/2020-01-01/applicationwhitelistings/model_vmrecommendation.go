package applicationwhitelistings

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VMRecommendation struct {
	ConfigurationStatus  *ConfigurationStatus  `json:"configurationStatus,omitempty"`
	EnforcementSupport   *EnforcementSupport   `json:"enforcementSupport,omitempty"`
	RecommendationAction *RecommendationAction `json:"recommendationAction,omitempty"`
	ResourceId           *string               `json:"resourceId,omitempty"`
}
