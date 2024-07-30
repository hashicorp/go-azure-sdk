package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowed string

const (
	DefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowed_Block DefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowed = "block"
	DefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowed_Warn  DefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowed = "warn"
	DefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowed_Wipe  DefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowed = "wipe"
)

func PossibleValuesForDefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowed() []string {
	return []string{
		string(DefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowed_Block),
		string(DefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowed_Warn),
		string(DefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowed_Wipe),
	}
}

func (s *DefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowed) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowed(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowed(input string) (*DefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowed, error) {
	vals := map[string]DefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowed{
		"block": DefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowed_Block,
		"warn":  DefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowed_Warn,
		"wipe":  DefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowed_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAppActionIfIosDeviceModelNotAllowed(input)
	return &out, nil
}
