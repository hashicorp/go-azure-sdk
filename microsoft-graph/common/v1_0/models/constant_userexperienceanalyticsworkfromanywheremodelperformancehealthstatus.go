package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus string

const (
	UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatusinsufficientData UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus = "InsufficientData"
	UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatusmeetingGoals     UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus = "MeetingGoals"
	UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatusneedsAttention   UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus = "NeedsAttention"
	UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatusunknown          UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus = "Unknown"
)

func PossibleValuesForUserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus() []string {
	return []string{
		string(UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatusinsufficientData),
		string(UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatusmeetingGoals),
		string(UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatusneedsAttention),
		string(UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatusunknown),
	}
}

func parseUserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus(input string) (*UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus, error) {
	vals := map[string]UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus{
		"insufficientdata": UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatusinsufficientData,
		"meetinggoals":     UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatusmeetingGoals,
		"needsattention":   UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatusneedsAttention,
		"unknown":          UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus(input)
	return &out, nil
}
