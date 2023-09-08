package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InformationProtectionContentLabelAssignmentMethod string

const (
	InformationProtectionContentLabelAssignmentMethodauto       InformationProtectionContentLabelAssignmentMethod = "Auto"
	InformationProtectionContentLabelAssignmentMethodprivileged InformationProtectionContentLabelAssignmentMethod = "Privileged"
	InformationProtectionContentLabelAssignmentMethodstandard   InformationProtectionContentLabelAssignmentMethod = "Standard"
)

func PossibleValuesForInformationProtectionContentLabelAssignmentMethod() []string {
	return []string{
		string(InformationProtectionContentLabelAssignmentMethodauto),
		string(InformationProtectionContentLabelAssignmentMethodprivileged),
		string(InformationProtectionContentLabelAssignmentMethodstandard),
	}
}

func parseInformationProtectionContentLabelAssignmentMethod(input string) (*InformationProtectionContentLabelAssignmentMethod, error) {
	vals := map[string]InformationProtectionContentLabelAssignmentMethod{
		"auto":       InformationProtectionContentLabelAssignmentMethodauto,
		"privileged": InformationProtectionContentLabelAssignmentMethodprivileged,
		"standard":   InformationProtectionContentLabelAssignmentMethodstandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InformationProtectionContentLabelAssignmentMethod(input)
	return &out, nil
}
