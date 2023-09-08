package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcondarkCircle    AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon = "DarkCircle"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcondarkSquare    AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon = "DarkSquare"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIconlightCircle   AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon = "LightCircle"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIconlightSquare   AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon = "LightSquare"
	AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIconnotConfigured AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon = "NotConfigured"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcondarkCircle),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcondarkSquare),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIconlightCircle),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIconlightSquare),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIconnotConfigured),
	}
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon{
		"darkcircle":    AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcondarkCircle,
		"darksquare":    AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcondarkSquare,
		"lightcircle":   AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIconlightCircle,
		"lightsquare":   AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIconlightSquare,
		"notconfigured": AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIconnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationKioskModeFolderIcon(input)
	return &out, nil
}
