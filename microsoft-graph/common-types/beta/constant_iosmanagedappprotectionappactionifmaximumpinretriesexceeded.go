package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionAppActionIfMaximumPinRetriesExceeded string

const (
	IosManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Block IosManagedAppProtectionAppActionIfMaximumPinRetriesExceeded = "block"
	IosManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Warn  IosManagedAppProtectionAppActionIfMaximumPinRetriesExceeded = "warn"
	IosManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Wipe  IosManagedAppProtectionAppActionIfMaximumPinRetriesExceeded = "wipe"
)

func PossibleValuesForIosManagedAppProtectionAppActionIfMaximumPinRetriesExceeded() []string {
	return []string{
		string(IosManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Block),
		string(IosManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Warn),
		string(IosManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Wipe),
	}
}

func (s *IosManagedAppProtectionAppActionIfMaximumPinRetriesExceeded) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosManagedAppProtectionAppActionIfMaximumPinRetriesExceeded(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosManagedAppProtectionAppActionIfMaximumPinRetriesExceeded(input string) (*IosManagedAppProtectionAppActionIfMaximumPinRetriesExceeded, error) {
	vals := map[string]IosManagedAppProtectionAppActionIfMaximumPinRetriesExceeded{
		"block": IosManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Block,
		"warn":  IosManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Warn,
		"wipe":  IosManagedAppProtectionAppActionIfMaximumPinRetriesExceeded_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionAppActionIfMaximumPinRetriesExceeded(input)
	return &out, nil
}
