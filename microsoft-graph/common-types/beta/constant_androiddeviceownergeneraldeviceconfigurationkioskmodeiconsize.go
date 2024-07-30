package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize_Large         AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize = "large"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize_Largest       AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize = "largest"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize_NotConfigured AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize = "notConfigured"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize_Regular       AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize = "regular"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize_Small         AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize = "small"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize_Smallest      AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize = "smallest"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize_Large),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize_Largest),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize_NotConfigured),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize_Regular),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize_Small),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize_Smallest),
	}
}

func (s *AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize{
		"large":         AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize_Large,
		"largest":       AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize_Largest,
		"notconfigured": AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize_NotConfigured,
		"regular":       AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize_Regular,
		"small":         AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize_Small,
		"smallest":      AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize_Smallest,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeIconSize(input)
	return &out, nil
}
