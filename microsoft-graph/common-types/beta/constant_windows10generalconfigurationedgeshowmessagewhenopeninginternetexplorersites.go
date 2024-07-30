package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSites string

const (
	Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSites_Disabled      Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSites = "disabled"
	Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSites_Enabled       Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSites = "enabled"
	Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSites_KeepGoing     Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSites = "keepGoing"
	Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSites_NotConfigured Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSites = "notConfigured"
)

func PossibleValuesForWindows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSites() []string {
	return []string{
		string(Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSites_Disabled),
		string(Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSites_Enabled),
		string(Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSites_KeepGoing),
		string(Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSites_NotConfigured),
	}
}

func (s *Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSites) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSites(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSites(input string) (*Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSites, error) {
	vals := map[string]Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSites{
		"disabled":      Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSites_Disabled,
		"enabled":       Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSites_Enabled,
		"keepgoing":     Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSites_KeepGoing,
		"notconfigured": Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSites_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationEdgeShowMessageWhenOpeningInternetExplorerSites(input)
	return &out, nil
}
