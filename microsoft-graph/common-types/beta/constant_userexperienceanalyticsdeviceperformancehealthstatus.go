package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDevicePerformanceHealthStatus string

const (
	UserExperienceAnalyticsDevicePerformanceHealthStatus_InsufficientData UserExperienceAnalyticsDevicePerformanceHealthStatus = "insufficientData"
	UserExperienceAnalyticsDevicePerformanceHealthStatus_MeetingGoals     UserExperienceAnalyticsDevicePerformanceHealthStatus = "meetingGoals"
	UserExperienceAnalyticsDevicePerformanceHealthStatus_NeedsAttention   UserExperienceAnalyticsDevicePerformanceHealthStatus = "needsAttention"
	UserExperienceAnalyticsDevicePerformanceHealthStatus_Unknown          UserExperienceAnalyticsDevicePerformanceHealthStatus = "unknown"
)

func PossibleValuesForUserExperienceAnalyticsDevicePerformanceHealthStatus() []string {
	return []string{
		string(UserExperienceAnalyticsDevicePerformanceHealthStatus_InsufficientData),
		string(UserExperienceAnalyticsDevicePerformanceHealthStatus_MeetingGoals),
		string(UserExperienceAnalyticsDevicePerformanceHealthStatus_NeedsAttention),
		string(UserExperienceAnalyticsDevicePerformanceHealthStatus_Unknown),
	}
}

func (s *UserExperienceAnalyticsDevicePerformanceHealthStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserExperienceAnalyticsDevicePerformanceHealthStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserExperienceAnalyticsDevicePerformanceHealthStatus(input string) (*UserExperienceAnalyticsDevicePerformanceHealthStatus, error) {
	vals := map[string]UserExperienceAnalyticsDevicePerformanceHealthStatus{
		"insufficientdata": UserExperienceAnalyticsDevicePerformanceHealthStatus_InsufficientData,
		"meetinggoals":     UserExperienceAnalyticsDevicePerformanceHealthStatus_MeetingGoals,
		"needsattention":   UserExperienceAnalyticsDevicePerformanceHealthStatus_NeedsAttention,
		"unknown":          UserExperienceAnalyticsDevicePerformanceHealthStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsDevicePerformanceHealthStatus(input)
	return &out, nil
}
