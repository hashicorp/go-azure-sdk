package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus string

const (
	UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus_InsufficientData UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus = "insufficientData"
	UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus_MeetingGoals     UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus = "meetingGoals"
	UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus_NeedsAttention   UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus = "needsAttention"
	UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus_Unknown          UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus = "unknown"
)

func PossibleValuesForUserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus() []string {
	return []string{
		string(UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus_InsufficientData),
		string(UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus_MeetingGoals),
		string(UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus_NeedsAttention),
		string(UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus_Unknown),
	}
}

func (s *UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus(input string) (*UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus, error) {
	vals := map[string]UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus{
		"insufficientdata": UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus_InsufficientData,
		"meetinggoals":     UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus_MeetingGoals,
		"needsattention":   UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus_NeedsAttention,
		"unknown":          UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsBatteryHealthModelPerformanceModelHealthStatus(input)
	return &out, nil
}
