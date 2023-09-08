package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosVppEBookAssignmentInstallIntent string

const (
	IosVppEBookAssignmentInstallIntentavailable                  IosVppEBookAssignmentInstallIntent = "Available"
	IosVppEBookAssignmentInstallIntentavailableWithoutEnrollment IosVppEBookAssignmentInstallIntent = "AvailableWithoutEnrollment"
	IosVppEBookAssignmentInstallIntentrequired                   IosVppEBookAssignmentInstallIntent = "Required"
	IosVppEBookAssignmentInstallIntentuninstall                  IosVppEBookAssignmentInstallIntent = "Uninstall"
)

func PossibleValuesForIosVppEBookAssignmentInstallIntent() []string {
	return []string{
		string(IosVppEBookAssignmentInstallIntentavailable),
		string(IosVppEBookAssignmentInstallIntentavailableWithoutEnrollment),
		string(IosVppEBookAssignmentInstallIntentrequired),
		string(IosVppEBookAssignmentInstallIntentuninstall),
	}
}

func parseIosVppEBookAssignmentInstallIntent(input string) (*IosVppEBookAssignmentInstallIntent, error) {
	vals := map[string]IosVppEBookAssignmentInstallIntent{
		"available":                  IosVppEBookAssignmentInstallIntentavailable,
		"availablewithoutenrollment": IosVppEBookAssignmentInstallIntentavailableWithoutEnrollment,
		"required":                   IosVppEBookAssignmentInstallIntentrequired,
		"uninstall":                  IosVppEBookAssignmentInstallIntentuninstall,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosVppEBookAssignmentInstallIntent(input)
	return &out, nil
}
