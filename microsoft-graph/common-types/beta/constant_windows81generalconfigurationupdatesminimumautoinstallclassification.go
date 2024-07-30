package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassification string

const (
	Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassification_Important               Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassification = "important"
	Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassification_None                    Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassification = "none"
	Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassification_RecommendedAndImportant Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassification = "recommendedAndImportant"
	Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassification_UserDefined             Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassification = "userDefined"
)

func PossibleValuesForWindows81GeneralConfigurationUpdatesMinimumAutoInstallClassification() []string {
	return []string{
		string(Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassification_Important),
		string(Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassification_None),
		string(Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassification_RecommendedAndImportant),
		string(Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassification_UserDefined),
	}
}

func (s *Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassification) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows81GeneralConfigurationUpdatesMinimumAutoInstallClassification(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows81GeneralConfigurationUpdatesMinimumAutoInstallClassification(input string) (*Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassification, error) {
	vals := map[string]Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassification{
		"important":               Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassification_Important,
		"none":                    Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassification_None,
		"recommendedandimportant": Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassification_RecommendedAndImportant,
		"userdefined":             Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassification_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassification(input)
	return &out, nil
}
