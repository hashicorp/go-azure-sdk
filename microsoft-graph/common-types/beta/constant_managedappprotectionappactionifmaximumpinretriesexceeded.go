package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppProtectionAppActionIfMaximumPinRetriesExceeded string

const (
	ManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Block ManagedAppProtectionAppActionIfMaximumPinRetriesExceeded = "block"
	ManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Warn  ManagedAppProtectionAppActionIfMaximumPinRetriesExceeded = "warn"
	ManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Wipe  ManagedAppProtectionAppActionIfMaximumPinRetriesExceeded = "wipe"
)

func PossibleValuesForManagedAppProtectionAppActionIfMaximumPinRetriesExceeded() []string {
	return []string{
		string(ManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Block),
		string(ManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Warn),
		string(ManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Wipe),
	}
}

func (s *ManagedAppProtectionAppActionIfMaximumPinRetriesExceeded) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppProtectionAppActionIfMaximumPinRetriesExceeded(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppProtectionAppActionIfMaximumPinRetriesExceeded(input string) (*ManagedAppProtectionAppActionIfMaximumPinRetriesExceeded, error) {
	vals := map[string]ManagedAppProtectionAppActionIfMaximumPinRetriesExceeded{
		"block": ManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Block,
		"warn":  ManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Warn,
		"wipe":  ManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppProtectionAppActionIfMaximumPinRetriesExceeded(input)
	return &out, nil
}
