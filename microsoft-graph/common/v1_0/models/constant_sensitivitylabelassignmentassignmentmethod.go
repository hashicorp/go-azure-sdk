package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SensitivityLabelAssignmentAssignmentMethod string

const (
	SensitivityLabelAssignmentAssignmentMethodauto       SensitivityLabelAssignmentAssignmentMethod = "Auto"
	SensitivityLabelAssignmentAssignmentMethodprivileged SensitivityLabelAssignmentAssignmentMethod = "Privileged"
	SensitivityLabelAssignmentAssignmentMethodstandard   SensitivityLabelAssignmentAssignmentMethod = "Standard"
)

func PossibleValuesForSensitivityLabelAssignmentAssignmentMethod() []string {
	return []string{
		string(SensitivityLabelAssignmentAssignmentMethodauto),
		string(SensitivityLabelAssignmentAssignmentMethodprivileged),
		string(SensitivityLabelAssignmentAssignmentMethodstandard),
	}
}

func parseSensitivityLabelAssignmentAssignmentMethod(input string) (*SensitivityLabelAssignmentAssignmentMethod, error) {
	vals := map[string]SensitivityLabelAssignmentAssignmentMethod{
		"auto":       SensitivityLabelAssignmentAssignmentMethodauto,
		"privileged": SensitivityLabelAssignmentAssignmentMethodprivileged,
		"standard":   SensitivityLabelAssignmentAssignmentMethodstandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SensitivityLabelAssignmentAssignmentMethod(input)
	return &out, nil
}
