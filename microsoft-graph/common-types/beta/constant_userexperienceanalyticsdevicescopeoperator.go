package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsDeviceScopeOperator string

const (
	UserExperienceAnalyticsDeviceScopeOperator_Equals UserExperienceAnalyticsDeviceScopeOperator = "equals"
	UserExperienceAnalyticsDeviceScopeOperator_None   UserExperienceAnalyticsDeviceScopeOperator = "none"
)

func PossibleValuesForUserExperienceAnalyticsDeviceScopeOperator() []string {
	return []string{
		string(UserExperienceAnalyticsDeviceScopeOperator_Equals),
		string(UserExperienceAnalyticsDeviceScopeOperator_None),
	}
}

func (s *UserExperienceAnalyticsDeviceScopeOperator) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserExperienceAnalyticsDeviceScopeOperator(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserExperienceAnalyticsDeviceScopeOperator(input string) (*UserExperienceAnalyticsDeviceScopeOperator, error) {
	vals := map[string]UserExperienceAnalyticsDeviceScopeOperator{
		"equals": UserExperienceAnalyticsDeviceScopeOperator_Equals,
		"none":   UserExperienceAnalyticsDeviceScopeOperator_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserExperienceAnalyticsDeviceScopeOperator(input)
	return &out, nil
}
