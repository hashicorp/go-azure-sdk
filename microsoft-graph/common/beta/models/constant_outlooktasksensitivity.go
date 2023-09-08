package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutlookTaskSensitivity string

const (
	OutlookTaskSensitivityconfidential OutlookTaskSensitivity = "Confidential"
	OutlookTaskSensitivitynormal       OutlookTaskSensitivity = "Normal"
	OutlookTaskSensitivitypersonal     OutlookTaskSensitivity = "Personal"
	OutlookTaskSensitivityprivate      OutlookTaskSensitivity = "Private"
)

func PossibleValuesForOutlookTaskSensitivity() []string {
	return []string{
		string(OutlookTaskSensitivityconfidential),
		string(OutlookTaskSensitivitynormal),
		string(OutlookTaskSensitivitypersonal),
		string(OutlookTaskSensitivityprivate),
	}
}

func parseOutlookTaskSensitivity(input string) (*OutlookTaskSensitivity, error) {
	vals := map[string]OutlookTaskSensitivity{
		"confidential": OutlookTaskSensitivityconfidential,
		"normal":       OutlookTaskSensitivitynormal,
		"personal":     OutlookTaskSensitivitypersonal,
		"private":      OutlookTaskSensitivityprivate,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OutlookTaskSensitivity(input)
	return &out, nil
}
