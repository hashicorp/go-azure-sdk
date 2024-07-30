package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAppActionIfDeviceLockNotSet string

const (
	AndroidManagedAppProtectionAppActionIfDeviceLockNotSet_Block AndroidManagedAppProtectionAppActionIfDeviceLockNotSet = "block"
	AndroidManagedAppProtectionAppActionIfDeviceLockNotSet_Warn  AndroidManagedAppProtectionAppActionIfDeviceLockNotSet = "warn"
	AndroidManagedAppProtectionAppActionIfDeviceLockNotSet_Wipe  AndroidManagedAppProtectionAppActionIfDeviceLockNotSet = "wipe"
)

func PossibleValuesForAndroidManagedAppProtectionAppActionIfDeviceLockNotSet() []string {
	return []string{
		string(AndroidManagedAppProtectionAppActionIfDeviceLockNotSet_Block),
		string(AndroidManagedAppProtectionAppActionIfDeviceLockNotSet_Warn),
		string(AndroidManagedAppProtectionAppActionIfDeviceLockNotSet_Wipe),
	}
}

func (s *AndroidManagedAppProtectionAppActionIfDeviceLockNotSet) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedAppProtectionAppActionIfDeviceLockNotSet(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedAppProtectionAppActionIfDeviceLockNotSet(input string) (*AndroidManagedAppProtectionAppActionIfDeviceLockNotSet, error) {
	vals := map[string]AndroidManagedAppProtectionAppActionIfDeviceLockNotSet{
		"block": AndroidManagedAppProtectionAppActionIfDeviceLockNotSet_Block,
		"warn":  AndroidManagedAppProtectionAppActionIfDeviceLockNotSet_Warn,
		"wipe":  AndroidManagedAppProtectionAppActionIfDeviceLockNotSet_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAppActionIfDeviceLockNotSet(input)
	return &out, nil
}
