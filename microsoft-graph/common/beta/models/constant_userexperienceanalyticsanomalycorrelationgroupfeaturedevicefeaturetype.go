package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType string

const (
	UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureTypeapplication  UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType = "Application"
	UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureTypedriver       UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType = "Driver"
	UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureTypemanufacturer UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType = "Manufacturer"
	UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureTypemodel        UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType = "Model"
	UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureTypeosVersion    UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType = "OsVersion"
)

func PossibleValuesForUserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType() []string {
	return []string{
		string(UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureTypeapplication),
		string(UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureTypedriver),
		string(UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureTypemanufacturer),
		string(UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureTypemodel),
		string(UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureTypeosVersion),
	}
}

func parseUserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType(input string) (*UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType, error) {
	vals := map[string]UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType{
		"application":  UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureTypeapplication,
		"driver":       UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureTypedriver,
		"manufacturer": UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureTypemanufacturer,
		"model":        UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureTypemodel,
		"osversion":    UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureTypeosVersion,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsAnomalyCorrelationGroupFeatureDeviceFeatureType(input)
	return &out, nil
}
