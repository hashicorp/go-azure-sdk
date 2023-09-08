package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LabelingOptionsAssignmentMethod string

const (
	LabelingOptionsAssignmentMethodauto       LabelingOptionsAssignmentMethod = "Auto"
	LabelingOptionsAssignmentMethodprivileged LabelingOptionsAssignmentMethod = "Privileged"
	LabelingOptionsAssignmentMethodstandard   LabelingOptionsAssignmentMethod = "Standard"
)

func PossibleValuesForLabelingOptionsAssignmentMethod() []string {
	return []string{
		string(LabelingOptionsAssignmentMethodauto),
		string(LabelingOptionsAssignmentMethodprivileged),
		string(LabelingOptionsAssignmentMethodstandard),
	}
}

func parseLabelingOptionsAssignmentMethod(input string) (*LabelingOptionsAssignmentMethod, error) {
	vals := map[string]LabelingOptionsAssignmentMethod{
		"auto":       LabelingOptionsAssignmentMethodauto,
		"privileged": LabelingOptionsAssignmentMethodprivileged,
		"standard":   LabelingOptionsAssignmentMethodstandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LabelingOptionsAssignmentMethod(input)
	return &out, nil
}
