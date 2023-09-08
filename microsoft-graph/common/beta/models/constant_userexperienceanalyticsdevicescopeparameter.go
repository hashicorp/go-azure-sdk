package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceScopeParameter string

const (
	UserExperienceAnalyticsDeviceScopeParameternone     UserExperienceAnalyticsDeviceScopeParameter = "None"
	UserExperienceAnalyticsDeviceScopeParameterscopeTag UserExperienceAnalyticsDeviceScopeParameter = "ScopeTag"
)

func PossibleValuesForUserExperienceAnalyticsDeviceScopeParameter() []string {
	return []string{
		string(UserExperienceAnalyticsDeviceScopeParameternone),
		string(UserExperienceAnalyticsDeviceScopeParameterscopeTag),
	}
}

func parseUserExperienceAnalyticsDeviceScopeParameter(input string) (*UserExperienceAnalyticsDeviceScopeParameter, error) {
	vals := map[string]UserExperienceAnalyticsDeviceScopeParameter{
		"none":     UserExperienceAnalyticsDeviceScopeParameternone,
		"scopetag": UserExperienceAnalyticsDeviceScopeParameterscopeTag,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsDeviceScopeParameter(input)
	return &out, nil
}
