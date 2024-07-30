package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionRequiredAndroidSafetyNetEvaluationType string

const (
	AndroidManagedAppProtectionRequiredAndroidSafetyNetEvaluationType_Basic          AndroidManagedAppProtectionRequiredAndroidSafetyNetEvaluationType = "basic"
	AndroidManagedAppProtectionRequiredAndroidSafetyNetEvaluationType_HardwareBacked AndroidManagedAppProtectionRequiredAndroidSafetyNetEvaluationType = "hardwareBacked"
)

func PossibleValuesForAndroidManagedAppProtectionRequiredAndroidSafetyNetEvaluationType() []string {
	return []string{
		string(AndroidManagedAppProtectionRequiredAndroidSafetyNetEvaluationType_Basic),
		string(AndroidManagedAppProtectionRequiredAndroidSafetyNetEvaluationType_HardwareBacked),
	}
}

func (s *AndroidManagedAppProtectionRequiredAndroidSafetyNetEvaluationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedAppProtectionRequiredAndroidSafetyNetEvaluationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedAppProtectionRequiredAndroidSafetyNetEvaluationType(input string) (*AndroidManagedAppProtectionRequiredAndroidSafetyNetEvaluationType, error) {
	vals := map[string]AndroidManagedAppProtectionRequiredAndroidSafetyNetEvaluationType{
		"basic":          AndroidManagedAppProtectionRequiredAndroidSafetyNetEvaluationType_Basic,
		"hardwarebacked": AndroidManagedAppProtectionRequiredAndroidSafetyNetEvaluationType_HardwareBacked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionRequiredAndroidSafetyNetEvaluationType(input)
	return &out, nil
}
