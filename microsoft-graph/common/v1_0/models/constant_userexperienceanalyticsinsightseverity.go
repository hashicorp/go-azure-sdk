package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsInsightSeverity string

const (
	UserExperienceAnalyticsInsightSeverityerror         UserExperienceAnalyticsInsightSeverity = "Error"
	UserExperienceAnalyticsInsightSeverityinformational UserExperienceAnalyticsInsightSeverity = "Informational"
	UserExperienceAnalyticsInsightSeveritynone          UserExperienceAnalyticsInsightSeverity = "None"
	UserExperienceAnalyticsInsightSeveritywarning       UserExperienceAnalyticsInsightSeverity = "Warning"
)

func PossibleValuesForUserExperienceAnalyticsInsightSeverity() []string {
	return []string{
		string(UserExperienceAnalyticsInsightSeverityerror),
		string(UserExperienceAnalyticsInsightSeverityinformational),
		string(UserExperienceAnalyticsInsightSeveritynone),
		string(UserExperienceAnalyticsInsightSeveritywarning),
	}
}

func parseUserExperienceAnalyticsInsightSeverity(input string) (*UserExperienceAnalyticsInsightSeverity, error) {
	vals := map[string]UserExperienceAnalyticsInsightSeverity{
		"error":         UserExperienceAnalyticsInsightSeverityerror,
		"informational": UserExperienceAnalyticsInsightSeverityinformational,
		"none":          UserExperienceAnalyticsInsightSeveritynone,
		"warning":       UserExperienceAnalyticsInsightSeveritywarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsInsightSeverity(input)
	return &out, nil
}
