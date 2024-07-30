package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DailyUserInsightMetricsRoot struct {
	ActiveUsers                *[]ActiveUsersMetric                     `json:"activeUsers,omitempty"`
	ActiveUsersBreakdown       *[]ActiveUsersBreakdownMetric            `json:"activeUsersBreakdown,omitempty"`
	Authentications            *[]AuthenticationsMetric                 `json:"authentications,omitempty"`
	Id                         *string                                  `json:"id,omitempty"`
	InactiveUsers              *[]DailyInactiveUsersMetric              `json:"inactiveUsers,omitempty"`
	InactiveUsersByApplication *[]DailyInactiveUsersByApplicationMetric `json:"inactiveUsersByApplication,omitempty"`
	MfaCompletions             *[]MfaCompletionMetric                   `json:"mfaCompletions,omitempty"`
	ODataType                  *string                                  `json:"@odata.type,omitempty"`
	SignUps                    *[]UserSignUpMetric                      `json:"signUps,omitempty"`
	Summary                    *[]InsightSummary                        `json:"summary,omitempty"`
	UserCount                  *[]UserCountMetric                       `json:"userCount,omitempty"`
}
