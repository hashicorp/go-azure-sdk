package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftCustomTrainingSettingSettingType string

const (
	MicrosoftCustomTrainingSettingSettingType_Custom           MicrosoftCustomTrainingSettingSettingType = "custom"
	MicrosoftCustomTrainingSettingSettingType_MicrosoftCustom  MicrosoftCustomTrainingSettingSettingType = "microsoftCustom"
	MicrosoftCustomTrainingSettingSettingType_MicrosoftManaged MicrosoftCustomTrainingSettingSettingType = "microsoftManaged"
	MicrosoftCustomTrainingSettingSettingType_NoTraining       MicrosoftCustomTrainingSettingSettingType = "noTraining"
)

func PossibleValuesForMicrosoftCustomTrainingSettingSettingType() []string {
	return []string{
		string(MicrosoftCustomTrainingSettingSettingType_Custom),
		string(MicrosoftCustomTrainingSettingSettingType_MicrosoftCustom),
		string(MicrosoftCustomTrainingSettingSettingType_MicrosoftManaged),
		string(MicrosoftCustomTrainingSettingSettingType_NoTraining),
	}
}

func (s *MicrosoftCustomTrainingSettingSettingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMicrosoftCustomTrainingSettingSettingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMicrosoftCustomTrainingSettingSettingType(input string) (*MicrosoftCustomTrainingSettingSettingType, error) {
	vals := map[string]MicrosoftCustomTrainingSettingSettingType{
		"custom":           MicrosoftCustomTrainingSettingSettingType_Custom,
		"microsoftcustom":  MicrosoftCustomTrainingSettingSettingType_MicrosoftCustom,
		"microsoftmanaged": MicrosoftCustomTrainingSettingSettingType_MicrosoftManaged,
		"notraining":       MicrosoftCustomTrainingSettingSettingType_NoTraining,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftCustomTrainingSettingSettingType(input)
	return &out, nil
}
