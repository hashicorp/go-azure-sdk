package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityContentLabelAssignmentMethod string

const (
	SecurityContentLabelAssignmentMethodauto       SecurityContentLabelAssignmentMethod = "Auto"
	SecurityContentLabelAssignmentMethodprivileged SecurityContentLabelAssignmentMethod = "Privileged"
	SecurityContentLabelAssignmentMethodstandard   SecurityContentLabelAssignmentMethod = "Standard"
)

func PossibleValuesForSecurityContentLabelAssignmentMethod() []string {
	return []string{
		string(SecurityContentLabelAssignmentMethodauto),
		string(SecurityContentLabelAssignmentMethodprivileged),
		string(SecurityContentLabelAssignmentMethodstandard),
	}
}

func parseSecurityContentLabelAssignmentMethod(input string) (*SecurityContentLabelAssignmentMethod, error) {
	vals := map[string]SecurityContentLabelAssignmentMethod{
		"auto":       SecurityContentLabelAssignmentMethodauto,
		"privileged": SecurityContentLabelAssignmentMethodprivileged,
		"standard":   SecurityContentLabelAssignmentMethodstandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityContentLabelAssignmentMethod(input)
	return &out, nil
}
