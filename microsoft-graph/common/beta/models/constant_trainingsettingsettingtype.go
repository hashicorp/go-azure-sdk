package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrainingSettingSettingType string

const (
	TrainingSettingSettingTypecustom           TrainingSettingSettingType = "Custom"
	TrainingSettingSettingTypemicrosoftCustom  TrainingSettingSettingType = "MicrosoftCustom"
	TrainingSettingSettingTypemicrosoftManaged TrainingSettingSettingType = "MicrosoftManaged"
	TrainingSettingSettingTypenoTraining       TrainingSettingSettingType = "NoTraining"
)

func PossibleValuesForTrainingSettingSettingType() []string {
	return []string{
		string(TrainingSettingSettingTypecustom),
		string(TrainingSettingSettingTypemicrosoftCustom),
		string(TrainingSettingSettingTypemicrosoftManaged),
		string(TrainingSettingSettingTypenoTraining),
	}
}

func parseTrainingSettingSettingType(input string) (*TrainingSettingSettingType, error) {
	vals := map[string]TrainingSettingSettingType{
		"custom":           TrainingSettingSettingTypecustom,
		"microsoftcustom":  TrainingSettingSettingTypemicrosoftCustom,
		"microsoftmanaged": TrainingSettingSettingTypemicrosoftManaged,
		"notraining":       TrainingSettingSettingTypenoTraining,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TrainingSettingSettingType(input)
	return &out, nil
}
