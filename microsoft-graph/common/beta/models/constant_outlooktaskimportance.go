package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutlookTaskImportance string

const (
	OutlookTaskImportancehigh   OutlookTaskImportance = "High"
	OutlookTaskImportancelow    OutlookTaskImportance = "Low"
	OutlookTaskImportancenormal OutlookTaskImportance = "Normal"
)

func PossibleValuesForOutlookTaskImportance() []string {
	return []string{
		string(OutlookTaskImportancehigh),
		string(OutlookTaskImportancelow),
		string(OutlookTaskImportancenormal),
	}
}

func parseOutlookTaskImportance(input string) (*OutlookTaskImportance, error) {
	vals := map[string]OutlookTaskImportance{
		"high":   OutlookTaskImportancehigh,
		"low":    OutlookTaskImportancelow,
		"normal": OutlookTaskImportancenormal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OutlookTaskImportance(input)
	return &out, nil
}
