package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EditionUpgradeConfigurationWindowsSMode string

const (
	EditionUpgradeConfigurationWindowsSModeblock         EditionUpgradeConfigurationWindowsSMode = "Block"
	EditionUpgradeConfigurationWindowsSModenoRestriction EditionUpgradeConfigurationWindowsSMode = "NoRestriction"
	EditionUpgradeConfigurationWindowsSModeunlock        EditionUpgradeConfigurationWindowsSMode = "Unlock"
)

func PossibleValuesForEditionUpgradeConfigurationWindowsSMode() []string {
	return []string{
		string(EditionUpgradeConfigurationWindowsSModeblock),
		string(EditionUpgradeConfigurationWindowsSModenoRestriction),
		string(EditionUpgradeConfigurationWindowsSModeunlock),
	}
}

func parseEditionUpgradeConfigurationWindowsSMode(input string) (*EditionUpgradeConfigurationWindowsSMode, error) {
	vals := map[string]EditionUpgradeConfigurationWindowsSMode{
		"block":         EditionUpgradeConfigurationWindowsSModeblock,
		"norestriction": EditionUpgradeConfigurationWindowsSModenoRestriction,
		"unlock":        EditionUpgradeConfigurationWindowsSModeunlock,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EditionUpgradeConfigurationWindowsSMode(input)
	return &out, nil
}
