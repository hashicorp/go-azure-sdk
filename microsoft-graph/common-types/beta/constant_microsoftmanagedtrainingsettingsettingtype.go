package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftManagedTrainingSettingSettingType string

const (
	MicrosoftManagedTrainingSettingSettingType_Custom           MicrosoftManagedTrainingSettingSettingType = "custom"
	MicrosoftManagedTrainingSettingSettingType_MicrosoftCustom  MicrosoftManagedTrainingSettingSettingType = "microsoftCustom"
	MicrosoftManagedTrainingSettingSettingType_MicrosoftManaged MicrosoftManagedTrainingSettingSettingType = "microsoftManaged"
	MicrosoftManagedTrainingSettingSettingType_NoTraining       MicrosoftManagedTrainingSettingSettingType = "noTraining"
)

func PossibleValuesForMicrosoftManagedTrainingSettingSettingType() []string {
	return []string{
		string(MicrosoftManagedTrainingSettingSettingType_Custom),
		string(MicrosoftManagedTrainingSettingSettingType_MicrosoftCustom),
		string(MicrosoftManagedTrainingSettingSettingType_MicrosoftManaged),
		string(MicrosoftManagedTrainingSettingSettingType_NoTraining),
	}
}

func (s *MicrosoftManagedTrainingSettingSettingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMicrosoftManagedTrainingSettingSettingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMicrosoftManagedTrainingSettingSettingType(input string) (*MicrosoftManagedTrainingSettingSettingType, error) {
	vals := map[string]MicrosoftManagedTrainingSettingSettingType{
		"custom":           MicrosoftManagedTrainingSettingSettingType_Custom,
		"microsoftcustom":  MicrosoftManagedTrainingSettingSettingType_MicrosoftCustom,
		"microsoftmanaged": MicrosoftManagedTrainingSettingSettingType_MicrosoftManaged,
		"notraining":       MicrosoftManagedTrainingSettingSettingType_NoTraining,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftManagedTrainingSettingSettingType(input)
	return &out, nil
}
