package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftManagedTrainingSettingSettingType string

const (
	MicrosoftManagedTrainingSettingSettingTypecustom           MicrosoftManagedTrainingSettingSettingType = "Custom"
	MicrosoftManagedTrainingSettingSettingTypemicrosoftCustom  MicrosoftManagedTrainingSettingSettingType = "MicrosoftCustom"
	MicrosoftManagedTrainingSettingSettingTypemicrosoftManaged MicrosoftManagedTrainingSettingSettingType = "MicrosoftManaged"
	MicrosoftManagedTrainingSettingSettingTypenoTraining       MicrosoftManagedTrainingSettingSettingType = "NoTraining"
)

func PossibleValuesForMicrosoftManagedTrainingSettingSettingType() []string {
	return []string{
		string(MicrosoftManagedTrainingSettingSettingTypecustom),
		string(MicrosoftManagedTrainingSettingSettingTypemicrosoftCustom),
		string(MicrosoftManagedTrainingSettingSettingTypemicrosoftManaged),
		string(MicrosoftManagedTrainingSettingSettingTypenoTraining),
	}
}

func parseMicrosoftManagedTrainingSettingSettingType(input string) (*MicrosoftManagedTrainingSettingSettingType, error) {
	vals := map[string]MicrosoftManagedTrainingSettingSettingType{
		"custom":           MicrosoftManagedTrainingSettingSettingTypecustom,
		"microsoftcustom":  MicrosoftManagedTrainingSettingSettingTypemicrosoftCustom,
		"microsoftmanaged": MicrosoftManagedTrainingSettingSettingTypemicrosoftManaged,
		"notraining":       MicrosoftManagedTrainingSettingSettingTypenoTraining,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftManagedTrainingSettingSettingType(input)
	return &out, nil
}
