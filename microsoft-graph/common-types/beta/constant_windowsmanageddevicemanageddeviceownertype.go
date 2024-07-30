package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedDeviceManagedDeviceOwnerType string

const (
	WindowsManagedDeviceManagedDeviceOwnerType_Company  WindowsManagedDeviceManagedDeviceOwnerType = "company"
	WindowsManagedDeviceManagedDeviceOwnerType_Personal WindowsManagedDeviceManagedDeviceOwnerType = "personal"
	WindowsManagedDeviceManagedDeviceOwnerType_Unknown  WindowsManagedDeviceManagedDeviceOwnerType = "unknown"
)

func PossibleValuesForWindowsManagedDeviceManagedDeviceOwnerType() []string {
	return []string{
		string(WindowsManagedDeviceManagedDeviceOwnerType_Company),
		string(WindowsManagedDeviceManagedDeviceOwnerType_Personal),
		string(WindowsManagedDeviceManagedDeviceOwnerType_Unknown),
	}
}

func (s *WindowsManagedDeviceManagedDeviceOwnerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsManagedDeviceManagedDeviceOwnerType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsManagedDeviceManagedDeviceOwnerType(input string) (*WindowsManagedDeviceManagedDeviceOwnerType, error) {
	vals := map[string]WindowsManagedDeviceManagedDeviceOwnerType{
		"company":  WindowsManagedDeviceManagedDeviceOwnerType_Company,
		"personal": WindowsManagedDeviceManagedDeviceOwnerType_Personal,
		"unknown":  WindowsManagedDeviceManagedDeviceOwnerType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedDeviceManagedDeviceOwnerType(input)
	return &out, nil
}
