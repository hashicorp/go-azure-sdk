package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType string

const (
	AndroidWorkProfileCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType_Basic          AndroidWorkProfileCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType = "basic"
	AndroidWorkProfileCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType_HardwareBacked AndroidWorkProfileCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType = "hardwareBacked"
)

func PossibleValuesForAndroidWorkProfileCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType() []string {
	return []string{
		string(AndroidWorkProfileCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType_Basic),
		string(AndroidWorkProfileCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType_HardwareBacked),
	}
}

func (s *AndroidWorkProfileCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType(input string) (*AndroidWorkProfileCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType, error) {
	vals := map[string]AndroidWorkProfileCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType{
		"basic":          AndroidWorkProfileCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType_Basic,
		"hardwarebacked": AndroidWorkProfileCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType_HardwareBacked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType(input)
	return &out, nil
}
