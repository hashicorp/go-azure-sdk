package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus string

const (
	UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatusinsufficientData UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus = "InsufficientData"
	UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatusmeetingGoals     UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus = "MeetingGoals"
	UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatusneedsAttention   UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus = "NeedsAttention"
	UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatusunknown          UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus = "Unknown"
)

func PossibleValuesForUserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus() []string {
	return []string{
		string(UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatusinsufficientData),
		string(UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatusmeetingGoals),
		string(UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatusneedsAttention),
		string(UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatusunknown),
	}
}

func parseUserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus(input string) (*UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus, error) {
	vals := map[string]UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus{
		"insufficientdata": UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatusinsufficientData,
		"meetinggoals":     UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatusmeetingGoals,
		"needsattention":   UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatusneedsAttention,
		"unknown":          UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus(input)
	return &out, nil
}
