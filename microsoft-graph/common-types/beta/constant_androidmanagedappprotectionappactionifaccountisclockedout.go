package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAppActionIfAccountIsClockedOut string

const (
	AndroidManagedAppProtectionAppActionIfAccountIsClockedOut_Block AndroidManagedAppProtectionAppActionIfAccountIsClockedOut = "block"
	AndroidManagedAppProtectionAppActionIfAccountIsClockedOut_Warn  AndroidManagedAppProtectionAppActionIfAccountIsClockedOut = "warn"
	AndroidManagedAppProtectionAppActionIfAccountIsClockedOut_Wipe  AndroidManagedAppProtectionAppActionIfAccountIsClockedOut = "wipe"
)

func PossibleValuesForAndroidManagedAppProtectionAppActionIfAccountIsClockedOut() []string {
	return []string{
		string(AndroidManagedAppProtectionAppActionIfAccountIsClockedOut_Block),
		string(AndroidManagedAppProtectionAppActionIfAccountIsClockedOut_Warn),
		string(AndroidManagedAppProtectionAppActionIfAccountIsClockedOut_Wipe),
	}
}

func (s *AndroidManagedAppProtectionAppActionIfAccountIsClockedOut) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedAppProtectionAppActionIfAccountIsClockedOut(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedAppProtectionAppActionIfAccountIsClockedOut(input string) (*AndroidManagedAppProtectionAppActionIfAccountIsClockedOut, error) {
	vals := map[string]AndroidManagedAppProtectionAppActionIfAccountIsClockedOut{
		"block": AndroidManagedAppProtectionAppActionIfAccountIsClockedOut_Block,
		"warn":  AndroidManagedAppProtectionAppActionIfAccountIsClockedOut_Warn,
		"wipe":  AndroidManagedAppProtectionAppActionIfAccountIsClockedOut_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAppActionIfAccountIsClockedOut(input)
	return &out, nil
}
