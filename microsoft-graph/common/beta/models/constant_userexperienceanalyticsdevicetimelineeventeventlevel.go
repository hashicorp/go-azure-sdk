package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceTimelineEventEventLevel string

const (
	UserExperienceAnalyticsDeviceTimelineEventEventLevelcritical    UserExperienceAnalyticsDeviceTimelineEventEventLevel = "Critical"
	UserExperienceAnalyticsDeviceTimelineEventEventLevelerror       UserExperienceAnalyticsDeviceTimelineEventEventLevel = "Error"
	UserExperienceAnalyticsDeviceTimelineEventEventLevelinformation UserExperienceAnalyticsDeviceTimelineEventEventLevel = "Information"
	UserExperienceAnalyticsDeviceTimelineEventEventLevelnone        UserExperienceAnalyticsDeviceTimelineEventEventLevel = "None"
	UserExperienceAnalyticsDeviceTimelineEventEventLevelverbose     UserExperienceAnalyticsDeviceTimelineEventEventLevel = "Verbose"
	UserExperienceAnalyticsDeviceTimelineEventEventLevelwarning     UserExperienceAnalyticsDeviceTimelineEventEventLevel = "Warning"
)

func PossibleValuesForUserExperienceAnalyticsDeviceTimelineEventEventLevel() []string {
	return []string{
		string(UserExperienceAnalyticsDeviceTimelineEventEventLevelcritical),
		string(UserExperienceAnalyticsDeviceTimelineEventEventLevelerror),
		string(UserExperienceAnalyticsDeviceTimelineEventEventLevelinformation),
		string(UserExperienceAnalyticsDeviceTimelineEventEventLevelnone),
		string(UserExperienceAnalyticsDeviceTimelineEventEventLevelverbose),
		string(UserExperienceAnalyticsDeviceTimelineEventEventLevelwarning),
	}
}

func parseUserExperienceAnalyticsDeviceTimelineEventEventLevel(input string) (*UserExperienceAnalyticsDeviceTimelineEventEventLevel, error) {
	vals := map[string]UserExperienceAnalyticsDeviceTimelineEventEventLevel{
		"critical":    UserExperienceAnalyticsDeviceTimelineEventEventLevelcritical,
		"error":       UserExperienceAnalyticsDeviceTimelineEventEventLevelerror,
		"information": UserExperienceAnalyticsDeviceTimelineEventEventLevelinformation,
		"none":        UserExperienceAnalyticsDeviceTimelineEventEventLevelnone,
		"verbose":     UserExperienceAnalyticsDeviceTimelineEventEventLevelverbose,
		"warning":     UserExperienceAnalyticsDeviceTimelineEventEventLevelwarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsDeviceTimelineEventEventLevel(input)
	return &out, nil
}
