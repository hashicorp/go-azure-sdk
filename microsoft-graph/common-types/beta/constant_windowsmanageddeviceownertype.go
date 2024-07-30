package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedDeviceOwnerType string

const (
	WindowsManagedDeviceOwnerType_Company  WindowsManagedDeviceOwnerType = "company"
	WindowsManagedDeviceOwnerType_Personal WindowsManagedDeviceOwnerType = "personal"
	WindowsManagedDeviceOwnerType_Unknown  WindowsManagedDeviceOwnerType = "unknown"
)

func PossibleValuesForWindowsManagedDeviceOwnerType() []string {
	return []string{
		string(WindowsManagedDeviceOwnerType_Company),
		string(WindowsManagedDeviceOwnerType_Personal),
		string(WindowsManagedDeviceOwnerType_Unknown),
	}
}

func (s *WindowsManagedDeviceOwnerType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsManagedDeviceOwnerType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsManagedDeviceOwnerType(input string) (*WindowsManagedDeviceOwnerType, error) {
	vals := map[string]WindowsManagedDeviceOwnerType{
		"company":  WindowsManagedDeviceOwnerType_Company,
		"personal": WindowsManagedDeviceOwnerType_Personal,
		"unknown":  WindowsManagedDeviceOwnerType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedDeviceOwnerType(input)
	return &out, nil
}
