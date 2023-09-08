package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus string

const (
	UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatusinsufficientData UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus = "InsufficientData"
	UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatusmeetingGoals     UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus = "MeetingGoals"
	UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatusneedsAttention   UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus = "NeedsAttention"
	UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatusunknown          UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus = "Unknown"
)

func PossibleValuesForUserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus() []string {
	return []string{
		string(UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatusinsufficientData),
		string(UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatusmeetingGoals),
		string(UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatusneedsAttention),
		string(UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatusunknown),
	}
}

func parseUserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus(input string) (*UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus, error) {
	vals := map[string]UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus{
		"insufficientdata": UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatusinsufficientData,
		"meetinggoals":     UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatusmeetingGoals,
		"needsattention":   UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatusneedsAttention,
		"unknown":          UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus(input)
	return &out, nil
}
