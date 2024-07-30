package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedAppleDeviceIdentityResultPlatform string

const (
	ImportedAppleDeviceIdentityResultPlatform_Android       ImportedAppleDeviceIdentityResultPlatform = "android"
	ImportedAppleDeviceIdentityResultPlatform_Ios           ImportedAppleDeviceIdentityResultPlatform = "ios"
	ImportedAppleDeviceIdentityResultPlatform_MacOS         ImportedAppleDeviceIdentityResultPlatform = "macOS"
	ImportedAppleDeviceIdentityResultPlatform_Unknown       ImportedAppleDeviceIdentityResultPlatform = "unknown"
	ImportedAppleDeviceIdentityResultPlatform_Windows       ImportedAppleDeviceIdentityResultPlatform = "windows"
	ImportedAppleDeviceIdentityResultPlatform_WindowsMobile ImportedAppleDeviceIdentityResultPlatform = "windowsMobile"
)

func PossibleValuesForImportedAppleDeviceIdentityResultPlatform() []string {
	return []string{
		string(ImportedAppleDeviceIdentityResultPlatform_Android),
		string(ImportedAppleDeviceIdentityResultPlatform_Ios),
		string(ImportedAppleDeviceIdentityResultPlatform_MacOS),
		string(ImportedAppleDeviceIdentityResultPlatform_Unknown),
		string(ImportedAppleDeviceIdentityResultPlatform_Windows),
		string(ImportedAppleDeviceIdentityResultPlatform_WindowsMobile),
	}
}

func (s *ImportedAppleDeviceIdentityResultPlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseImportedAppleDeviceIdentityResultPlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseImportedAppleDeviceIdentityResultPlatform(input string) (*ImportedAppleDeviceIdentityResultPlatform, error) {
	vals := map[string]ImportedAppleDeviceIdentityResultPlatform{
		"android":       ImportedAppleDeviceIdentityResultPlatform_Android,
		"ios":           ImportedAppleDeviceIdentityResultPlatform_Ios,
		"macos":         ImportedAppleDeviceIdentityResultPlatform_MacOS,
		"unknown":       ImportedAppleDeviceIdentityResultPlatform_Unknown,
		"windows":       ImportedAppleDeviceIdentityResultPlatform_Windows,
		"windowsmobile": ImportedAppleDeviceIdentityResultPlatform_WindowsMobile,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedAppleDeviceIdentityResultPlatform(input)
	return &out, nil
}
