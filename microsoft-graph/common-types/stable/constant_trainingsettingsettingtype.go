package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrainingSettingSettingType string

const (
	TrainingSettingSettingType_Custom           TrainingSettingSettingType = "custom"
	TrainingSettingSettingType_MicrosoftCustom  TrainingSettingSettingType = "microsoftCustom"
	TrainingSettingSettingType_MicrosoftManaged TrainingSettingSettingType = "microsoftManaged"
	TrainingSettingSettingType_NoTraining       TrainingSettingSettingType = "noTraining"
)

func PossibleValuesForTrainingSettingSettingType() []string {
	return []string{
		string(TrainingSettingSettingType_Custom),
		string(TrainingSettingSettingType_MicrosoftCustom),
		string(TrainingSettingSettingType_MicrosoftManaged),
		string(TrainingSettingSettingType_NoTraining),
	}
}

func (s *TrainingSettingSettingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTrainingSettingSettingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTrainingSettingSettingType(input string) (*TrainingSettingSettingType, error) {
	vals := map[string]TrainingSettingSettingType{
		"custom":           TrainingSettingSettingType_Custom,
		"microsoftcustom":  TrainingSettingSettingType_MicrosoftCustom,
		"microsoftmanaged": TrainingSettingSettingType_MicrosoftManaged,
		"notraining":       TrainingSettingSettingType_NoTraining,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TrainingSettingSettingType(input)
	return &out, nil
}
