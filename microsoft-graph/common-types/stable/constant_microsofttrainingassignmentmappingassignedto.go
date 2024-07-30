package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftTrainingAssignmentMappingAssignedTo string

const (
	MicrosoftTrainingAssignmentMappingAssignedTo_AllUsers          MicrosoftTrainingAssignmentMappingAssignedTo = "allUsers"
	MicrosoftTrainingAssignmentMappingAssignedTo_ClickedPayload    MicrosoftTrainingAssignmentMappingAssignedTo = "clickedPayload"
	MicrosoftTrainingAssignmentMappingAssignedTo_Compromised       MicrosoftTrainingAssignmentMappingAssignedTo = "compromised"
	MicrosoftTrainingAssignmentMappingAssignedTo_DidNothing        MicrosoftTrainingAssignmentMappingAssignedTo = "didNothing"
	MicrosoftTrainingAssignmentMappingAssignedTo_None              MicrosoftTrainingAssignmentMappingAssignedTo = "none"
	MicrosoftTrainingAssignmentMappingAssignedTo_ReadButNotClicked MicrosoftTrainingAssignmentMappingAssignedTo = "readButNotClicked"
	MicrosoftTrainingAssignmentMappingAssignedTo_ReportedPhish     MicrosoftTrainingAssignmentMappingAssignedTo = "reportedPhish"
)

func PossibleValuesForMicrosoftTrainingAssignmentMappingAssignedTo() []string {
	return []string{
		string(MicrosoftTrainingAssignmentMappingAssignedTo_AllUsers),
		string(MicrosoftTrainingAssignmentMappingAssignedTo_ClickedPayload),
		string(MicrosoftTrainingAssignmentMappingAssignedTo_Compromised),
		string(MicrosoftTrainingAssignmentMappingAssignedTo_DidNothing),
		string(MicrosoftTrainingAssignmentMappingAssignedTo_None),
		string(MicrosoftTrainingAssignmentMappingAssignedTo_ReadButNotClicked),
		string(MicrosoftTrainingAssignmentMappingAssignedTo_ReportedPhish),
	}
}

func (s *MicrosoftTrainingAssignmentMappingAssignedTo) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMicrosoftTrainingAssignmentMappingAssignedTo(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMicrosoftTrainingAssignmentMappingAssignedTo(input string) (*MicrosoftTrainingAssignmentMappingAssignedTo, error) {
	vals := map[string]MicrosoftTrainingAssignmentMappingAssignedTo{
		"allusers":          MicrosoftTrainingAssignmentMappingAssignedTo_AllUsers,
		"clickedpayload":    MicrosoftTrainingAssignmentMappingAssignedTo_ClickedPayload,
		"compromised":       MicrosoftTrainingAssignmentMappingAssignedTo_Compromised,
		"didnothing":        MicrosoftTrainingAssignmentMappingAssignedTo_DidNothing,
		"none":              MicrosoftTrainingAssignmentMappingAssignedTo_None,
		"readbutnotclicked": MicrosoftTrainingAssignmentMappingAssignedTo_ReadButNotClicked,
		"reportedphish":     MicrosoftTrainingAssignmentMappingAssignedTo_ReportedPhish,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftTrainingAssignmentMappingAssignedTo(input)
	return &out, nil
}
