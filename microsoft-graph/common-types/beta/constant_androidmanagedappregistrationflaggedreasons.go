package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppRegistrationFlaggedReasons string

const (
	AndroidManagedAppRegistrationFlaggedReasons_AndroidBootloaderUnlocked AndroidManagedAppRegistrationFlaggedReasons = "androidBootloaderUnlocked"
	AndroidManagedAppRegistrationFlaggedReasons_AndroidFactoryRomModified AndroidManagedAppRegistrationFlaggedReasons = "androidFactoryRomModified"
	AndroidManagedAppRegistrationFlaggedReasons_None                      AndroidManagedAppRegistrationFlaggedReasons = "none"
	AndroidManagedAppRegistrationFlaggedReasons_RootedDevice              AndroidManagedAppRegistrationFlaggedReasons = "rootedDevice"
)

func PossibleValuesForAndroidManagedAppRegistrationFlaggedReasons() []string {
	return []string{
		string(AndroidManagedAppRegistrationFlaggedReasons_AndroidBootloaderUnlocked),
		string(AndroidManagedAppRegistrationFlaggedReasons_AndroidFactoryRomModified),
		string(AndroidManagedAppRegistrationFlaggedReasons_None),
		string(AndroidManagedAppRegistrationFlaggedReasons_RootedDevice),
	}
}

func (s *AndroidManagedAppRegistrationFlaggedReasons) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedAppRegistrationFlaggedReasons(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedAppRegistrationFlaggedReasons(input string) (*AndroidManagedAppRegistrationFlaggedReasons, error) {
	vals := map[string]AndroidManagedAppRegistrationFlaggedReasons{
		"androidbootloaderunlocked": AndroidManagedAppRegistrationFlaggedReasons_AndroidBootloaderUnlocked,
		"androidfactoryrommodified": AndroidManagedAppRegistrationFlaggedReasons_AndroidFactoryRomModified,
		"none":                      AndroidManagedAppRegistrationFlaggedReasons_None,
		"rooteddevice":              AndroidManagedAppRegistrationFlaggedReasons_RootedDevice,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppRegistrationFlaggedReasons(input)
	return &out, nil
}
