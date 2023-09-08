package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassification string

const (
	Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassificationimportant               Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassification = "Important"
	Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassificationnone                    Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassification = "None"
	Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassificationrecommendedAndImportant Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassification = "RecommendedAndImportant"
	Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassificationuserDefined             Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassification = "UserDefined"
)

func PossibleValuesForWindows81GeneralConfigurationUpdatesMinimumAutoInstallClassification() []string {
	return []string{
		string(Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassificationimportant),
		string(Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassificationnone),
		string(Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassificationrecommendedAndImportant),
		string(Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassificationuserDefined),
	}
}

func parseWindows81GeneralConfigurationUpdatesMinimumAutoInstallClassification(input string) (*Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassification, error) {
	vals := map[string]Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassification{
		"important":               Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassificationimportant,
		"none":                    Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassificationnone,
		"recommendedandimportant": Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassificationrecommendedAndImportant,
		"userdefined":             Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassificationuserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81GeneralConfigurationUpdatesMinimumAutoInstallClassification(input)
	return &out, nil
}
