package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType string

const (
	AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType_AllowPersonalToWork AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType = "allowPersonalToWork"
	AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType_DeviceDefault       AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType = "deviceDefault"
	AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType_NoRestrictions      AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType = "noRestrictions"
	AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType_PreventAny          AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType = "preventAny"
)

func PossibleValuesForAndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType() []string {
	return []string{
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType_AllowPersonalToWork),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType_DeviceDefault),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType_NoRestrictions),
		string(AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType_PreventAny),
	}
}

func (s *AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType(input string) (*AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType, error) {
	vals := map[string]AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType{
		"allowpersonaltowork": AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType_AllowPersonalToWork,
		"devicedefault":       AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType_DeviceDefault,
		"norestrictions":      AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType_NoRestrictions,
		"preventany":          AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType_PreventAny,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkGeneralDeviceConfigurationWorkProfileDataSharingType(input)
	return &out, nil
}
