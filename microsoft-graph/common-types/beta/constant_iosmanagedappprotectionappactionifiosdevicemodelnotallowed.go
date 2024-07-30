package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionAppActionIfIosDeviceModelNotAllowed string

const (
	IosManagedAppProtectionAppActionIfIosDeviceModelNotAllowed_Block IosManagedAppProtectionAppActionIfIosDeviceModelNotAllowed = "block"
	IosManagedAppProtectionAppActionIfIosDeviceModelNotAllowed_Warn  IosManagedAppProtectionAppActionIfIosDeviceModelNotAllowed = "warn"
	IosManagedAppProtectionAppActionIfIosDeviceModelNotAllowed_Wipe  IosManagedAppProtectionAppActionIfIosDeviceModelNotAllowed = "wipe"
)

func PossibleValuesForIosManagedAppProtectionAppActionIfIosDeviceModelNotAllowed() []string {
	return []string{
		string(IosManagedAppProtectionAppActionIfIosDeviceModelNotAllowed_Block),
		string(IosManagedAppProtectionAppActionIfIosDeviceModelNotAllowed_Warn),
		string(IosManagedAppProtectionAppActionIfIosDeviceModelNotAllowed_Wipe),
	}
}

func (s *IosManagedAppProtectionAppActionIfIosDeviceModelNotAllowed) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosManagedAppProtectionAppActionIfIosDeviceModelNotAllowed(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosManagedAppProtectionAppActionIfIosDeviceModelNotAllowed(input string) (*IosManagedAppProtectionAppActionIfIosDeviceModelNotAllowed, error) {
	vals := map[string]IosManagedAppProtectionAppActionIfIosDeviceModelNotAllowed{
		"block": IosManagedAppProtectionAppActionIfIosDeviceModelNotAllowed_Block,
		"warn":  IosManagedAppProtectionAppActionIfIosDeviceModelNotAllowed_Warn,
		"wipe":  IosManagedAppProtectionAppActionIfIosDeviceModelNotAllowed_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionAppActionIfIosDeviceModelNotAllowed(input)
	return &out, nil
}
