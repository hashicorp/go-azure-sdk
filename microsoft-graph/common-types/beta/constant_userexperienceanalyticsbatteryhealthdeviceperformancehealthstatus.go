package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsBatteryHealthDevicePerformanceHealthStatus string

const (
	UserExperienceAnalyticsBatteryHealthDevicePerformanceHealthStatus_InsufficientData UserExperienceAnalyticsBatteryHealthDevicePerformanceHealthStatus = "insufficientData"
	UserExperienceAnalyticsBatteryHealthDevicePerformanceHealthStatus_MeetingGoals     UserExperienceAnalyticsBatteryHealthDevicePerformanceHealthStatus = "meetingGoals"
	UserExperienceAnalyticsBatteryHealthDevicePerformanceHealthStatus_NeedsAttention   UserExperienceAnalyticsBatteryHealthDevicePerformanceHealthStatus = "needsAttention"
	UserExperienceAnalyticsBatteryHealthDevicePerformanceHealthStatus_Unknown          UserExperienceAnalyticsBatteryHealthDevicePerformanceHealthStatus = "unknown"
)

func PossibleValuesForUserExperienceAnalyticsBatteryHealthDevicePerformanceHealthStatus() []string {
	return []string{
		string(UserExperienceAnalyticsBatteryHealthDevicePerformanceHealthStatus_InsufficientData),
		string(UserExperienceAnalyticsBatteryHealthDevicePerformanceHealthStatus_MeetingGoals),
		string(UserExperienceAnalyticsBatteryHealthDevicePerformanceHealthStatus_NeedsAttention),
		string(UserExperienceAnalyticsBatteryHealthDevicePerformanceHealthStatus_Unknown),
	}
}

func (s *UserExperienceAnalyticsBatteryHealthDevicePerformanceHealthStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserExperienceAnalyticsBatteryHealthDevicePerformanceHealthStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserExperienceAnalyticsBatteryHealthDevicePerformanceHealthStatus(input string) (*UserExperienceAnalyticsBatteryHealthDevicePerformanceHealthStatus, error) {
	vals := map[string]UserExperienceAnalyticsBatteryHealthDevicePerformanceHealthStatus{
		"insufficientdata": UserExperienceAnalyticsBatteryHealthDevicePerformanceHealthStatus_InsufficientData,
		"meetinggoals":     UserExperienceAnalyticsBatteryHealthDevicePerformanceHealthStatus_MeetingGoals,
		"needsattention":   UserExperienceAnalyticsBatteryHealthDevicePerformanceHealthStatus_NeedsAttention,
		"unknown":          UserExperienceAnalyticsBatteryHealthDevicePerformanceHealthStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsBatteryHealthDevicePerformanceHealthStatus(input)
	return &out, nil
}
