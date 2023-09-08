package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus string

const (
	UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatusinsufficientData UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus = "InsufficientData"
	UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatusmeetingGoals     UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus = "MeetingGoals"
	UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatusneedsAttention   UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus = "NeedsAttention"
	UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatusunknown          UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus = "Unknown"
)

func PossibleValuesForUserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus() []string {
	return []string{
		string(UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatusinsufficientData),
		string(UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatusmeetingGoals),
		string(UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatusneedsAttention),
		string(UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatusunknown),
	}
}

func parseUserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus(input string) (*UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus, error) {
	vals := map[string]UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus{
		"insufficientdata": UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatusinsufficientData,
		"meetinggoals":     UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatusmeetingGoals,
		"needsattention":   UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatusneedsAttention,
		"unknown":          UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus(input)
	return &out, nil
}
