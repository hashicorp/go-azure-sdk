package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityLabelingOptionsAssignmentMethod string

const (
	SecurityLabelingOptionsAssignmentMethodauto       SecurityLabelingOptionsAssignmentMethod = "Auto"
	SecurityLabelingOptionsAssignmentMethodprivileged SecurityLabelingOptionsAssignmentMethod = "Privileged"
	SecurityLabelingOptionsAssignmentMethodstandard   SecurityLabelingOptionsAssignmentMethod = "Standard"
)

func PossibleValuesForSecurityLabelingOptionsAssignmentMethod() []string {
	return []string{
		string(SecurityLabelingOptionsAssignmentMethodauto),
		string(SecurityLabelingOptionsAssignmentMethodprivileged),
		string(SecurityLabelingOptionsAssignmentMethodstandard),
	}
}

func parseSecurityLabelingOptionsAssignmentMethod(input string) (*SecurityLabelingOptionsAssignmentMethod, error) {
	vals := map[string]SecurityLabelingOptionsAssignmentMethod{
		"auto":       SecurityLabelingOptionsAssignmentMethodauto,
		"privileged": SecurityLabelingOptionsAssignmentMethodprivileged,
		"standard":   SecurityLabelingOptionsAssignmentMethodstandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityLabelingOptionsAssignmentMethod(input)
	return &out, nil
}
