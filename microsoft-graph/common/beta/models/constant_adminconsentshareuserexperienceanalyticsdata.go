package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdminConsentShareUserExperienceAnalyticsData string

const (
	AdminConsentShareUserExperienceAnalyticsDatagranted       AdminConsentShareUserExperienceAnalyticsData = "Granted"
	AdminConsentShareUserExperienceAnalyticsDatanotConfigured AdminConsentShareUserExperienceAnalyticsData = "NotConfigured"
	AdminConsentShareUserExperienceAnalyticsDatanotGranted    AdminConsentShareUserExperienceAnalyticsData = "NotGranted"
)

func PossibleValuesForAdminConsentShareUserExperienceAnalyticsData() []string {
	return []string{
		string(AdminConsentShareUserExperienceAnalyticsDatagranted),
		string(AdminConsentShareUserExperienceAnalyticsDatanotConfigured),
		string(AdminConsentShareUserExperienceAnalyticsDatanotGranted),
	}
}

func parseAdminConsentShareUserExperienceAnalyticsData(input string) (*AdminConsentShareUserExperienceAnalyticsData, error) {
	vals := map[string]AdminConsentShareUserExperienceAnalyticsData{
		"granted":       AdminConsentShareUserExperienceAnalyticsDatagranted,
		"notconfigured": AdminConsentShareUserExperienceAnalyticsDatanotConfigured,
		"notgranted":    AdminConsentShareUserExperienceAnalyticsDatanotGranted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AdminConsentShareUserExperienceAnalyticsData(input)
	return &out, nil
}
