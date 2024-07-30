package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed string

const (
	AndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed_Block AndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed = "block"
	AndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed_Warn  AndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed = "warn"
	AndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed_Wipe  AndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed = "wipe"
)

func PossibleValuesForAndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed() []string {
	return []string{
		string(AndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed_Block),
		string(AndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed_Warn),
		string(AndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed_Wipe),
	}
}

func (s *AndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed(input string) (*AndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed, error) {
	vals := map[string]AndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed{
		"block": AndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed_Block,
		"warn":  AndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed_Warn,
		"wipe":  AndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAppActionIfAndroidSafetyNetAppsVerificationFailed(input)
	return &out, nil
}
