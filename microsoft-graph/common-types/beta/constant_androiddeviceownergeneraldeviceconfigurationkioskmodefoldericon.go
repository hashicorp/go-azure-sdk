package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon_DarkCircle    AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon = "darkCircle"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon_DarkSquare    AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon = "darkSquare"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon_LightCircle   AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon = "lightCircle"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon_LightSquare   AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon = "lightSquare"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon_NotConfigured AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon = "notConfigured"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon_DarkCircle),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon_DarkSquare),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon_LightCircle),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon_LightSquare),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon_NotConfigured),
	}
}

func (s *AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon{
		"darkcircle":    AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon_DarkCircle,
		"darksquare":    AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon_DarkSquare,
		"lightcircle":   AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon_LightCircle,
		"lightsquare":   AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon_LightSquare,
		"notconfigured": AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon(input)
	return &out, nil
}
