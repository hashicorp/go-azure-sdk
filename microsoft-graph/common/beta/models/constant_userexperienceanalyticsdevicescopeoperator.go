package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceScopeOperator string

const (
	UserExperienceAnalyticsDeviceScopeOperatorequals UserExperienceAnalyticsDeviceScopeOperator = "Equals"
	UserExperienceAnalyticsDeviceScopeOperatornone   UserExperienceAnalyticsDeviceScopeOperator = "None"
)

func PossibleValuesForUserExperienceAnalyticsDeviceScopeOperator() []string {
	return []string{
		string(UserExperienceAnalyticsDeviceScopeOperatorequals),
		string(UserExperienceAnalyticsDeviceScopeOperatornone),
	}
}

func parseUserExperienceAnalyticsDeviceScopeOperator(input string) (*UserExperienceAnalyticsDeviceScopeOperator, error) {
	vals := map[string]UserExperienceAnalyticsDeviceScopeOperator{
		"equals": UserExperienceAnalyticsDeviceScopeOperatorequals,
		"none":   UserExperienceAnalyticsDeviceScopeOperatornone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsDeviceScopeOperator(input)
	return &out, nil
}
