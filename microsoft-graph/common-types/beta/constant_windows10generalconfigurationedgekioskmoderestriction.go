package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationEdgeKioskModeRestriction string

const (
	Windows10GeneralConfigurationEdgeKioskModeRestriction_DigitalSignage          Windows10GeneralConfigurationEdgeKioskModeRestriction = "digitalSignage"
	Windows10GeneralConfigurationEdgeKioskModeRestriction_NormalMode              Windows10GeneralConfigurationEdgeKioskModeRestriction = "normalMode"
	Windows10GeneralConfigurationEdgeKioskModeRestriction_NotConfigured           Windows10GeneralConfigurationEdgeKioskModeRestriction = "notConfigured"
	Windows10GeneralConfigurationEdgeKioskModeRestriction_PublicBrowsingMultiApp  Windows10GeneralConfigurationEdgeKioskModeRestriction = "publicBrowsingMultiApp"
	Windows10GeneralConfigurationEdgeKioskModeRestriction_PublicBrowsingSingleApp Windows10GeneralConfigurationEdgeKioskModeRestriction = "publicBrowsingSingleApp"
)

func PossibleValuesForWindows10GeneralConfigurationEdgeKioskModeRestriction() []string {
	return []string{
		string(Windows10GeneralConfigurationEdgeKioskModeRestriction_DigitalSignage),
		string(Windows10GeneralConfigurationEdgeKioskModeRestriction_NormalMode),
		string(Windows10GeneralConfigurationEdgeKioskModeRestriction_NotConfigured),
		string(Windows10GeneralConfigurationEdgeKioskModeRestriction_PublicBrowsingMultiApp),
		string(Windows10GeneralConfigurationEdgeKioskModeRestriction_PublicBrowsingSingleApp),
	}
}

func (s *Windows10GeneralConfigurationEdgeKioskModeRestriction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationEdgeKioskModeRestriction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationEdgeKioskModeRestriction(input string) (*Windows10GeneralConfigurationEdgeKioskModeRestriction, error) {
	vals := map[string]Windows10GeneralConfigurationEdgeKioskModeRestriction{
		"digitalsignage":          Windows10GeneralConfigurationEdgeKioskModeRestriction_DigitalSignage,
		"normalmode":              Windows10GeneralConfigurationEdgeKioskModeRestriction_NormalMode,
		"notconfigured":           Windows10GeneralConfigurationEdgeKioskModeRestriction_NotConfigured,
		"publicbrowsingmultiapp":  Windows10GeneralConfigurationEdgeKioskModeRestriction_PublicBrowsingMultiApp,
		"publicbrowsingsingleapp": Windows10GeneralConfigurationEdgeKioskModeRestriction_PublicBrowsingSingleApp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationEdgeKioskModeRestriction(input)
	return &out, nil
}
