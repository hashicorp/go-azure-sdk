package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAnomalyCorrelationGroupOverviewCorrelationGroupPrevalence string

const (
	UserExperienceAnalyticsAnomalyCorrelationGroupOverviewCorrelationGroupPrevalence_High   UserExperienceAnalyticsAnomalyCorrelationGroupOverviewCorrelationGroupPrevalence = "high"
	UserExperienceAnalyticsAnomalyCorrelationGroupOverviewCorrelationGroupPrevalence_Low    UserExperienceAnalyticsAnomalyCorrelationGroupOverviewCorrelationGroupPrevalence = "low"
	UserExperienceAnalyticsAnomalyCorrelationGroupOverviewCorrelationGroupPrevalence_Medium UserExperienceAnalyticsAnomalyCorrelationGroupOverviewCorrelationGroupPrevalence = "medium"
)

func PossibleValuesForUserExperienceAnalyticsAnomalyCorrelationGroupOverviewCorrelationGroupPrevalence() []string {
	return []string{
		string(UserExperienceAnalyticsAnomalyCorrelationGroupOverviewCorrelationGroupPrevalence_High),
		string(UserExperienceAnalyticsAnomalyCorrelationGroupOverviewCorrelationGroupPrevalence_Low),
		string(UserExperienceAnalyticsAnomalyCorrelationGroupOverviewCorrelationGroupPrevalence_Medium),
	}
}

func (s *UserExperienceAnalyticsAnomalyCorrelationGroupOverviewCorrelationGroupPrevalence) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserExperienceAnalyticsAnomalyCorrelationGroupOverviewCorrelationGroupPrevalence(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserExperienceAnalyticsAnomalyCorrelationGroupOverviewCorrelationGroupPrevalence(input string) (*UserExperienceAnalyticsAnomalyCorrelationGroupOverviewCorrelationGroupPrevalence, error) {
	vals := map[string]UserExperienceAnalyticsAnomalyCorrelationGroupOverviewCorrelationGroupPrevalence{
		"high":   UserExperienceAnalyticsAnomalyCorrelationGroupOverviewCorrelationGroupPrevalence_High,
		"low":    UserExperienceAnalyticsAnomalyCorrelationGroupOverviewCorrelationGroupPrevalence_Low,
		"medium": UserExperienceAnalyticsAnomalyCorrelationGroupOverviewCorrelationGroupPrevalence_Medium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsAnomalyCorrelationGroupOverviewCorrelationGroupPrevalence(input)
	return &out, nil
}
