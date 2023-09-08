package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAnomalySeverity string

const (
	UserExperienceAnalyticsAnomalySeverityhigh          UserExperienceAnalyticsAnomalySeverity = "High"
	UserExperienceAnalyticsAnomalySeverityinformational UserExperienceAnalyticsAnomalySeverity = "Informational"
	UserExperienceAnalyticsAnomalySeveritylow           UserExperienceAnalyticsAnomalySeverity = "Low"
	UserExperienceAnalyticsAnomalySeveritymedium        UserExperienceAnalyticsAnomalySeverity = "Medium"
	UserExperienceAnalyticsAnomalySeverityother         UserExperienceAnalyticsAnomalySeverity = "Other"
)

func PossibleValuesForUserExperienceAnalyticsAnomalySeverity() []string {
	return []string{
		string(UserExperienceAnalyticsAnomalySeverityhigh),
		string(UserExperienceAnalyticsAnomalySeverityinformational),
		string(UserExperienceAnalyticsAnomalySeveritylow),
		string(UserExperienceAnalyticsAnomalySeveritymedium),
		string(UserExperienceAnalyticsAnomalySeverityother),
	}
}

func parseUserExperienceAnalyticsAnomalySeverity(input string) (*UserExperienceAnalyticsAnomalySeverity, error) {
	vals := map[string]UserExperienceAnalyticsAnomalySeverity{
		"high":          UserExperienceAnalyticsAnomalySeverityhigh,
		"informational": UserExperienceAnalyticsAnomalySeverityinformational,
		"low":           UserExperienceAnalyticsAnomalySeveritylow,
		"medium":        UserExperienceAnalyticsAnomalySeveritymedium,
		"other":         UserExperienceAnalyticsAnomalySeverityother,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsAnomalySeverity(input)
	return &out, nil
}
