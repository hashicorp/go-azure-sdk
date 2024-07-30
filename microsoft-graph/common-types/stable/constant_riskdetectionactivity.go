package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RiskDetectionActivity string

const (
	RiskDetectionActivity_ServicePrincipal RiskDetectionActivity = "servicePrincipal"
	RiskDetectionActivity_Signin           RiskDetectionActivity = "signin"
	RiskDetectionActivity_User             RiskDetectionActivity = "user"
)

func PossibleValuesForRiskDetectionActivity() []string {
	return []string{
		string(RiskDetectionActivity_ServicePrincipal),
		string(RiskDetectionActivity_Signin),
		string(RiskDetectionActivity_User),
	}
}

func (s *RiskDetectionActivity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRiskDetectionActivity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRiskDetectionActivity(input string) (*RiskDetectionActivity, error) {
	vals := map[string]RiskDetectionActivity{
		"serviceprincipal": RiskDetectionActivity_ServicePrincipal,
		"signin":           RiskDetectionActivity_Signin,
		"user":             RiskDetectionActivity_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RiskDetectionActivity(input)
	return &out, nil
}
