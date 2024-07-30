package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType string

const (
	DefaultManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType_Enabled DefaultManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType = "enabled"
	DefaultManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType_None    DefaultManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType = "none"
)

func PossibleValuesForDefaultManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType() []string {
	return []string{
		string(DefaultManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType_Enabled),
		string(DefaultManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType_None),
	}
}

func (s *DefaultManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType(input string) (*DefaultManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType, error) {
	vals := map[string]DefaultManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType{
		"enabled": DefaultManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType_Enabled,
		"none":    DefaultManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType(input)
	return &out, nil
}
