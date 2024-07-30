package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftTrainingAssignmentMappingSettingType string

const (
	MicrosoftTrainingAssignmentMappingSettingType_Custom           MicrosoftTrainingAssignmentMappingSettingType = "custom"
	MicrosoftTrainingAssignmentMappingSettingType_MicrosoftCustom  MicrosoftTrainingAssignmentMappingSettingType = "microsoftCustom"
	MicrosoftTrainingAssignmentMappingSettingType_MicrosoftManaged MicrosoftTrainingAssignmentMappingSettingType = "microsoftManaged"
	MicrosoftTrainingAssignmentMappingSettingType_NoTraining       MicrosoftTrainingAssignmentMappingSettingType = "noTraining"
)

func PossibleValuesForMicrosoftTrainingAssignmentMappingSettingType() []string {
	return []string{
		string(MicrosoftTrainingAssignmentMappingSettingType_Custom),
		string(MicrosoftTrainingAssignmentMappingSettingType_MicrosoftCustom),
		string(MicrosoftTrainingAssignmentMappingSettingType_MicrosoftManaged),
		string(MicrosoftTrainingAssignmentMappingSettingType_NoTraining),
	}
}

func (s *MicrosoftTrainingAssignmentMappingSettingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMicrosoftTrainingAssignmentMappingSettingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMicrosoftTrainingAssignmentMappingSettingType(input string) (*MicrosoftTrainingAssignmentMappingSettingType, error) {
	vals := map[string]MicrosoftTrainingAssignmentMappingSettingType{
		"custom":           MicrosoftTrainingAssignmentMappingSettingType_Custom,
		"microsoftcustom":  MicrosoftTrainingAssignmentMappingSettingType_MicrosoftCustom,
		"microsoftmanaged": MicrosoftTrainingAssignmentMappingSettingType_MicrosoftManaged,
		"notraining":       MicrosoftTrainingAssignmentMappingSettingType_NoTraining,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftTrainingAssignmentMappingSettingType(input)
	return &out, nil
}
