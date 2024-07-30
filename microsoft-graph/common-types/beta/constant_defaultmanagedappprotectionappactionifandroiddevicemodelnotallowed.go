package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed string

const (
	DefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed_Block DefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed = "block"
	DefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed_Warn  DefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed = "warn"
	DefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed_Wipe  DefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed = "wipe"
)

func PossibleValuesForDefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed() []string {
	return []string{
		string(DefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed_Block),
		string(DefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed_Warn),
		string(DefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed_Wipe),
	}
}

func (s *DefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed(input string) (*DefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed, error) {
	vals := map[string]DefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed{
		"block": DefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed_Block,
		"warn":  DefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed_Warn,
		"wipe":  DefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAppActionIfAndroidDeviceModelNotAllowed(input)
	return &out, nil
}
