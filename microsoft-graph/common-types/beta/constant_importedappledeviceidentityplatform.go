package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedAppleDeviceIdentityPlatform string

const (
	ImportedAppleDeviceIdentityPlatform_Android       ImportedAppleDeviceIdentityPlatform = "android"
	ImportedAppleDeviceIdentityPlatform_Ios           ImportedAppleDeviceIdentityPlatform = "ios"
	ImportedAppleDeviceIdentityPlatform_MacOS         ImportedAppleDeviceIdentityPlatform = "macOS"
	ImportedAppleDeviceIdentityPlatform_Unknown       ImportedAppleDeviceIdentityPlatform = "unknown"
	ImportedAppleDeviceIdentityPlatform_Windows       ImportedAppleDeviceIdentityPlatform = "windows"
	ImportedAppleDeviceIdentityPlatform_WindowsMobile ImportedAppleDeviceIdentityPlatform = "windowsMobile"
)

func PossibleValuesForImportedAppleDeviceIdentityPlatform() []string {
	return []string{
		string(ImportedAppleDeviceIdentityPlatform_Android),
		string(ImportedAppleDeviceIdentityPlatform_Ios),
		string(ImportedAppleDeviceIdentityPlatform_MacOS),
		string(ImportedAppleDeviceIdentityPlatform_Unknown),
		string(ImportedAppleDeviceIdentityPlatform_Windows),
		string(ImportedAppleDeviceIdentityPlatform_WindowsMobile),
	}
}

func (s *ImportedAppleDeviceIdentityPlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseImportedAppleDeviceIdentityPlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseImportedAppleDeviceIdentityPlatform(input string) (*ImportedAppleDeviceIdentityPlatform, error) {
	vals := map[string]ImportedAppleDeviceIdentityPlatform{
		"android":       ImportedAppleDeviceIdentityPlatform_Android,
		"ios":           ImportedAppleDeviceIdentityPlatform_Ios,
		"macos":         ImportedAppleDeviceIdentityPlatform_MacOS,
		"unknown":       ImportedAppleDeviceIdentityPlatform_Unknown,
		"windows":       ImportedAppleDeviceIdentityPlatform_Windows,
		"windowsmobile": ImportedAppleDeviceIdentityPlatform_WindowsMobile,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedAppleDeviceIdentityPlatform(input)
	return &out, nil
}
