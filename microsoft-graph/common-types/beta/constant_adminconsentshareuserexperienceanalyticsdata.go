package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdminConsentShareUserExperienceAnalyticsData string

const (
	AdminConsentShareUserExperienceAnalyticsData_Granted       AdminConsentShareUserExperienceAnalyticsData = "granted"
	AdminConsentShareUserExperienceAnalyticsData_NotConfigured AdminConsentShareUserExperienceAnalyticsData = "notConfigured"
	AdminConsentShareUserExperienceAnalyticsData_NotGranted    AdminConsentShareUserExperienceAnalyticsData = "notGranted"
)

func PossibleValuesForAdminConsentShareUserExperienceAnalyticsData() []string {
	return []string{
		string(AdminConsentShareUserExperienceAnalyticsData_Granted),
		string(AdminConsentShareUserExperienceAnalyticsData_NotConfigured),
		string(AdminConsentShareUserExperienceAnalyticsData_NotGranted),
	}
}

func (s *AdminConsentShareUserExperienceAnalyticsData) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAdminConsentShareUserExperienceAnalyticsData(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAdminConsentShareUserExperienceAnalyticsData(input string) (*AdminConsentShareUserExperienceAnalyticsData, error) {
	vals := map[string]AdminConsentShareUserExperienceAnalyticsData{
		"granted":       AdminConsentShareUserExperienceAnalyticsData_Granted,
		"notconfigured": AdminConsentShareUserExperienceAnalyticsData_NotConfigured,
		"notgranted":    AdminConsentShareUserExperienceAnalyticsData_NotGranted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AdminConsentShareUserExperienceAnalyticsData(input)
	return &out, nil
}
