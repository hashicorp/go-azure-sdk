package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81GeneralConfigurationMinimumAutoInstallClassification string

const (
	Windows81GeneralConfigurationMinimumAutoInstallClassificationimportant               Windows81GeneralConfigurationMinimumAutoInstallClassification = "Important"
	Windows81GeneralConfigurationMinimumAutoInstallClassificationnone                    Windows81GeneralConfigurationMinimumAutoInstallClassification = "None"
	Windows81GeneralConfigurationMinimumAutoInstallClassificationrecommendedAndImportant Windows81GeneralConfigurationMinimumAutoInstallClassification = "RecommendedAndImportant"
	Windows81GeneralConfigurationMinimumAutoInstallClassificationuserDefined             Windows81GeneralConfigurationMinimumAutoInstallClassification = "UserDefined"
)

func PossibleValuesForWindows81GeneralConfigurationMinimumAutoInstallClassification() []string {
	return []string{
		string(Windows81GeneralConfigurationMinimumAutoInstallClassificationimportant),
		string(Windows81GeneralConfigurationMinimumAutoInstallClassificationnone),
		string(Windows81GeneralConfigurationMinimumAutoInstallClassificationrecommendedAndImportant),
		string(Windows81GeneralConfigurationMinimumAutoInstallClassificationuserDefined),
	}
}

func parseWindows81GeneralConfigurationMinimumAutoInstallClassification(input string) (*Windows81GeneralConfigurationMinimumAutoInstallClassification, error) {
	vals := map[string]Windows81GeneralConfigurationMinimumAutoInstallClassification{
		"important":               Windows81GeneralConfigurationMinimumAutoInstallClassificationimportant,
		"none":                    Windows81GeneralConfigurationMinimumAutoInstallClassificationnone,
		"recommendedandimportant": Windows81GeneralConfigurationMinimumAutoInstallClassificationrecommendedAndImportant,
		"userdefined":             Windows81GeneralConfigurationMinimumAutoInstallClassificationuserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81GeneralConfigurationMinimumAutoInstallClassification(input)
	return &out, nil
}
