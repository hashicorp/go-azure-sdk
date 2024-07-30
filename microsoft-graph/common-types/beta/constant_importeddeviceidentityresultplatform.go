package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedDeviceIdentityResultPlatform string

const (
	ImportedDeviceIdentityResultPlatform_Android       ImportedDeviceIdentityResultPlatform = "android"
	ImportedDeviceIdentityResultPlatform_Ios           ImportedDeviceIdentityResultPlatform = "ios"
	ImportedDeviceIdentityResultPlatform_MacOS         ImportedDeviceIdentityResultPlatform = "macOS"
	ImportedDeviceIdentityResultPlatform_Unknown       ImportedDeviceIdentityResultPlatform = "unknown"
	ImportedDeviceIdentityResultPlatform_Windows       ImportedDeviceIdentityResultPlatform = "windows"
	ImportedDeviceIdentityResultPlatform_WindowsMobile ImportedDeviceIdentityResultPlatform = "windowsMobile"
)

func PossibleValuesForImportedDeviceIdentityResultPlatform() []string {
	return []string{
		string(ImportedDeviceIdentityResultPlatform_Android),
		string(ImportedDeviceIdentityResultPlatform_Ios),
		string(ImportedDeviceIdentityResultPlatform_MacOS),
		string(ImportedDeviceIdentityResultPlatform_Unknown),
		string(ImportedDeviceIdentityResultPlatform_Windows),
		string(ImportedDeviceIdentityResultPlatform_WindowsMobile),
	}
}

func (s *ImportedDeviceIdentityResultPlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseImportedDeviceIdentityResultPlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseImportedDeviceIdentityResultPlatform(input string) (*ImportedDeviceIdentityResultPlatform, error) {
	vals := map[string]ImportedDeviceIdentityResultPlatform{
		"android":       ImportedDeviceIdentityResultPlatform_Android,
		"ios":           ImportedDeviceIdentityResultPlatform_Ios,
		"macos":         ImportedDeviceIdentityResultPlatform_MacOS,
		"unknown":       ImportedDeviceIdentityResultPlatform_Unknown,
		"windows":       ImportedDeviceIdentityResultPlatform_Windows,
		"windowsmobile": ImportedDeviceIdentityResultPlatform_WindowsMobile,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedDeviceIdentityResultPlatform(input)
	return &out, nil
}
