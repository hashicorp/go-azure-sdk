package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomTrainingSettingSettingType string

const (
	CustomTrainingSettingSettingType_Custom           CustomTrainingSettingSettingType = "custom"
	CustomTrainingSettingSettingType_MicrosoftCustom  CustomTrainingSettingSettingType = "microsoftCustom"
	CustomTrainingSettingSettingType_MicrosoftManaged CustomTrainingSettingSettingType = "microsoftManaged"
	CustomTrainingSettingSettingType_NoTraining       CustomTrainingSettingSettingType = "noTraining"
)

func PossibleValuesForCustomTrainingSettingSettingType() []string {
	return []string{
		string(CustomTrainingSettingSettingType_Custom),
		string(CustomTrainingSettingSettingType_MicrosoftCustom),
		string(CustomTrainingSettingSettingType_MicrosoftManaged),
		string(CustomTrainingSettingSettingType_NoTraining),
	}
}

func (s *CustomTrainingSettingSettingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCustomTrainingSettingSettingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCustomTrainingSettingSettingType(input string) (*CustomTrainingSettingSettingType, error) {
	vals := map[string]CustomTrainingSettingSettingType{
		"custom":           CustomTrainingSettingSettingType_Custom,
		"microsoftcustom":  CustomTrainingSettingSettingType_MicrosoftCustom,
		"microsoftmanaged": CustomTrainingSettingSettingType_MicrosoftManaged,
		"notraining":       CustomTrainingSettingSettingType_NoTraining,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomTrainingSettingSettingType(input)
	return &out, nil
}
