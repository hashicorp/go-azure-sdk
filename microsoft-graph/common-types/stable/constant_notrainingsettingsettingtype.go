package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NoTrainingSettingSettingType string

const (
	NoTrainingSettingSettingType_Custom           NoTrainingSettingSettingType = "custom"
	NoTrainingSettingSettingType_MicrosoftCustom  NoTrainingSettingSettingType = "microsoftCustom"
	NoTrainingSettingSettingType_MicrosoftManaged NoTrainingSettingSettingType = "microsoftManaged"
	NoTrainingSettingSettingType_NoTraining       NoTrainingSettingSettingType = "noTraining"
)

func PossibleValuesForNoTrainingSettingSettingType() []string {
	return []string{
		string(NoTrainingSettingSettingType_Custom),
		string(NoTrainingSettingSettingType_MicrosoftCustom),
		string(NoTrainingSettingSettingType_MicrosoftManaged),
		string(NoTrainingSettingSettingType_NoTraining),
	}
}

func (s *NoTrainingSettingSettingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNoTrainingSettingSettingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNoTrainingSettingSettingType(input string) (*NoTrainingSettingSettingType, error) {
	vals := map[string]NoTrainingSettingSettingType{
		"custom":           NoTrainingSettingSettingType_Custom,
		"microsoftcustom":  NoTrainingSettingSettingType_MicrosoftCustom,
		"microsoftmanaged": NoTrainingSettingSettingType_MicrosoftManaged,
		"notraining":       NoTrainingSettingSettingType_NoTraining,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NoTrainingSettingSettingType(input)
	return &out, nil
}
