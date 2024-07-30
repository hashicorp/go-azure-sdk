package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType string

const (
	AndroidForWorkCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType_Basic          AndroidForWorkCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType = "basic"
	AndroidForWorkCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType_HardwareBacked AndroidForWorkCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType = "hardwareBacked"
)

func PossibleValuesForAndroidForWorkCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType() []string {
	return []string{
		string(AndroidForWorkCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType_Basic),
		string(AndroidForWorkCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType_HardwareBacked),
	}
}

func (s *AndroidForWorkCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType(input string) (*AndroidForWorkCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType, error) {
	vals := map[string]AndroidForWorkCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType{
		"basic":          AndroidForWorkCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType_Basic,
		"hardwarebacked": AndroidForWorkCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType_HardwareBacked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkCompliancePolicySecurityRequiredAndroidSafetyNetEvaluationType(input)
	return &out, nil
}
