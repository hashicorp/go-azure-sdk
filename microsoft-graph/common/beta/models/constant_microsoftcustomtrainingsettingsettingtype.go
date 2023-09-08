package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftCustomTrainingSettingSettingType string

const (
	MicrosoftCustomTrainingSettingSettingTypecustom           MicrosoftCustomTrainingSettingSettingType = "Custom"
	MicrosoftCustomTrainingSettingSettingTypemicrosoftCustom  MicrosoftCustomTrainingSettingSettingType = "MicrosoftCustom"
	MicrosoftCustomTrainingSettingSettingTypemicrosoftManaged MicrosoftCustomTrainingSettingSettingType = "MicrosoftManaged"
	MicrosoftCustomTrainingSettingSettingTypenoTraining       MicrosoftCustomTrainingSettingSettingType = "NoTraining"
)

func PossibleValuesForMicrosoftCustomTrainingSettingSettingType() []string {
	return []string{
		string(MicrosoftCustomTrainingSettingSettingTypecustom),
		string(MicrosoftCustomTrainingSettingSettingTypemicrosoftCustom),
		string(MicrosoftCustomTrainingSettingSettingTypemicrosoftManaged),
		string(MicrosoftCustomTrainingSettingSettingTypenoTraining),
	}
}

func parseMicrosoftCustomTrainingSettingSettingType(input string) (*MicrosoftCustomTrainingSettingSettingType, error) {
	vals := map[string]MicrosoftCustomTrainingSettingSettingType{
		"custom":           MicrosoftCustomTrainingSettingSettingTypecustom,
		"microsoftcustom":  MicrosoftCustomTrainingSettingSettingTypemicrosoftCustom,
		"microsoftmanaged": MicrosoftCustomTrainingSettingSettingTypemicrosoftManaged,
		"notraining":       MicrosoftCustomTrainingSettingSettingTypenoTraining,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftCustomTrainingSettingSettingType(input)
	return &out, nil
}
