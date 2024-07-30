package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed string

const (
	AndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed_Block AndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed = "block"
	AndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed_Warn  AndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed = "warn"
	AndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed_Wipe  AndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed = "wipe"
)

func PossibleValuesForAndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed() []string {
	return []string{
		string(AndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed_Block),
		string(AndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed_Warn),
		string(AndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed_Wipe),
	}
}

func (s *AndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed(input string) (*AndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed, error) {
	vals := map[string]AndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed{
		"block": AndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed_Block,
		"warn":  AndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed_Warn,
		"wipe":  AndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed(input)
	return &out, nil
}
