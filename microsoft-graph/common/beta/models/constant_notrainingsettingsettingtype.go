package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NoTrainingSettingSettingType string

const (
	NoTrainingSettingSettingTypecustom           NoTrainingSettingSettingType = "Custom"
	NoTrainingSettingSettingTypemicrosoftCustom  NoTrainingSettingSettingType = "MicrosoftCustom"
	NoTrainingSettingSettingTypemicrosoftManaged NoTrainingSettingSettingType = "MicrosoftManaged"
	NoTrainingSettingSettingTypenoTraining       NoTrainingSettingSettingType = "NoTraining"
)

func PossibleValuesForNoTrainingSettingSettingType() []string {
	return []string{
		string(NoTrainingSettingSettingTypecustom),
		string(NoTrainingSettingSettingTypemicrosoftCustom),
		string(NoTrainingSettingSettingTypemicrosoftManaged),
		string(NoTrainingSettingSettingTypenoTraining),
	}
}

func parseNoTrainingSettingSettingType(input string) (*NoTrainingSettingSettingType, error) {
	vals := map[string]NoTrainingSettingSettingType{
		"custom":           NoTrainingSettingSettingTypecustom,
		"microsoftcustom":  NoTrainingSettingSettingTypemicrosoftCustom,
		"microsoftmanaged": NoTrainingSettingSettingTypemicrosoftManaged,
		"notraining":       NoTrainingSettingSettingTypenoTraining,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NoTrainingSettingSettingType(input)
	return &out, nil
}
