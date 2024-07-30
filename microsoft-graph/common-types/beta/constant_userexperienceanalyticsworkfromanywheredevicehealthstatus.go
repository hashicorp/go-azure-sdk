package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus string

const (
	UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus_InsufficientData UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus = "insufficientData"
	UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus_MeetingGoals     UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus = "meetingGoals"
	UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus_NeedsAttention   UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus = "needsAttention"
	UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus_Unknown          UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus = "unknown"
)

func PossibleValuesForUserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus() []string {
	return []string{
		string(UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus_InsufficientData),
		string(UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus_MeetingGoals),
		string(UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus_NeedsAttention),
		string(UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus_Unknown),
	}
}

func (s *UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus(input string) (*UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus, error) {
	vals := map[string]UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus{
		"insufficientdata": UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus_InsufficientData,
		"meetinggoals":     UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus_MeetingGoals,
		"needsattention":   UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus_NeedsAttention,
		"unknown":          UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsWorkFromAnywhereDeviceHealthStatus(input)
	return &out, nil
}
