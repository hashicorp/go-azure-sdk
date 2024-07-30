package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAnomalyDeviceDeviceStatus string

const (
	UserExperienceAnalyticsAnomalyDeviceDeviceStatus_Affected  UserExperienceAnalyticsAnomalyDeviceDeviceStatus = "affected"
	UserExperienceAnalyticsAnomalyDeviceDeviceStatus_Anomalous UserExperienceAnalyticsAnomalyDeviceDeviceStatus = "anomalous"
	UserExperienceAnalyticsAnomalyDeviceDeviceStatus_AtRisk    UserExperienceAnalyticsAnomalyDeviceDeviceStatus = "atRisk"
)

func PossibleValuesForUserExperienceAnalyticsAnomalyDeviceDeviceStatus() []string {
	return []string{
		string(UserExperienceAnalyticsAnomalyDeviceDeviceStatus_Affected),
		string(UserExperienceAnalyticsAnomalyDeviceDeviceStatus_Anomalous),
		string(UserExperienceAnalyticsAnomalyDeviceDeviceStatus_AtRisk),
	}
}

func (s *UserExperienceAnalyticsAnomalyDeviceDeviceStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserExperienceAnalyticsAnomalyDeviceDeviceStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserExperienceAnalyticsAnomalyDeviceDeviceStatus(input string) (*UserExperienceAnalyticsAnomalyDeviceDeviceStatus, error) {
	vals := map[string]UserExperienceAnalyticsAnomalyDeviceDeviceStatus{
		"affected":  UserExperienceAnalyticsAnomalyDeviceDeviceStatus_Affected,
		"anomalous": UserExperienceAnalyticsAnomalyDeviceDeviceStatus_Anomalous,
		"atrisk":    UserExperienceAnalyticsAnomalyDeviceDeviceStatus_AtRisk,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsAnomalyDeviceDeviceStatus(input)
	return &out, nil
}
