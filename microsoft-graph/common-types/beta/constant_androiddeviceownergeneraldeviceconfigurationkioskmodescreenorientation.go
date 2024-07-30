package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientation string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientation_AutoRotate    AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientation = "autoRotate"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientation_Landscape     AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientation = "landscape"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientation_NotConfigured AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientation = "notConfigured"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientation_Portrait      AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientation = "portrait"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientation() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientation_AutoRotate),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientation_Landscape),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientation_NotConfigured),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientation_Portrait),
	}
}

func (s *AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientation(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientation, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientation{
		"autorotate":    AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientation_AutoRotate,
		"landscape":     AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientation_Landscape,
		"notconfigured": AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientation_NotConfigured,
		"portrait":      AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientation_Portrait,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeScreenOrientation(input)
	return &out, nil
}
