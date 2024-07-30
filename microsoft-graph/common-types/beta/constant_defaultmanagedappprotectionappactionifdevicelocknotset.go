package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAppActionIfDeviceLockNotSet string

const (
	DefaultManagedAppProtectionAppActionIfDeviceLockNotSet_Block DefaultManagedAppProtectionAppActionIfDeviceLockNotSet = "block"
	DefaultManagedAppProtectionAppActionIfDeviceLockNotSet_Warn  DefaultManagedAppProtectionAppActionIfDeviceLockNotSet = "warn"
	DefaultManagedAppProtectionAppActionIfDeviceLockNotSet_Wipe  DefaultManagedAppProtectionAppActionIfDeviceLockNotSet = "wipe"
)

func PossibleValuesForDefaultManagedAppProtectionAppActionIfDeviceLockNotSet() []string {
	return []string{
		string(DefaultManagedAppProtectionAppActionIfDeviceLockNotSet_Block),
		string(DefaultManagedAppProtectionAppActionIfDeviceLockNotSet_Warn),
		string(DefaultManagedAppProtectionAppActionIfDeviceLockNotSet_Wipe),
	}
}

func (s *DefaultManagedAppProtectionAppActionIfDeviceLockNotSet) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultManagedAppProtectionAppActionIfDeviceLockNotSet(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultManagedAppProtectionAppActionIfDeviceLockNotSet(input string) (*DefaultManagedAppProtectionAppActionIfDeviceLockNotSet, error) {
	vals := map[string]DefaultManagedAppProtectionAppActionIfDeviceLockNotSet{
		"block": DefaultManagedAppProtectionAppActionIfDeviceLockNotSet_Block,
		"warn":  DefaultManagedAppProtectionAppActionIfDeviceLockNotSet_Warn,
		"wipe":  DefaultManagedAppProtectionAppActionIfDeviceLockNotSet_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAppActionIfDeviceLockNotSet(input)
	return &out, nil
}
