package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsAnomalyAnomalyType string

const (
	UserExperienceAnalyticsAnomalyAnomalyType_Application UserExperienceAnalyticsAnomalyAnomalyType = "application"
	UserExperienceAnalyticsAnomalyAnomalyType_Device      UserExperienceAnalyticsAnomalyAnomalyType = "device"
	UserExperienceAnalyticsAnomalyAnomalyType_Driver      UserExperienceAnalyticsAnomalyAnomalyType = "driver"
	UserExperienceAnalyticsAnomalyAnomalyType_Other       UserExperienceAnalyticsAnomalyAnomalyType = "other"
	UserExperienceAnalyticsAnomalyAnomalyType_StopError   UserExperienceAnalyticsAnomalyAnomalyType = "stopError"
)

func PossibleValuesForUserExperienceAnalyticsAnomalyAnomalyType() []string {
	return []string{
		string(UserExperienceAnalyticsAnomalyAnomalyType_Application),
		string(UserExperienceAnalyticsAnomalyAnomalyType_Device),
		string(UserExperienceAnalyticsAnomalyAnomalyType_Driver),
		string(UserExperienceAnalyticsAnomalyAnomalyType_Other),
		string(UserExperienceAnalyticsAnomalyAnomalyType_StopError),
	}
}

func (s *UserExperienceAnalyticsAnomalyAnomalyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserExperienceAnalyticsAnomalyAnomalyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserExperienceAnalyticsAnomalyAnomalyType(input string) (*UserExperienceAnalyticsAnomalyAnomalyType, error) {
	vals := map[string]UserExperienceAnalyticsAnomalyAnomalyType{
		"application": UserExperienceAnalyticsAnomalyAnomalyType_Application,
		"device":      UserExperienceAnalyticsAnomalyAnomalyType_Device,
		"driver":      UserExperienceAnalyticsAnomalyAnomalyType_Driver,
		"other":       UserExperienceAnalyticsAnomalyAnomalyType_Other,
		"stoperror":   UserExperienceAnalyticsAnomalyAnomalyType_StopError,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsAnomalyAnomalyType(input)
	return &out, nil
}
