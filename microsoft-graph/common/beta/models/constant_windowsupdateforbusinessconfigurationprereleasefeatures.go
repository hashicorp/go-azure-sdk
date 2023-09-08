package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateForBusinessConfigurationPrereleaseFeatures string

const (
	WindowsUpdateForBusinessConfigurationPrereleaseFeaturesnotAllowed                  WindowsUpdateForBusinessConfigurationPrereleaseFeatures = "NotAllowed"
	WindowsUpdateForBusinessConfigurationPrereleaseFeaturessettingsAndExperimentations WindowsUpdateForBusinessConfigurationPrereleaseFeatures = "SettingsAndExperimentations"
	WindowsUpdateForBusinessConfigurationPrereleaseFeaturessettingsOnly                WindowsUpdateForBusinessConfigurationPrereleaseFeatures = "SettingsOnly"
	WindowsUpdateForBusinessConfigurationPrereleaseFeaturesuserDefined                 WindowsUpdateForBusinessConfigurationPrereleaseFeatures = "UserDefined"
)

func PossibleValuesForWindowsUpdateForBusinessConfigurationPrereleaseFeatures() []string {
	return []string{
		string(WindowsUpdateForBusinessConfigurationPrereleaseFeaturesnotAllowed),
		string(WindowsUpdateForBusinessConfigurationPrereleaseFeaturessettingsAndExperimentations),
		string(WindowsUpdateForBusinessConfigurationPrereleaseFeaturessettingsOnly),
		string(WindowsUpdateForBusinessConfigurationPrereleaseFeaturesuserDefined),
	}
}

func parseWindowsUpdateForBusinessConfigurationPrereleaseFeatures(input string) (*WindowsUpdateForBusinessConfigurationPrereleaseFeatures, error) {
	vals := map[string]WindowsUpdateForBusinessConfigurationPrereleaseFeatures{
		"notallowed":                  WindowsUpdateForBusinessConfigurationPrereleaseFeaturesnotAllowed,
		"settingsandexperimentations": WindowsUpdateForBusinessConfigurationPrereleaseFeaturessettingsAndExperimentations,
		"settingsonly":                WindowsUpdateForBusinessConfigurationPrereleaseFeaturessettingsOnly,
		"userdefined":                 WindowsUpdateForBusinessConfigurationPrereleaseFeaturesuserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdateForBusinessConfigurationPrereleaseFeatures(input)
	return &out, nil
}
