package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81GeneralConfigurationMinimumAutoInstallClassification string

const (
	Windows81GeneralConfigurationMinimumAutoInstallClassification_Important               Windows81GeneralConfigurationMinimumAutoInstallClassification = "important"
	Windows81GeneralConfigurationMinimumAutoInstallClassification_None                    Windows81GeneralConfigurationMinimumAutoInstallClassification = "none"
	Windows81GeneralConfigurationMinimumAutoInstallClassification_RecommendedAndImportant Windows81GeneralConfigurationMinimumAutoInstallClassification = "recommendedAndImportant"
	Windows81GeneralConfigurationMinimumAutoInstallClassification_UserDefined             Windows81GeneralConfigurationMinimumAutoInstallClassification = "userDefined"
)

func PossibleValuesForWindows81GeneralConfigurationMinimumAutoInstallClassification() []string {
	return []string{
		string(Windows81GeneralConfigurationMinimumAutoInstallClassification_Important),
		string(Windows81GeneralConfigurationMinimumAutoInstallClassification_None),
		string(Windows81GeneralConfigurationMinimumAutoInstallClassification_RecommendedAndImportant),
		string(Windows81GeneralConfigurationMinimumAutoInstallClassification_UserDefined),
	}
}

func (s *Windows81GeneralConfigurationMinimumAutoInstallClassification) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows81GeneralConfigurationMinimumAutoInstallClassification(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows81GeneralConfigurationMinimumAutoInstallClassification(input string) (*Windows81GeneralConfigurationMinimumAutoInstallClassification, error) {
	vals := map[string]Windows81GeneralConfigurationMinimumAutoInstallClassification{
		"important":               Windows81GeneralConfigurationMinimumAutoInstallClassification_Important,
		"none":                    Windows81GeneralConfigurationMinimumAutoInstallClassification_None,
		"recommendedandimportant": Windows81GeneralConfigurationMinimumAutoInstallClassification_RecommendedAndImportant,
		"userdefined":             Windows81GeneralConfigurationMinimumAutoInstallClassification_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81GeneralConfigurationMinimumAutoInstallClassification(input)
	return &out, nil
}
