package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceeded string

const (
	DefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Block DefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceeded = "block"
	DefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Warn  DefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceeded = "warn"
	DefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Wipe  DefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceeded = "wipe"
)

func PossibleValuesForDefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceeded() []string {
	return []string{
		string(DefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Block),
		string(DefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Warn),
		string(DefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Wipe),
	}
}

func (s *DefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceeded) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceeded(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceeded(input string) (*DefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceeded, error) {
	vals := map[string]DefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceeded{
		"block": DefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Block,
		"warn":  DefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Warn,
		"wipe":  DefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAppActionIfMaximumPinRetriesExceeded(input)
	return &out, nil
}
