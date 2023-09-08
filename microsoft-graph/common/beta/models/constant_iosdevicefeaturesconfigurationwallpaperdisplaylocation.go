package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosDeviceFeaturesConfigurationWallpaperDisplayLocation string

const (
	IosDeviceFeaturesConfigurationWallpaperDisplayLocationhomeScreen         IosDeviceFeaturesConfigurationWallpaperDisplayLocation = "HomeScreen"
	IosDeviceFeaturesConfigurationWallpaperDisplayLocationlockAndHomeScreens IosDeviceFeaturesConfigurationWallpaperDisplayLocation = "LockAndHomeScreens"
	IosDeviceFeaturesConfigurationWallpaperDisplayLocationlockScreen         IosDeviceFeaturesConfigurationWallpaperDisplayLocation = "LockScreen"
	IosDeviceFeaturesConfigurationWallpaperDisplayLocationnotConfigured      IosDeviceFeaturesConfigurationWallpaperDisplayLocation = "NotConfigured"
)

func PossibleValuesForIosDeviceFeaturesConfigurationWallpaperDisplayLocation() []string {
	return []string{
		string(IosDeviceFeaturesConfigurationWallpaperDisplayLocationhomeScreen),
		string(IosDeviceFeaturesConfigurationWallpaperDisplayLocationlockAndHomeScreens),
		string(IosDeviceFeaturesConfigurationWallpaperDisplayLocationlockScreen),
		string(IosDeviceFeaturesConfigurationWallpaperDisplayLocationnotConfigured),
	}
}

func parseIosDeviceFeaturesConfigurationWallpaperDisplayLocation(input string) (*IosDeviceFeaturesConfigurationWallpaperDisplayLocation, error) {
	vals := map[string]IosDeviceFeaturesConfigurationWallpaperDisplayLocation{
		"homescreen":         IosDeviceFeaturesConfigurationWallpaperDisplayLocationhomeScreen,
		"lockandhomescreens": IosDeviceFeaturesConfigurationWallpaperDisplayLocationlockAndHomeScreens,
		"lockscreen":         IosDeviceFeaturesConfigurationWallpaperDisplayLocationlockScreen,
		"notconfigured":      IosDeviceFeaturesConfigurationWallpaperDisplayLocationnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosDeviceFeaturesConfigurationWallpaperDisplayLocation(input)
	return &out, nil
}
