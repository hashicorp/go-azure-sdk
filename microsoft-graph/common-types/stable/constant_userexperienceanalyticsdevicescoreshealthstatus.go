package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceScoresHealthStatus string

const (
	UserExperienceAnalyticsDeviceScoresHealthStatus_InsufficientData UserExperienceAnalyticsDeviceScoresHealthStatus = "insufficientData"
	UserExperienceAnalyticsDeviceScoresHealthStatus_MeetingGoals     UserExperienceAnalyticsDeviceScoresHealthStatus = "meetingGoals"
	UserExperienceAnalyticsDeviceScoresHealthStatus_NeedsAttention   UserExperienceAnalyticsDeviceScoresHealthStatus = "needsAttention"
	UserExperienceAnalyticsDeviceScoresHealthStatus_Unknown          UserExperienceAnalyticsDeviceScoresHealthStatus = "unknown"
)

func PossibleValuesForUserExperienceAnalyticsDeviceScoresHealthStatus() []string {
	return []string{
		string(UserExperienceAnalyticsDeviceScoresHealthStatus_InsufficientData),
		string(UserExperienceAnalyticsDeviceScoresHealthStatus_MeetingGoals),
		string(UserExperienceAnalyticsDeviceScoresHealthStatus_NeedsAttention),
		string(UserExperienceAnalyticsDeviceScoresHealthStatus_Unknown),
	}
}

func (s *UserExperienceAnalyticsDeviceScoresHealthStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserExperienceAnalyticsDeviceScoresHealthStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserExperienceAnalyticsDeviceScoresHealthStatus(input string) (*UserExperienceAnalyticsDeviceScoresHealthStatus, error) {
	vals := map[string]UserExperienceAnalyticsDeviceScoresHealthStatus{
		"insufficientdata": UserExperienceAnalyticsDeviceScoresHealthStatus_InsufficientData,
		"meetinggoals":     UserExperienceAnalyticsDeviceScoresHealthStatus_MeetingGoals,
		"needsattention":   UserExperienceAnalyticsDeviceScoresHealthStatus_NeedsAttention,
		"unknown":          UserExperienceAnalyticsDeviceScoresHealthStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsDeviceScoresHealthStatus(input)
	return &out, nil
}
