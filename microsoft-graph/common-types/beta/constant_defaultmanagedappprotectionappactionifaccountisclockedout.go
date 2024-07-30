package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAppActionIfAccountIsClockedOut string

const (
	DefaultManagedAppProtectionAppActionIfAccountIsClockedOut_Block DefaultManagedAppProtectionAppActionIfAccountIsClockedOut = "block"
	DefaultManagedAppProtectionAppActionIfAccountIsClockedOut_Warn  DefaultManagedAppProtectionAppActionIfAccountIsClockedOut = "warn"
	DefaultManagedAppProtectionAppActionIfAccountIsClockedOut_Wipe  DefaultManagedAppProtectionAppActionIfAccountIsClockedOut = "wipe"
)

func PossibleValuesForDefaultManagedAppProtectionAppActionIfAccountIsClockedOut() []string {
	return []string{
		string(DefaultManagedAppProtectionAppActionIfAccountIsClockedOut_Block),
		string(DefaultManagedAppProtectionAppActionIfAccountIsClockedOut_Warn),
		string(DefaultManagedAppProtectionAppActionIfAccountIsClockedOut_Wipe),
	}
}

func (s *DefaultManagedAppProtectionAppActionIfAccountIsClockedOut) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultManagedAppProtectionAppActionIfAccountIsClockedOut(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultManagedAppProtectionAppActionIfAccountIsClockedOut(input string) (*DefaultManagedAppProtectionAppActionIfAccountIsClockedOut, error) {
	vals := map[string]DefaultManagedAppProtectionAppActionIfAccountIsClockedOut{
		"block": DefaultManagedAppProtectionAppActionIfAccountIsClockedOut_Block,
		"warn":  DefaultManagedAppProtectionAppActionIfAccountIsClockedOut_Warn,
		"wipe":  DefaultManagedAppProtectionAppActionIfAccountIsClockedOut_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAppActionIfAccountIsClockedOut(input)
	return &out, nil
}
