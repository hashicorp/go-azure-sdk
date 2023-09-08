package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftTrainingAssignmentMappingAssignedTo string

const (
	MicrosoftTrainingAssignmentMappingAssignedToallUsers          MicrosoftTrainingAssignmentMappingAssignedTo = "AllUsers"
	MicrosoftTrainingAssignmentMappingAssignedToclickedPayload    MicrosoftTrainingAssignmentMappingAssignedTo = "ClickedPayload"
	MicrosoftTrainingAssignmentMappingAssignedTocompromised       MicrosoftTrainingAssignmentMappingAssignedTo = "Compromised"
	MicrosoftTrainingAssignmentMappingAssignedTodidNothing        MicrosoftTrainingAssignmentMappingAssignedTo = "DidNothing"
	MicrosoftTrainingAssignmentMappingAssignedTonone              MicrosoftTrainingAssignmentMappingAssignedTo = "None"
	MicrosoftTrainingAssignmentMappingAssignedToreadButNotClicked MicrosoftTrainingAssignmentMappingAssignedTo = "ReadButNotClicked"
	MicrosoftTrainingAssignmentMappingAssignedToreportedPhish     MicrosoftTrainingAssignmentMappingAssignedTo = "ReportedPhish"
)

func PossibleValuesForMicrosoftTrainingAssignmentMappingAssignedTo() []string {
	return []string{
		string(MicrosoftTrainingAssignmentMappingAssignedToallUsers),
		string(MicrosoftTrainingAssignmentMappingAssignedToclickedPayload),
		string(MicrosoftTrainingAssignmentMappingAssignedTocompromised),
		string(MicrosoftTrainingAssignmentMappingAssignedTodidNothing),
		string(MicrosoftTrainingAssignmentMappingAssignedTonone),
		string(MicrosoftTrainingAssignmentMappingAssignedToreadButNotClicked),
		string(MicrosoftTrainingAssignmentMappingAssignedToreportedPhish),
	}
}

func parseMicrosoftTrainingAssignmentMappingAssignedTo(input string) (*MicrosoftTrainingAssignmentMappingAssignedTo, error) {
	vals := map[string]MicrosoftTrainingAssignmentMappingAssignedTo{
		"allusers":          MicrosoftTrainingAssignmentMappingAssignedToallUsers,
		"clickedpayload":    MicrosoftTrainingAssignmentMappingAssignedToclickedPayload,
		"compromised":       MicrosoftTrainingAssignmentMappingAssignedTocompromised,
		"didnothing":        MicrosoftTrainingAssignmentMappingAssignedTodidNothing,
		"none":              MicrosoftTrainingAssignmentMappingAssignedTonone,
		"readbutnotclicked": MicrosoftTrainingAssignmentMappingAssignedToreadButNotClicked,
		"reportedphish":     MicrosoftTrainingAssignmentMappingAssignedToreportedPhish,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftTrainingAssignmentMappingAssignedTo(input)
	return &out, nil
}
