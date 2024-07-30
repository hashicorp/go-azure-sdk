package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomTrainingSettingAssignedTo string

const (
	CustomTrainingSettingAssignedTo_AllUsers          CustomTrainingSettingAssignedTo = "allUsers"
	CustomTrainingSettingAssignedTo_ClickedPayload    CustomTrainingSettingAssignedTo = "clickedPayload"
	CustomTrainingSettingAssignedTo_Compromised       CustomTrainingSettingAssignedTo = "compromised"
	CustomTrainingSettingAssignedTo_DidNothing        CustomTrainingSettingAssignedTo = "didNothing"
	CustomTrainingSettingAssignedTo_None              CustomTrainingSettingAssignedTo = "none"
	CustomTrainingSettingAssignedTo_ReadButNotClicked CustomTrainingSettingAssignedTo = "readButNotClicked"
	CustomTrainingSettingAssignedTo_ReportedPhish     CustomTrainingSettingAssignedTo = "reportedPhish"
)

func PossibleValuesForCustomTrainingSettingAssignedTo() []string {
	return []string{
		string(CustomTrainingSettingAssignedTo_AllUsers),
		string(CustomTrainingSettingAssignedTo_ClickedPayload),
		string(CustomTrainingSettingAssignedTo_Compromised),
		string(CustomTrainingSettingAssignedTo_DidNothing),
		string(CustomTrainingSettingAssignedTo_None),
		string(CustomTrainingSettingAssignedTo_ReadButNotClicked),
		string(CustomTrainingSettingAssignedTo_ReportedPhish),
	}
}

func (s *CustomTrainingSettingAssignedTo) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCustomTrainingSettingAssignedTo(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCustomTrainingSettingAssignedTo(input string) (*CustomTrainingSettingAssignedTo, error) {
	vals := map[string]CustomTrainingSettingAssignedTo{
		"allusers":          CustomTrainingSettingAssignedTo_AllUsers,
		"clickedpayload":    CustomTrainingSettingAssignedTo_ClickedPayload,
		"compromised":       CustomTrainingSettingAssignedTo_Compromised,
		"didnothing":        CustomTrainingSettingAssignedTo_DidNothing,
		"none":              CustomTrainingSettingAssignedTo_None,
		"readbutnotclicked": CustomTrainingSettingAssignedTo_ReadButNotClicked,
		"reportedphish":     CustomTrainingSettingAssignedTo_ReportedPhish,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomTrainingSettingAssignedTo(input)
	return &out, nil
}
