package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType string

const (
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType_AllowPersonalToWork AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType = "allowPersonalToWork"
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType_DeviceDefault       AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType = "deviceDefault"
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType_NoRestrictions      AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType = "noRestrictions"
	AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType_PreventAny          AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType = "preventAny"
)

func PossibleValuesForAndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType() []string {
	return []string{
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType_AllowPersonalToWork),
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType_DeviceDefault),
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType_NoRestrictions),
		string(AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType_PreventAny),
	}
}

func (s *AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType(input string) (*AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType, error) {
	vals := map[string]AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType{
		"allowpersonaltowork": AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType_AllowPersonalToWork,
		"devicedefault":       AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType_DeviceDefault,
		"norestrictions":      AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType_NoRestrictions,
		"preventany":          AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType_PreventAny,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileGeneralDeviceConfigurationWorkProfileDataSharingType(input)
	return &out, nil
}
