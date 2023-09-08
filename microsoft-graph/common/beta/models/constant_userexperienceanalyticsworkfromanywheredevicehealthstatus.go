package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus string

const (
	UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatusinsufficientData UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus = "InsufficientData"
	UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatusmeetingGoals     UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus = "MeetingGoals"
	UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatusneedsAttention   UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus = "NeedsAttention"
	UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatusunknown          UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus = "Unknown"
)

func PossibleValuesForUserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus() []string {
	return []string{
		string(UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatusinsufficientData),
		string(UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatusmeetingGoals),
		string(UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatusneedsAttention),
		string(UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatusunknown),
	}
}

func parseUserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus(input string) (*UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus, error) {
	vals := map[string]UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus{
		"insufficientdata": UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatusinsufficientData,
		"meetinggoals":     UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatusmeetingGoals,
		"needsattention":   UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatusneedsAttention,
		"unknown":          UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus(input)
	return &out, nil
}
