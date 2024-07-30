package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionAppActionIfDeviceComplianceRequired string

const (
	IosManagedAppProtectionAppActionIfDeviceComplianceRequired_Block IosManagedAppProtectionAppActionIfDeviceComplianceRequired = "block"
	IosManagedAppProtectionAppActionIfDeviceComplianceRequired_Warn  IosManagedAppProtectionAppActionIfDeviceComplianceRequired = "warn"
	IosManagedAppProtectionAppActionIfDeviceComplianceRequired_Wipe  IosManagedAppProtectionAppActionIfDeviceComplianceRequired = "wipe"
)

func PossibleValuesForIosManagedAppProtectionAppActionIfDeviceComplianceRequired() []string {
	return []string{
		string(IosManagedAppProtectionAppActionIfDeviceComplianceRequired_Block),
		string(IosManagedAppProtectionAppActionIfDeviceComplianceRequired_Warn),
		string(IosManagedAppProtectionAppActionIfDeviceComplianceRequired_Wipe),
	}
}

func (s *IosManagedAppProtectionAppActionIfDeviceComplianceRequired) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosManagedAppProtectionAppActionIfDeviceComplianceRequired(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosManagedAppProtectionAppActionIfDeviceComplianceRequired(input string) (*IosManagedAppProtectionAppActionIfDeviceComplianceRequired, error) {
	vals := map[string]IosManagedAppProtectionAppActionIfDeviceComplianceRequired{
		"block": IosManagedAppProtectionAppActionIfDeviceComplianceRequired_Block,
		"warn":  IosManagedAppProtectionAppActionIfDeviceComplianceRequired_Warn,
		"wipe":  IosManagedAppProtectionAppActionIfDeviceComplianceRequired_Wipe,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionAppActionIfDeviceComplianceRequired(input)
	return &out, nil
}
