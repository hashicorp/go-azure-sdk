package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadCompatibleAssignmentFilterAssignmentFilterManagementType string

const (
	PayloadCompatibleAssignmentFilterAssignmentFilterManagementTypeapps    PayloadCompatibleAssignmentFilterAssignmentFilterManagementType = "Apps"
	PayloadCompatibleAssignmentFilterAssignmentFilterManagementTypedevices PayloadCompatibleAssignmentFilterAssignmentFilterManagementType = "Devices"
)

func PossibleValuesForPayloadCompatibleAssignmentFilterAssignmentFilterManagementType() []string {
	return []string{
		string(PayloadCompatibleAssignmentFilterAssignmentFilterManagementTypeapps),
		string(PayloadCompatibleAssignmentFilterAssignmentFilterManagementTypedevices),
	}
}

func parsePayloadCompatibleAssignmentFilterAssignmentFilterManagementType(input string) (*PayloadCompatibleAssignmentFilterAssignmentFilterManagementType, error) {
	vals := map[string]PayloadCompatibleAssignmentFilterAssignmentFilterManagementType{
		"apps":    PayloadCompatibleAssignmentFilterAssignmentFilterManagementTypeapps,
		"devices": PayloadCompatibleAssignmentFilterAssignmentFilterManagementTypedevices,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PayloadCompatibleAssignmentFilterAssignmentFilterManagementType(input)
	return &out, nil
}
