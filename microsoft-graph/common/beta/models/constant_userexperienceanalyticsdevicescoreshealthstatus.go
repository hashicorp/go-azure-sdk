package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceScoresHealthStatus string

const (
	UserExperienceAnalyticsDeviceScoresHealthStatusinsufficientData UserExperienceAnalyticsDeviceScoresHealthStatus = "InsufficientData"
	UserExperienceAnalyticsDeviceScoresHealthStatusmeetingGoals     UserExperienceAnalyticsDeviceScoresHealthStatus = "MeetingGoals"
	UserExperienceAnalyticsDeviceScoresHealthStatusneedsAttention   UserExperienceAnalyticsDeviceScoresHealthStatus = "NeedsAttention"
	UserExperienceAnalyticsDeviceScoresHealthStatusunknown          UserExperienceAnalyticsDeviceScoresHealthStatus = "Unknown"
)

func PossibleValuesForUserExperienceAnalyticsDeviceScoresHealthStatus() []string {
	return []string{
		string(UserExperienceAnalyticsDeviceScoresHealthStatusinsufficientData),
		string(UserExperienceAnalyticsDeviceScoresHealthStatusmeetingGoals),
		string(UserExperienceAnalyticsDeviceScoresHealthStatusneedsAttention),
		string(UserExperienceAnalyticsDeviceScoresHealthStatusunknown),
	}
}

func parseUserExperienceAnalyticsDeviceScoresHealthStatus(input string) (*UserExperienceAnalyticsDeviceScoresHealthStatus, error) {
	vals := map[string]UserExperienceAnalyticsDeviceScoresHealthStatus{
		"insufficientdata": UserExperienceAnalyticsDeviceScoresHealthStatusinsufficientData,
		"meetinggoals":     UserExperienceAnalyticsDeviceScoresHealthStatusmeetingGoals,
		"needsattention":   UserExperienceAnalyticsDeviceScoresHealthStatusneedsAttention,
		"unknown":          UserExperienceAnalyticsDeviceScoresHealthStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsDeviceScoresHealthStatus(input)
	return &out, nil
}
