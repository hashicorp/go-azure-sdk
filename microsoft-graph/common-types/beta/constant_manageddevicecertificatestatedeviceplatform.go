package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceCertificateStateDevicePlatform string

const (
	ManagedDeviceCertificateStateDevicePlatform_Android                            ManagedDeviceCertificateStateDevicePlatform = "android"
	ManagedDeviceCertificateStateDevicePlatform_AndroidAOSP                        ManagedDeviceCertificateStateDevicePlatform = "androidAOSP"
	ManagedDeviceCertificateStateDevicePlatform_AndroidForWork                     ManagedDeviceCertificateStateDevicePlatform = "androidForWork"
	ManagedDeviceCertificateStateDevicePlatform_AndroidMobileApplicationManagement ManagedDeviceCertificateStateDevicePlatform = "androidMobileApplicationManagement"
	ManagedDeviceCertificateStateDevicePlatform_AndroidWorkProfile                 ManagedDeviceCertificateStateDevicePlatform = "androidWorkProfile"
	ManagedDeviceCertificateStateDevicePlatform_IOS                                ManagedDeviceCertificateStateDevicePlatform = "iOS"
	ManagedDeviceCertificateStateDevicePlatform_IOSMobileApplicationManagement     ManagedDeviceCertificateStateDevicePlatform = "iOSMobileApplicationManagement"
	ManagedDeviceCertificateStateDevicePlatform_MacOS                              ManagedDeviceCertificateStateDevicePlatform = "macOS"
	ManagedDeviceCertificateStateDevicePlatform_Unknown                            ManagedDeviceCertificateStateDevicePlatform = "unknown"
	ManagedDeviceCertificateStateDevicePlatform_Windows10AndLater                  ManagedDeviceCertificateStateDevicePlatform = "windows10AndLater"
	ManagedDeviceCertificateStateDevicePlatform_Windows81AndLater                  ManagedDeviceCertificateStateDevicePlatform = "windows81AndLater"
	ManagedDeviceCertificateStateDevicePlatform_WindowsPhone81                     ManagedDeviceCertificateStateDevicePlatform = "windowsPhone81"
)

func PossibleValuesForManagedDeviceCertificateStateDevicePlatform() []string {
	return []string{
		string(ManagedDeviceCertificateStateDevicePlatform_Android),
		string(ManagedDeviceCertificateStateDevicePlatform_AndroidAOSP),
		string(ManagedDeviceCertificateStateDevicePlatform_AndroidForWork),
		string(ManagedDeviceCertificateStateDevicePlatform_AndroidMobileApplicationManagement),
		string(ManagedDeviceCertificateStateDevicePlatform_AndroidWorkProfile),
		string(ManagedDeviceCertificateStateDevicePlatform_IOS),
		string(ManagedDeviceCertificateStateDevicePlatform_IOSMobileApplicationManagement),
		string(ManagedDeviceCertificateStateDevicePlatform_MacOS),
		string(ManagedDeviceCertificateStateDevicePlatform_Unknown),
		string(ManagedDeviceCertificateStateDevicePlatform_Windows10AndLater),
		string(ManagedDeviceCertificateStateDevicePlatform_Windows81AndLater),
		string(ManagedDeviceCertificateStateDevicePlatform_WindowsPhone81),
	}
}

func (s *ManagedDeviceCertificateStateDevicePlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceCertificateStateDevicePlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceCertificateStateDevicePlatform(input string) (*ManagedDeviceCertificateStateDevicePlatform, error) {
	vals := map[string]ManagedDeviceCertificateStateDevicePlatform{
		"android":                            ManagedDeviceCertificateStateDevicePlatform_Android,
		"androidaosp":                        ManagedDeviceCertificateStateDevicePlatform_AndroidAOSP,
		"androidforwork":                     ManagedDeviceCertificateStateDevicePlatform_AndroidForWork,
		"androidmobileapplicationmanagement": ManagedDeviceCertificateStateDevicePlatform_AndroidMobileApplicationManagement,
		"androidworkprofile":                 ManagedDeviceCertificateStateDevicePlatform_AndroidWorkProfile,
		"ios":                                ManagedDeviceCertificateStateDevicePlatform_IOS,
		"iosmobileapplicationmanagement":     ManagedDeviceCertificateStateDevicePlatform_IOSMobileApplicationManagement,
		"macos":                              ManagedDeviceCertificateStateDevicePlatform_MacOS,
		"unknown":                            ManagedDeviceCertificateStateDevicePlatform_Unknown,
		"windows10andlater":                  ManagedDeviceCertificateStateDevicePlatform_Windows10AndLater,
		"windows81andlater":                  ManagedDeviceCertificateStateDevicePlatform_Windows81AndLater,
		"windowsphone81":                     ManagedDeviceCertificateStateDevicePlatform_WindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceCertificateStateDevicePlatform(input)
	return &out, nil
}
