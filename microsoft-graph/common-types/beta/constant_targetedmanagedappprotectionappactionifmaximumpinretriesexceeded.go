package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceeded string

const (
	TargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Block TargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceeded = "block"
	TargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Warn  TargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceeded = "warn"
	TargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Wipe  TargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceeded = "wipe"
)

func PossibleValuesForTargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceeded() []string {
	return []string{
		string(TargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Block),
		string(TargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Warn),
		string(TargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Wipe),
	}
}

func (s *TargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceeded) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceeded(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceeded(input string) (*TargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceeded, error) {
	vals := map[string]TargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceeded{
		"block": TargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Block,
		"warn":  TargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Warn,
		"wipe":  TargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppProtectionAppActionIfMaximumPinRetriesExceeded(input)
	return &out, nil
}
