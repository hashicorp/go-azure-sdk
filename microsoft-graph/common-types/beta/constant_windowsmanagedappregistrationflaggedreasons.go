package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedAppRegistrationFlaggedReasons string

const (
	WindowsManagedAppRegistrationFlaggedReasons_AndroidBootloaderUnlocked WindowsManagedAppRegistrationFlaggedReasons = "androidBootloaderUnlocked"
	WindowsManagedAppRegistrationFlaggedReasons_AndroidFactoryRomModified WindowsManagedAppRegistrationFlaggedReasons = "androidFactoryRomModified"
	WindowsManagedAppRegistrationFlaggedReasons_None                      WindowsManagedAppRegistrationFlaggedReasons = "none"
	WindowsManagedAppRegistrationFlaggedReasons_RootedDevice              WindowsManagedAppRegistrationFlaggedReasons = "rootedDevice"
)

func PossibleValuesForWindowsManagedAppRegistrationFlaggedReasons() []string {
	return []string{
		string(WindowsManagedAppRegistrationFlaggedReasons_AndroidBootloaderUnlocked),
		string(WindowsManagedAppRegistrationFlaggedReasons_AndroidFactoryRomModified),
		string(WindowsManagedAppRegistrationFlaggedReasons_None),
		string(WindowsManagedAppRegistrationFlaggedReasons_RootedDevice),
	}
}

func (s *WindowsManagedAppRegistrationFlaggedReasons) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsManagedAppRegistrationFlaggedReasons(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsManagedAppRegistrationFlaggedReasons(input string) (*WindowsManagedAppRegistrationFlaggedReasons, error) {
	vals := map[string]WindowsManagedAppRegistrationFlaggedReasons{
		"androidbootloaderunlocked": WindowsManagedAppRegistrationFlaggedReasons_AndroidBootloaderUnlocked,
		"androidfactoryrommodified": WindowsManagedAppRegistrationFlaggedReasons_AndroidFactoryRomModified,
		"none":                      WindowsManagedAppRegistrationFlaggedReasons_None,
		"rooteddevice":              WindowsManagedAppRegistrationFlaggedReasons_RootedDevice,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedAppRegistrationFlaggedReasons(input)
	return &out, nil
}
