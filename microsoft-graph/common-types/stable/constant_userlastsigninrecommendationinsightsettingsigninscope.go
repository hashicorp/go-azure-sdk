package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserLastSignInRecommendationInsightSettingSignInScope string

const (
	UserLastSignInRecommendationInsightSettingSignInScope_Application UserLastSignInRecommendationInsightSettingSignInScope = "application"
	UserLastSignInRecommendationInsightSettingSignInScope_Tenant      UserLastSignInRecommendationInsightSettingSignInScope = "tenant"
)

func PossibleValuesForUserLastSignInRecommendationInsightSettingSignInScope() []string {
	return []string{
		string(UserLastSignInRecommendationInsightSettingSignInScope_Application),
		string(UserLastSignInRecommendationInsightSettingSignInScope_Tenant),
	}
}

func (s *UserLastSignInRecommendationInsightSettingSignInScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserLastSignInRecommendationInsightSettingSignInScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserLastSignInRecommendationInsightSettingSignInScope(input string) (*UserLastSignInRecommendationInsightSettingSignInScope, error) {
	vals := map[string]UserLastSignInRecommendationInsightSettingSignInScope{
		"application": UserLastSignInRecommendationInsightSettingSignInScope_Application,
		"tenant":      UserLastSignInRecommendationInsightSettingSignInScope_Tenant,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserLastSignInRecommendationInsightSettingSignInScope(input)
	return &out, nil
}
