package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfiguration string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfiguration_Bottom        AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfiguration = "bottom"
	AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfiguration_Hide          AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfiguration = "hide"
	AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfiguration_NotConfigured AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfiguration = "notConfigured"
	AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfiguration_Top           AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfiguration = "top"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfiguration() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfiguration_Bottom),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfiguration_Hide),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfiguration_NotConfigured),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfiguration_Top),
	}
}

func (s *AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfiguration(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfiguration(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfiguration, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfiguration{
		"bottom":        AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfiguration_Bottom,
		"hide":          AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfiguration_Hide,
		"notconfigured": AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfiguration_NotConfigured,
		"top":           AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfiguration_Top,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfiguration(input)
	return &out, nil
}
