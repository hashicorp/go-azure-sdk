package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus string

const (
	UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus_InsufficientData UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus = "insufficientData"
	UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus_MeetingGoals     UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus = "meetingGoals"
	UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus_NeedsAttention   UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus = "needsAttention"
	UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus_Unknown          UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus = "unknown"
)

func PossibleValuesForUserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus() []string {
	return []string{
		string(UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus_InsufficientData),
		string(UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus_MeetingGoals),
		string(UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus_NeedsAttention),
		string(UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus_Unknown),
	}
}

func (s *UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus(input string) (*UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus, error) {
	vals := map[string]UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus{
		"insufficientdata": UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus_InsufficientData,
		"meetinggoals":     UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus_MeetingGoals,
		"needsattention":   UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus_NeedsAttention,
		"unknown":          UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsWorkFromAnywhereModelPerformanceHealthStatus(input)
	return &out, nil
}
