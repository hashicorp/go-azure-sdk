package applicationwhitelistings

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdaptiveApplicationControlGroupData struct {
	ConfigurationStatus  *ConfigurationStatus                      `json:"configurationStatus,omitempty"`
	EnforcementMode      *EnforcementMode                          `json:"enforcementMode,omitempty"`
	Issues               *[]AdaptiveApplicationControlIssueSummary `json:"issues,omitempty"`
	PathRecommendations  *[]PathRecommendation                     `json:"pathRecommendations,omitempty"`
	ProtectionMode       *ProtectionMode                           `json:"protectionMode,omitempty"`
	RecommendationStatus *RecommendationStatus                     `json:"recommendationStatus,omitempty"`
	SourceSystem         *SourceSystem                             `json:"sourceSystem,omitempty"`
	VMRecommendations    *[]VMRecommendation                       `json:"vmRecommendations,omitempty"`
}
