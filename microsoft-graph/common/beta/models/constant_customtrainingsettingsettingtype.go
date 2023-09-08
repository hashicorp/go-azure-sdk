package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomTrainingSettingSettingType string

const (
	CustomTrainingSettingSettingTypecustom           CustomTrainingSettingSettingType = "Custom"
	CustomTrainingSettingSettingTypemicrosoftCustom  CustomTrainingSettingSettingType = "MicrosoftCustom"
	CustomTrainingSettingSettingTypemicrosoftManaged CustomTrainingSettingSettingType = "MicrosoftManaged"
	CustomTrainingSettingSettingTypenoTraining       CustomTrainingSettingSettingType = "NoTraining"
)

func PossibleValuesForCustomTrainingSettingSettingType() []string {
	return []string{
		string(CustomTrainingSettingSettingTypecustom),
		string(CustomTrainingSettingSettingTypemicrosoftCustom),
		string(CustomTrainingSettingSettingTypemicrosoftManaged),
		string(CustomTrainingSettingSettingTypenoTraining),
	}
}

func parseCustomTrainingSettingSettingType(input string) (*CustomTrainingSettingSettingType, error) {
	vals := map[string]CustomTrainingSettingSettingType{
		"custom":           CustomTrainingSettingSettingTypecustom,
		"microsoftcustom":  CustomTrainingSettingSettingTypemicrosoftCustom,
		"microsoftmanaged": CustomTrainingSettingSettingTypemicrosoftManaged,
		"notraining":       CustomTrainingSettingSettingTypenoTraining,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomTrainingSettingSettingType(input)
	return &out, nil
}
