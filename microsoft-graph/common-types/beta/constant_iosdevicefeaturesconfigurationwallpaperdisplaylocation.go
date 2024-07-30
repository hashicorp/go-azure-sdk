package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosDeviceFeaturesConfigurationWallpaperDisplayLocation string

const (
	IosDeviceFeaturesConfigurationWallpaperDisplayLocation_HomeScreen         IosDeviceFeaturesConfigurationWallpaperDisplayLocation = "homeScreen"
	IosDeviceFeaturesConfigurationWallpaperDisplayLocation_LockAndHomeScreens IosDeviceFeaturesConfigurationWallpaperDisplayLocation = "lockAndHomeScreens"
	IosDeviceFeaturesConfigurationWallpaperDisplayLocation_LockScreen         IosDeviceFeaturesConfigurationWallpaperDisplayLocation = "lockScreen"
	IosDeviceFeaturesConfigurationWallpaperDisplayLocation_NotConfigured      IosDeviceFeaturesConfigurationWallpaperDisplayLocation = "notConfigured"
)

func PossibleValuesForIosDeviceFeaturesConfigurationWallpaperDisplayLocation() []string {
	return []string{
		string(IosDeviceFeaturesConfigurationWallpaperDisplayLocation_HomeScreen),
		string(IosDeviceFeaturesConfigurationWallpaperDisplayLocation_LockAndHomeScreens),
		string(IosDeviceFeaturesConfigurationWallpaperDisplayLocation_LockScreen),
		string(IosDeviceFeaturesConfigurationWallpaperDisplayLocation_NotConfigured),
	}
}

func (s *IosDeviceFeaturesConfigurationWallpaperDisplayLocation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosDeviceFeaturesConfigurationWallpaperDisplayLocation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosDeviceFeaturesConfigurationWallpaperDisplayLocation(input string) (*IosDeviceFeaturesConfigurationWallpaperDisplayLocation, error) {
	vals := map[string]IosDeviceFeaturesConfigurationWallpaperDisplayLocation{
		"homescreen":         IosDeviceFeaturesConfigurationWallpaperDisplayLocation_HomeScreen,
		"lockandhomescreens": IosDeviceFeaturesConfigurationWallpaperDisplayLocation_LockAndHomeScreens,
		"lockscreen":         IosDeviceFeaturesConfigurationWallpaperDisplayLocation_LockScreen,
		"notconfigured":      IosDeviceFeaturesConfigurationWallpaperDisplayLocation_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosDeviceFeaturesConfigurationWallpaperDisplayLocation(input)
	return &out, nil
}
