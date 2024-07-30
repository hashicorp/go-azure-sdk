package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus string

const (
	UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus_InsufficientData UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus = "insufficientData"
	UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus_MeetingGoals     UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus = "meetingGoals"
	UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus_NeedsAttention   UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus = "needsAttention"
	UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus_Unknown          UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus = "unknown"
)

func PossibleValuesForUserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus() []string {
	return []string{
		string(UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus_InsufficientData),
		string(UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus_MeetingGoals),
		string(UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus_NeedsAttention),
		string(UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus_Unknown),
	}
}

func (s *UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus(input string) (*UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus, error) {
	vals := map[string]UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus{
		"insufficientdata": UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus_InsufficientData,
		"meetinggoals":     UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus_MeetingGoals,
		"needsattention":   UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus_NeedsAttention,
		"unknown":          UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsAppHealthDeviceModelPerformanceHealthStatus(input)
	return &out, nil
}
