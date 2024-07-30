package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed string

const (
	DefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed_Block DefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed = "block"
	DefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed_Warn  DefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed = "warn"
	DefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed_Wipe  DefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed = "wipe"
)

func PossibleValuesForDefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed() []string {
	return []string{
		string(DefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed_Block),
		string(DefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed_Warn),
		string(DefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed_Wipe),
	}
}

func (s *DefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed(input string) (*DefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed, error) {
	vals := map[string]DefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed{
		"block": DefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed_Block,
		"warn":  DefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed_Warn,
		"wipe":  DefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAppActionIfAndroidSafetyNetDeviceAttestationFailed(input)
	return &out, nil
}
