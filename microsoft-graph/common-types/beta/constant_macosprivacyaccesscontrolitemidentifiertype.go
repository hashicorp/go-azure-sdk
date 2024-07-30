package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPrivacyAccessControlItemIdentifierType string

const (
	MacOSPrivacyAccessControlItemIdentifierType_BundleID MacOSPrivacyAccessControlItemIdentifierType = "bundleID"
	MacOSPrivacyAccessControlItemIdentifierType_Path     MacOSPrivacyAccessControlItemIdentifierType = "path"
)

func PossibleValuesForMacOSPrivacyAccessControlItemIdentifierType() []string {
	return []string{
		string(MacOSPrivacyAccessControlItemIdentifierType_BundleID),
		string(MacOSPrivacyAccessControlItemIdentifierType_Path),
	}
}

func (s *MacOSPrivacyAccessControlItemIdentifierType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSPrivacyAccessControlItemIdentifierType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSPrivacyAccessControlItemIdentifierType(input string) (*MacOSPrivacyAccessControlItemIdentifierType, error) {
	vals := map[string]MacOSPrivacyAccessControlItemIdentifierType{
		"bundleid": MacOSPrivacyAccessControlItemIdentifierType_BundleID,
		"path":     MacOSPrivacyAccessControlItemIdentifierType_Path,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPrivacyAccessControlItemIdentifierType(input)
	return &out, nil
}
