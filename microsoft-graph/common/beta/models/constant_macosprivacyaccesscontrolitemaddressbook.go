package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPrivacyAccessControlItemAddressBook string

const (
	MacOSPrivacyAccessControlItemAddressBookdisabled      MacOSPrivacyAccessControlItemAddressBook = "Disabled"
	MacOSPrivacyAccessControlItemAddressBookenabled       MacOSPrivacyAccessControlItemAddressBook = "Enabled"
	MacOSPrivacyAccessControlItemAddressBooknotConfigured MacOSPrivacyAccessControlItemAddressBook = "NotConfigured"
)

func PossibleValuesForMacOSPrivacyAccessControlItemAddressBook() []string {
	return []string{
		string(MacOSPrivacyAccessControlItemAddressBookdisabled),
		string(MacOSPrivacyAccessControlItemAddressBookenabled),
		string(MacOSPrivacyAccessControlItemAddressBooknotConfigured),
	}
}

func parseMacOSPrivacyAccessControlItemAddressBook(input string) (*MacOSPrivacyAccessControlItemAddressBook, error) {
	vals := map[string]MacOSPrivacyAccessControlItemAddressBook{
		"disabled":      MacOSPrivacyAccessControlItemAddressBookdisabled,
		"enabled":       MacOSPrivacyAccessControlItemAddressBookenabled,
		"notconfigured": MacOSPrivacyAccessControlItemAddressBooknotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPrivacyAccessControlItemAddressBook(input)
	return &out, nil
}
