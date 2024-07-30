package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceTimelineEventEventLevel string

const (
	UserExperienceAnalyticsDeviceTimelineEventEventLevel_Critical    UserExperienceAnalyticsDeviceTimelineEventEventLevel = "critical"
	UserExperienceAnalyticsDeviceTimelineEventEventLevel_Error       UserExperienceAnalyticsDeviceTimelineEventEventLevel = "error"
	UserExperienceAnalyticsDeviceTimelineEventEventLevel_Information UserExperienceAnalyticsDeviceTimelineEventEventLevel = "information"
	UserExperienceAnalyticsDeviceTimelineEventEventLevel_None        UserExperienceAnalyticsDeviceTimelineEventEventLevel = "none"
	UserExperienceAnalyticsDeviceTimelineEventEventLevel_Verbose     UserExperienceAnalyticsDeviceTimelineEventEventLevel = "verbose"
	UserExperienceAnalyticsDeviceTimelineEventEventLevel_Warning     UserExperienceAnalyticsDeviceTimelineEventEventLevel = "warning"
)

func PossibleValuesForUserExperienceAnalyticsDeviceTimelineEventEventLevel() []string {
	return []string{
		string(UserExperienceAnalyticsDeviceTimelineEventEventLevel_Critical),
		string(UserExperienceAnalyticsDeviceTimelineEventEventLevel_Error),
		string(UserExperienceAnalyticsDeviceTimelineEventEventLevel_Information),
		string(UserExperienceAnalyticsDeviceTimelineEventEventLevel_None),
		string(UserExperienceAnalyticsDeviceTimelineEventEventLevel_Verbose),
		string(UserExperienceAnalyticsDeviceTimelineEventEventLevel_Warning),
	}
}

func (s *UserExperienceAnalyticsDeviceTimelineEventEventLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserExperienceAnalyticsDeviceTimelineEventEventLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserExperienceAnalyticsDeviceTimelineEventEventLevel(input string) (*UserExperienceAnalyticsDeviceTimelineEventEventLevel, error) {
	vals := map[string]UserExperienceAnalyticsDeviceTimelineEventEventLevel{
		"critical":    UserExperienceAnalyticsDeviceTimelineEventEventLevel_Critical,
		"error":       UserExperienceAnalyticsDeviceTimelineEventEventLevel_Error,
		"information": UserExperienceAnalyticsDeviceTimelineEventEventLevel_Information,
		"none":        UserExperienceAnalyticsDeviceTimelineEventEventLevel_None,
		"verbose":     UserExperienceAnalyticsDeviceTimelineEventEventLevel_Verbose,
		"warning":     UserExperienceAnalyticsDeviceTimelineEventEventLevel_Warning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsDeviceTimelineEventEventLevel(input)
	return &out, nil
}
