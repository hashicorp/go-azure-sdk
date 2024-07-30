package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed string

const (
	AndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed_Block AndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed = "block"
	AndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed_Warn  AndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed = "warn"
	AndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed_Wipe  AndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed = "wipe"
)

func PossibleValuesForAndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed() []string {
	return []string{
		string(AndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed_Block),
		string(AndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed_Warn),
		string(AndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed_Wipe),
	}
}

func (s *AndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed(input string) (*AndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed, error) {
	vals := map[string]AndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed{
		"block": AndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed_Block,
		"warn":  AndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed_Warn,
		"wipe":  AndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed(input)
	return &out, nil
}
