package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftTrainingAssignmentMappingSettingType string

const (
	MicrosoftTrainingAssignmentMappingSettingTypecustom           MicrosoftTrainingAssignmentMappingSettingType = "Custom"
	MicrosoftTrainingAssignmentMappingSettingTypemicrosoftCustom  MicrosoftTrainingAssignmentMappingSettingType = "MicrosoftCustom"
	MicrosoftTrainingAssignmentMappingSettingTypemicrosoftManaged MicrosoftTrainingAssignmentMappingSettingType = "MicrosoftManaged"
	MicrosoftTrainingAssignmentMappingSettingTypenoTraining       MicrosoftTrainingAssignmentMappingSettingType = "NoTraining"
)

func PossibleValuesForMicrosoftTrainingAssignmentMappingSettingType() []string {
	return []string{
		string(MicrosoftTrainingAssignmentMappingSettingTypecustom),
		string(MicrosoftTrainingAssignmentMappingSettingTypemicrosoftCustom),
		string(MicrosoftTrainingAssignmentMappingSettingTypemicrosoftManaged),
		string(MicrosoftTrainingAssignmentMappingSettingTypenoTraining),
	}
}

func parseMicrosoftTrainingAssignmentMappingSettingType(input string) (*MicrosoftTrainingAssignmentMappingSettingType, error) {
	vals := map[string]MicrosoftTrainingAssignmentMappingSettingType{
		"custom":           MicrosoftTrainingAssignmentMappingSettingTypecustom,
		"microsoftcustom":  MicrosoftTrainingAssignmentMappingSettingTypemicrosoftCustom,
		"microsoftmanaged": MicrosoftTrainingAssignmentMappingSettingTypemicrosoftManaged,
		"notraining":       MicrosoftTrainingAssignmentMappingSettingTypenoTraining,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftTrainingAssignmentMappingSettingType(input)
	return &out, nil
}
