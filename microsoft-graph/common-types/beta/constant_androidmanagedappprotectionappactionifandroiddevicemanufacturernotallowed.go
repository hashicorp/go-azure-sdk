package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed string

const (
	AndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed_Block AndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed = "block"
	AndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed_Warn  AndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed = "warn"
	AndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed_Wipe  AndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed = "wipe"
)

func PossibleValuesForAndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed() []string {
	return []string{
		string(AndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed_Block),
		string(AndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed_Warn),
		string(AndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed_Wipe),
	}
}

func (s *AndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed(input string) (*AndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed, error) {
	vals := map[string]AndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed{
		"block": AndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed_Block,
		"warn":  AndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed_Warn,
		"wipe":  AndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAppActionIfAndroidDeviceManufacturerNotAllowed(input)
	return &out, nil
}
