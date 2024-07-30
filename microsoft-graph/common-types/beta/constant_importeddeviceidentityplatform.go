package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedDeviceIdentityPlatform string

const (
	ImportedDeviceIdentityPlatform_Android       ImportedDeviceIdentityPlatform = "android"
	ImportedDeviceIdentityPlatform_Ios           ImportedDeviceIdentityPlatform = "ios"
	ImportedDeviceIdentityPlatform_MacOS         ImportedDeviceIdentityPlatform = "macOS"
	ImportedDeviceIdentityPlatform_Unknown       ImportedDeviceIdentityPlatform = "unknown"
	ImportedDeviceIdentityPlatform_Windows       ImportedDeviceIdentityPlatform = "windows"
	ImportedDeviceIdentityPlatform_WindowsMobile ImportedDeviceIdentityPlatform = "windowsMobile"
)

func PossibleValuesForImportedDeviceIdentityPlatform() []string {
	return []string{
		string(ImportedDeviceIdentityPlatform_Android),
		string(ImportedDeviceIdentityPlatform_Ios),
		string(ImportedDeviceIdentityPlatform_MacOS),
		string(ImportedDeviceIdentityPlatform_Unknown),
		string(ImportedDeviceIdentityPlatform_Windows),
		string(ImportedDeviceIdentityPlatform_WindowsMobile),
	}
}

func (s *ImportedDeviceIdentityPlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseImportedDeviceIdentityPlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseImportedDeviceIdentityPlatform(input string) (*ImportedDeviceIdentityPlatform, error) {
	vals := map[string]ImportedDeviceIdentityPlatform{
		"android":       ImportedDeviceIdentityPlatform_Android,
		"ios":           ImportedDeviceIdentityPlatform_Ios,
		"macos":         ImportedDeviceIdentityPlatform_MacOS,
		"unknown":       ImportedDeviceIdentityPlatform_Unknown,
		"windows":       ImportedDeviceIdentityPlatform_Windows,
		"windowsmobile": ImportedDeviceIdentityPlatform_WindowsMobile,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedDeviceIdentityPlatform(input)
	return &out, nil
}
