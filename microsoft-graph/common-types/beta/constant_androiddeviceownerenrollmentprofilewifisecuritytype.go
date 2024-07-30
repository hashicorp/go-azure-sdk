package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerEnrollmentProfileWifiSecurityType string

const (
	AndroidDeviceOwnerEnrollmentProfileWifiSecurityType_None AndroidDeviceOwnerEnrollmentProfileWifiSecurityType = "none"
	AndroidDeviceOwnerEnrollmentProfileWifiSecurityType_Wep  AndroidDeviceOwnerEnrollmentProfileWifiSecurityType = "wep"
	AndroidDeviceOwnerEnrollmentProfileWifiSecurityType_Wpa  AndroidDeviceOwnerEnrollmentProfileWifiSecurityType = "wpa"
)

func PossibleValuesForAndroidDeviceOwnerEnrollmentProfileWifiSecurityType() []string {
	return []string{
		string(AndroidDeviceOwnerEnrollmentProfileWifiSecurityType_None),
		string(AndroidDeviceOwnerEnrollmentProfileWifiSecurityType_Wep),
		string(AndroidDeviceOwnerEnrollmentProfileWifiSecurityType_Wpa),
	}
}

func (s *AndroidDeviceOwnerEnrollmentProfileWifiSecurityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerEnrollmentProfileWifiSecurityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerEnrollmentProfileWifiSecurityType(input string) (*AndroidDeviceOwnerEnrollmentProfileWifiSecurityType, error) {
	vals := map[string]AndroidDeviceOwnerEnrollmentProfileWifiSecurityType{
		"none": AndroidDeviceOwnerEnrollmentProfileWifiSecurityType_None,
		"wep":  AndroidDeviceOwnerEnrollmentProfileWifiSecurityType_Wep,
		"wpa":  AndroidDeviceOwnerEnrollmentProfileWifiSecurityType_Wpa,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerEnrollmentProfileWifiSecurityType(input)
	return &out, nil
}
