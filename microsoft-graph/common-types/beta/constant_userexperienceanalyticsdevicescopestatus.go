package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceScopeStatus string

const (
	UserExperienceAnalyticsDeviceScopeStatus_Completed        UserExperienceAnalyticsDeviceScopeStatus = "completed"
	UserExperienceAnalyticsDeviceScopeStatus_Computing        UserExperienceAnalyticsDeviceScopeStatus = "computing"
	UserExperienceAnalyticsDeviceScopeStatus_InsufficientData UserExperienceAnalyticsDeviceScopeStatus = "insufficientData"
	UserExperienceAnalyticsDeviceScopeStatus_None             UserExperienceAnalyticsDeviceScopeStatus = "none"
)

func PossibleValuesForUserExperienceAnalyticsDeviceScopeStatus() []string {
	return []string{
		string(UserExperienceAnalyticsDeviceScopeStatus_Completed),
		string(UserExperienceAnalyticsDeviceScopeStatus_Computing),
		string(UserExperienceAnalyticsDeviceScopeStatus_InsufficientData),
		string(UserExperienceAnalyticsDeviceScopeStatus_None),
	}
}

func (s *UserExperienceAnalyticsDeviceScopeStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserExperienceAnalyticsDeviceScopeStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserExperienceAnalyticsDeviceScopeStatus(input string) (*UserExperienceAnalyticsDeviceScopeStatus, error) {
	vals := map[string]UserExperienceAnalyticsDeviceScopeStatus{
		"completed":        UserExperienceAnalyticsDeviceScopeStatus_Completed,
		"computing":        UserExperienceAnalyticsDeviceScopeStatus_Computing,
		"insufficientdata": UserExperienceAnalyticsDeviceScopeStatus_InsufficientData,
		"none":             UserExperienceAnalyticsDeviceScopeStatus_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsDeviceScopeStatus(input)
	return &out, nil
}
