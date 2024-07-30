package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType string

const (
	AndroidManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType_Enabled AndroidManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType = "enabled"
	AndroidManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType_None    AndroidManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType = "none"
)

func PossibleValuesForAndroidManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType() []string {
	return []string{
		string(AndroidManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType_Enabled),
		string(AndroidManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType_None),
	}
}

func (s *AndroidManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType(input string) (*AndroidManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType, error) {
	vals := map[string]AndroidManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType{
		"enabled": AndroidManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType_Enabled,
		"none":    AndroidManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionRequiredAndroidSafetyNetAppsVerificationType(input)
	return &out, nil
}
