package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed string

const (
	DefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed_Block DefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed = "block"
	DefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed_Warn  DefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed = "warn"
	DefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed_Wipe  DefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed = "wipe"
)

func PossibleValuesForDefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed() []string {
	return []string{
		string(DefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed_Block),
		string(DefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed_Warn),
		string(DefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed_Wipe),
	}
}

func (s *DefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed(input string) (*DefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed, error) {
	vals := map[string]DefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed{
		"block": DefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed_Block,
		"warn":  DefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed_Warn,
		"wipe":  DefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed(input)
	return &out, nil
}
