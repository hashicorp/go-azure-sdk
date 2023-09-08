package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceScopeStatus string

const (
	UserExperienceAnalyticsDeviceScopeStatuscompleted        UserExperienceAnalyticsDeviceScopeStatus = "Completed"
	UserExperienceAnalyticsDeviceScopeStatuscomputing        UserExperienceAnalyticsDeviceScopeStatus = "Computing"
	UserExperienceAnalyticsDeviceScopeStatusinsufficientData UserExperienceAnalyticsDeviceScopeStatus = "InsufficientData"
	UserExperienceAnalyticsDeviceScopeStatusnone             UserExperienceAnalyticsDeviceScopeStatus = "None"
)

func PossibleValuesForUserExperienceAnalyticsDeviceScopeStatus() []string {
	return []string{
		string(UserExperienceAnalyticsDeviceScopeStatuscompleted),
		string(UserExperienceAnalyticsDeviceScopeStatuscomputing),
		string(UserExperienceAnalyticsDeviceScopeStatusinsufficientData),
		string(UserExperienceAnalyticsDeviceScopeStatusnone),
	}
}

func parseUserExperienceAnalyticsDeviceScopeStatus(input string) (*UserExperienceAnalyticsDeviceScopeStatus, error) {
	vals := map[string]UserExperienceAnalyticsDeviceScopeStatus{
		"completed":        UserExperienceAnalyticsDeviceScopeStatuscompleted,
		"computing":        UserExperienceAnalyticsDeviceScopeStatuscomputing,
		"insufficientdata": UserExperienceAnalyticsDeviceScopeStatusinsufficientData,
		"none":             UserExperienceAnalyticsDeviceScopeStatusnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsDeviceScopeStatus(input)
	return &out, nil
}
