package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus string

const (
	UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus_InsufficientData UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus = "insufficientData"
	UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus_MeetingGoals     UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus = "meetingGoals"
	UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus_NeedsAttention   UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus = "needsAttention"
	UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus_Unknown          UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus = "unknown"
)

func PossibleValuesForUserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus() []string {
	return []string{
		string(UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus_InsufficientData),
		string(UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus_MeetingGoals),
		string(UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus_NeedsAttention),
		string(UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus_Unknown),
	}
}

func (s *UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus(input string) (*UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus, error) {
	vals := map[string]UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus{
		"insufficientdata": UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus_InsufficientData,
		"meetinggoals":     UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus_MeetingGoals,
		"needsattention":   UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus_NeedsAttention,
		"unknown":          UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsAppHealthDevicePerformanceHealthStatus(input)
	return &out, nil
}
