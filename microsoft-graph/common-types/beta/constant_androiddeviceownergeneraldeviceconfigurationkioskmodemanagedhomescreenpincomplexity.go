package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexity string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexity_Complex       AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexity = "complex"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexity_NotConfigured AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexity = "notConfigured"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexity_Simple        AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexity = "simple"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexity() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexity_Complex),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexity_NotConfigured),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexity_Simple),
	}
}

func (s *AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexity(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexity, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexity{
		"complex":       AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexity_Complex,
		"notconfigured": AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexity_NotConfigured,
		"simple":        AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexity_Simple,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeManagedHomeScreenPinComplexity(input)
	return &out, nil
}
