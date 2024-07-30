package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomExtensionStageSettingStage string

const (
	CustomExtensionStageSettingStage_AssignmentFourteenDaysBeforeExpiration CustomExtensionStageSettingStage = "assignmentFourteenDaysBeforeExpiration"
	CustomExtensionStageSettingStage_AssignmentOneDayBeforeExpiration       CustomExtensionStageSettingStage = "assignmentOneDayBeforeExpiration"
	CustomExtensionStageSettingStage_AssignmentRequestApproved              CustomExtensionStageSettingStage = "assignmentRequestApproved"
	CustomExtensionStageSettingStage_AssignmentRequestCreated               CustomExtensionStageSettingStage = "assignmentRequestCreated"
	CustomExtensionStageSettingStage_AssignmentRequestGranted               CustomExtensionStageSettingStage = "assignmentRequestGranted"
	CustomExtensionStageSettingStage_AssignmentRequestRemoved               CustomExtensionStageSettingStage = "assignmentRequestRemoved"
)

func PossibleValuesForCustomExtensionStageSettingStage() []string {
	return []string{
		string(CustomExtensionStageSettingStage_AssignmentFourteenDaysBeforeExpiration),
		string(CustomExtensionStageSettingStage_AssignmentOneDayBeforeExpiration),
		string(CustomExtensionStageSettingStage_AssignmentRequestApproved),
		string(CustomExtensionStageSettingStage_AssignmentRequestCreated),
		string(CustomExtensionStageSettingStage_AssignmentRequestGranted),
		string(CustomExtensionStageSettingStage_AssignmentRequestRemoved),
	}
}

func (s *CustomExtensionStageSettingStage) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCustomExtensionStageSettingStage(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCustomExtensionStageSettingStage(input string) (*CustomExtensionStageSettingStage, error) {
	vals := map[string]CustomExtensionStageSettingStage{
		"assignmentfourteendaysbeforeexpiration": CustomExtensionStageSettingStage_AssignmentFourteenDaysBeforeExpiration,
		"assignmentonedaybeforeexpiration":       CustomExtensionStageSettingStage_AssignmentOneDayBeforeExpiration,
		"assignmentrequestapproved":              CustomExtensionStageSettingStage_AssignmentRequestApproved,
		"assignmentrequestcreated":               CustomExtensionStageSettingStage_AssignmentRequestCreated,
		"assignmentrequestgranted":               CustomExtensionStageSettingStage_AssignmentRequestGranted,
		"assignmentrequestremoved":               CustomExtensionStageSettingStage_AssignmentRequestRemoved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomExtensionStageSettingStage(input)
	return &out, nil
}
