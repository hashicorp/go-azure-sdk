package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsModelScoresHealthStatus string

const (
	UserExperienceAnalyticsModelScoresHealthStatus_InsufficientData UserExperienceAnalyticsModelScoresHealthStatus = "insufficientData"
	UserExperienceAnalyticsModelScoresHealthStatus_MeetingGoals     UserExperienceAnalyticsModelScoresHealthStatus = "meetingGoals"
	UserExperienceAnalyticsModelScoresHealthStatus_NeedsAttention   UserExperienceAnalyticsModelScoresHealthStatus = "needsAttention"
	UserExperienceAnalyticsModelScoresHealthStatus_Unknown          UserExperienceAnalyticsModelScoresHealthStatus = "unknown"
)

func PossibleValuesForUserExperienceAnalyticsModelScoresHealthStatus() []string {
	return []string{
		string(UserExperienceAnalyticsModelScoresHealthStatus_InsufficientData),
		string(UserExperienceAnalyticsModelScoresHealthStatus_MeetingGoals),
		string(UserExperienceAnalyticsModelScoresHealthStatus_NeedsAttention),
		string(UserExperienceAnalyticsModelScoresHealthStatus_Unknown),
	}
}

func (s *UserExperienceAnalyticsModelScoresHealthStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserExperienceAnalyticsModelScoresHealthStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserExperienceAnalyticsModelScoresHealthStatus(input string) (*UserExperienceAnalyticsModelScoresHealthStatus, error) {
	vals := map[string]UserExperienceAnalyticsModelScoresHealthStatus{
		"insufficientdata": UserExperienceAnalyticsModelScoresHealthStatus_InsufficientData,
		"meetinggoals":     UserExperienceAnalyticsModelScoresHealthStatus_MeetingGoals,
		"needsattention":   UserExperienceAnalyticsModelScoresHealthStatus_NeedsAttention,
		"unknown":          UserExperienceAnalyticsModelScoresHealthStatus_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsModelScoresHealthStatus(input)
	return &out, nil
}
