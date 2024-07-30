package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlatformCredentialAuthenticationMethodPlatform string

const (
	PlatformCredentialAuthenticationMethodPlatform_Android PlatformCredentialAuthenticationMethodPlatform = "android"
	PlatformCredentialAuthenticationMethodPlatform_IOS     PlatformCredentialAuthenticationMethodPlatform = "iOS"
	PlatformCredentialAuthenticationMethodPlatform_Linux   PlatformCredentialAuthenticationMethodPlatform = "linux"
	PlatformCredentialAuthenticationMethodPlatform_MacOS   PlatformCredentialAuthenticationMethodPlatform = "macOS"
	PlatformCredentialAuthenticationMethodPlatform_Unknown PlatformCredentialAuthenticationMethodPlatform = "unknown"
	PlatformCredentialAuthenticationMethodPlatform_Windows PlatformCredentialAuthenticationMethodPlatform = "windows"
)

func PossibleValuesForPlatformCredentialAuthenticationMethodPlatform() []string {
	return []string{
		string(PlatformCredentialAuthenticationMethodPlatform_Android),
		string(PlatformCredentialAuthenticationMethodPlatform_IOS),
		string(PlatformCredentialAuthenticationMethodPlatform_Linux),
		string(PlatformCredentialAuthenticationMethodPlatform_MacOS),
		string(PlatformCredentialAuthenticationMethodPlatform_Unknown),
		string(PlatformCredentialAuthenticationMethodPlatform_Windows),
	}
}

func (s *PlatformCredentialAuthenticationMethodPlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlatformCredentialAuthenticationMethodPlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlatformCredentialAuthenticationMethodPlatform(input string) (*PlatformCredentialAuthenticationMethodPlatform, error) {
	vals := map[string]PlatformCredentialAuthenticationMethodPlatform{
		"android": PlatformCredentialAuthenticationMethodPlatform_Android,
		"ios":     PlatformCredentialAuthenticationMethodPlatform_IOS,
		"linux":   PlatformCredentialAuthenticationMethodPlatform_Linux,
		"macos":   PlatformCredentialAuthenticationMethodPlatform_MacOS,
		"unknown": PlatformCredentialAuthenticationMethodPlatform_Unknown,
		"windows": PlatformCredentialAuthenticationMethodPlatform_Windows,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlatformCredentialAuthenticationMethodPlatform(input)
	return &out, nil
}
