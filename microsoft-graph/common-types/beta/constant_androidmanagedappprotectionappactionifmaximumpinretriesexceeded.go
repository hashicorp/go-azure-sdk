package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceeded string

const (
	AndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Block AndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceeded = "block"
	AndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Warn  AndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceeded = "warn"
	AndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Wipe  AndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceeded = "wipe"
)

func PossibleValuesForAndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceeded() []string {
	return []string{
		string(AndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Block),
		string(AndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Warn),
		string(AndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Wipe),
	}
}

func (s *AndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceeded) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceeded(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceeded(input string) (*AndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceeded, error) {
	vals := map[string]AndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceeded{
		"block": AndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Block,
		"warn":  AndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Warn,
		"wipe":  AndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAppActionIfMaximumPinRetriesExceeded(input)
	return &out, nil
}
