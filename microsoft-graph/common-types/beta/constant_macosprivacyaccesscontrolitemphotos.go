package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPrivacyAccessControlItemPhotos string

const (
	MacOSPrivacyAccessControlItemPhotos_Disabled      MacOSPrivacyAccessControlItemPhotos = "disabled"
	MacOSPrivacyAccessControlItemPhotos_Enabled       MacOSPrivacyAccessControlItemPhotos = "enabled"
	MacOSPrivacyAccessControlItemPhotos_NotConfigured MacOSPrivacyAccessControlItemPhotos = "notConfigured"
)

func PossibleValuesForMacOSPrivacyAccessControlItemPhotos() []string {
	return []string{
		string(MacOSPrivacyAccessControlItemPhotos_Disabled),
		string(MacOSPrivacyAccessControlItemPhotos_Enabled),
		string(MacOSPrivacyAccessControlItemPhotos_NotConfigured),
	}
}

func (s *MacOSPrivacyAccessControlItemPhotos) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSPrivacyAccessControlItemPhotos(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSPrivacyAccessControlItemPhotos(input string) (*MacOSPrivacyAccessControlItemPhotos, error) {
	vals := map[string]MacOSPrivacyAccessControlItemPhotos{
		"disabled":      MacOSPrivacyAccessControlItemPhotos_Disabled,
		"enabled":       MacOSPrivacyAccessControlItemPhotos_Enabled,
		"notconfigured": MacOSPrivacyAccessControlItemPhotos_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPrivacyAccessControlItemPhotos(input)
	return &out, nil
}
