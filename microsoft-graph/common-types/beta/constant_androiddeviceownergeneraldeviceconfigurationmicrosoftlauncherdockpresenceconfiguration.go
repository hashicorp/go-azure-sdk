package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfiguration string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfiguration_Disabled      AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfiguration = "disabled"
	AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfiguration_Hide          AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfiguration = "hide"
	AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfiguration_NotConfigured AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfiguration = "notConfigured"
	AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfiguration_Show          AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfiguration = "show"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfiguration() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfiguration_Disabled),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfiguration_Hide),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfiguration_NotConfigured),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfiguration_Show),
	}
}

func (s *AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfiguration(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfiguration(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfiguration, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfiguration{
		"disabled":      AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfiguration_Disabled,
		"hide":          AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfiguration_Hide,
		"notconfigured": AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfiguration_NotConfigured,
		"show":          AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfiguration_Show,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherDockPresenceConfiguration(input)
	return &out, nil
}
