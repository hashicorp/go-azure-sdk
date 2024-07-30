package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsBatteryHealthOsPerformanceOsHealthStatus string

const (
	UserExperienceAnalyticsBatteryHealthOsPerformanceOsHealthStatus_InsufficientData UserExperienceAnalyticsBatteryHealthOsPerformanceOsHealthStatus = "insufficientData"
	UserExperienceAnalyticsBatteryHealthOsPerformanceOsHealthStatus_MeetingGoals     UserExperienceAnalyticsBatteryHealthOsPerformanceOsHealthStatus = "meetingGoals"
	UserExperienceAnalyticsBatteryHealthOsPerformanceOsHealthStatus_NeedsAttention   UserExperienceAnalyticsBatteryHealthOsPerformanceOsHealthStatus = "needsAttention"
	UserExperienceAnalyticsBatteryHealthOsPerformanceOsHealthStatus_Unknown          UserExperienceAnalyticsBatteryHealthOsPerformanceOsHealthStatus = "unknown"
)

func PossibleValuesForUserExperienceAnalyticsBatteryHealthOsPerformanceOsHealthStatus() []string {
	return []string{
		string(UserExperienceAnalyticsBatteryHealthOsPerformanceOsHealthStatus_InsufficientData),
		string(UserExperienceAnalyticsBatteryHealthOsPerformanceOsHealthStatus_MeetingGoals),
		string(UserExperienceAnalyticsBatteryHealthOsPerformanceOsHealthStatus_NeedsAttention),
		string(UserExperienceAnalyticsBatteryHealthOsPerformanceOsHealthStatus_Unknown),
	}
}

func (s *UserExperienceAnalyticsBatteryHealthOsPerformanceOsHealthStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserExperienceAnalyticsBatteryHealthOsPerformanceOsHealthStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserExperienceAnalyticsBatteryHealthOsPerformanceOsHealthStatus(input string) (*UserExperienceAnalyticsBatteryHealthOsPerformanceOsHealthStatus, error) {
	vals := map[string]UserExperienceAnalyticsBatteryHealthOsPerformanceOsHealthStatus{
		"insufficientdata": UserExperienceAnalyticsBatteryHealthOsPerformanceOsHealthStatus_InsufficientData,
		"meetinggoals":     UserExperienceAnalyticsBatteryHealthOsPerformanceOsHealthStatus_MeetingGoals,
		"needsattention":   UserExperienceAnalyticsBatteryHealthOsPerformanceOsHealthStatus_NeedsAttention,
		"unknown":          UserExperienceAnalyticsBatteryHealthOsPerformanceOsHealthStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsBatteryHealthOsPerformanceOsHealthStatus(input)
	return &out, nil
}
