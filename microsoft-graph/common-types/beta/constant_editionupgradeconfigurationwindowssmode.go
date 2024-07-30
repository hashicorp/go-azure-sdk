package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EditionUpgradeConfigurationWindowsSMode string

const (
	EditionUpgradeConfigurationWindowsSMode_Block         EditionUpgradeConfigurationWindowsSMode = "block"
	EditionUpgradeConfigurationWindowsSMode_NoRestriction EditionUpgradeConfigurationWindowsSMode = "noRestriction"
	EditionUpgradeConfigurationWindowsSMode_Unlock        EditionUpgradeConfigurationWindowsSMode = "unlock"
)

func PossibleValuesForEditionUpgradeConfigurationWindowsSMode() []string {
	return []string{
		string(EditionUpgradeConfigurationWindowsSMode_Block),
		string(EditionUpgradeConfigurationWindowsSMode_NoRestriction),
		string(EditionUpgradeConfigurationWindowsSMode_Unlock),
	}
}

func (s *EditionUpgradeConfigurationWindowsSMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEditionUpgradeConfigurationWindowsSMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEditionUpgradeConfigurationWindowsSMode(input string) (*EditionUpgradeConfigurationWindowsSMode, error) {
	vals := map[string]EditionUpgradeConfigurationWindowsSMode{
		"block":         EditionUpgradeConfigurationWindowsSMode_Block,
		"norestriction": EditionUpgradeConfigurationWindowsSMode_NoRestriction,
		"unlock":        EditionUpgradeConfigurationWindowsSMode_Unlock,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EditionUpgradeConfigurationWindowsSMode(input)
	return &out, nil
}
