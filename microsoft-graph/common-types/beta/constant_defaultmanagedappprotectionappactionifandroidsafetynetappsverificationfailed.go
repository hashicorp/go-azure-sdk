package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed string

const (
	DefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed_Block DefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed = "block"
	DefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed_Warn  DefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed = "warn"
	DefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed_Wipe  DefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed = "wipe"
)

func PossibleValuesForDefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed() []string {
	return []string{
		string(DefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed_Block),
		string(DefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed_Warn),
		string(DefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed_Wipe),
	}
}

func (s *DefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed(input string) (*DefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed, error) {
	vals := map[string]DefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed{
		"block": DefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed_Block,
		"warn":  DefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed_Warn,
		"wipe":  DefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed(input)
	return &out, nil
}
