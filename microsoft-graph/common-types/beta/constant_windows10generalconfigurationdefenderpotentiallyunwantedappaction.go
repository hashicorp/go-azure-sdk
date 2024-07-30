package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppAction string

const (
	Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppAction_Audit         Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppAction = "audit"
	Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppAction_Block         Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppAction = "block"
	Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppAction_DeviceDefault Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppAction = "deviceDefault"
)

func PossibleValuesForWindows10GeneralConfigurationDefenderPotentiallyUnwantedAppAction() []string {
	return []string{
		string(Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppAction_Audit),
		string(Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppAction_Block),
		string(Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppAction_DeviceDefault),
	}
}

func (s *Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationDefenderPotentiallyUnwantedAppAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationDefenderPotentiallyUnwantedAppAction(input string) (*Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppAction, error) {
	vals := map[string]Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppAction{
		"audit":         Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppAction_Audit,
		"block":         Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppAction_Block,
		"devicedefault": Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppAction_DeviceDefault,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppAction(input)
	return &out, nil
}
