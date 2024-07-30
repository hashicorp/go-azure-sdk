package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceScopeParameter string

const (
	UserExperienceAnalyticsDeviceScopeParameter_None     UserExperienceAnalyticsDeviceScopeParameter = "none"
	UserExperienceAnalyticsDeviceScopeParameter_ScopeTag UserExperienceAnalyticsDeviceScopeParameter = "scopeTag"
)

func PossibleValuesForUserExperienceAnalyticsDeviceScopeParameter() []string {
	return []string{
		string(UserExperienceAnalyticsDeviceScopeParameter_None),
		string(UserExperienceAnalyticsDeviceScopeParameter_ScopeTag),
	}
}

func (s *UserExperienceAnalyticsDeviceScopeParameter) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserExperienceAnalyticsDeviceScopeParameter(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserExperienceAnalyticsDeviceScopeParameter(input string) (*UserExperienceAnalyticsDeviceScopeParameter, error) {
	vals := map[string]UserExperienceAnalyticsDeviceScopeParameter{
		"none":     UserExperienceAnalyticsDeviceScopeParameter_None,
		"scopetag": UserExperienceAnalyticsDeviceScopeParameter_ScopeTag,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsDeviceScopeParameter(input)
	return &out, nil
}
