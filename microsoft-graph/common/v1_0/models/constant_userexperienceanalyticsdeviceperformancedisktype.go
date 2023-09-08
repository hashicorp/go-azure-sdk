package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDevicePerformanceDiskType string

const (
	UserExperienceAnalyticsDevicePerformanceDiskTypehdd     UserExperienceAnalyticsDevicePerformanceDiskType = "Hdd"
	UserExperienceAnalyticsDevicePerformanceDiskTypessd     UserExperienceAnalyticsDevicePerformanceDiskType = "Ssd"
	UserExperienceAnalyticsDevicePerformanceDiskTypeunknown UserExperienceAnalyticsDevicePerformanceDiskType = "Unknown"
)

func PossibleValuesForUserExperienceAnalyticsDevicePerformanceDiskType() []string {
	return []string{
		string(UserExperienceAnalyticsDevicePerformanceDiskTypehdd),
		string(UserExperienceAnalyticsDevicePerformanceDiskTypessd),
		string(UserExperienceAnalyticsDevicePerformanceDiskTypeunknown),
	}
}

func parseUserExperienceAnalyticsDevicePerformanceDiskType(input string) (*UserExperienceAnalyticsDevicePerformanceDiskType, error) {
	vals := map[string]UserExperienceAnalyticsDevicePerformanceDiskType{
		"hdd":     UserExperienceAnalyticsDevicePerformanceDiskTypehdd,
		"ssd":     UserExperienceAnalyticsDevicePerformanceDiskTypessd,
		"unknown": UserExperienceAnalyticsDevicePerformanceDiskTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsDevicePerformanceDiskType(input)
	return &out, nil
}
