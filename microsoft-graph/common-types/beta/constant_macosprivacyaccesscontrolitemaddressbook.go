package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPrivacyAccessControlItemAddressBook string

const (
	MacOSPrivacyAccessControlItemAddressBook_Disabled      MacOSPrivacyAccessControlItemAddressBook = "disabled"
	MacOSPrivacyAccessControlItemAddressBook_Enabled       MacOSPrivacyAccessControlItemAddressBook = "enabled"
	MacOSPrivacyAccessControlItemAddressBook_NotConfigured MacOSPrivacyAccessControlItemAddressBook = "notConfigured"
)

func PossibleValuesForMacOSPrivacyAccessControlItemAddressBook() []string {
	return []string{
		string(MacOSPrivacyAccessControlItemAddressBook_Disabled),
		string(MacOSPrivacyAccessControlItemAddressBook_Enabled),
		string(MacOSPrivacyAccessControlItemAddressBook_NotConfigured),
	}
}

func (s *MacOSPrivacyAccessControlItemAddressBook) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSPrivacyAccessControlItemAddressBook(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSPrivacyAccessControlItemAddressBook(input string) (*MacOSPrivacyAccessControlItemAddressBook, error) {
	vals := map[string]MacOSPrivacyAccessControlItemAddressBook{
		"disabled":      MacOSPrivacyAccessControlItemAddressBook_Disabled,
		"enabled":       MacOSPrivacyAccessControlItemAddressBook_Enabled,
		"notconfigured": MacOSPrivacyAccessControlItemAddressBook_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPrivacyAccessControlItemAddressBook(input)
	return &out, nil
}
