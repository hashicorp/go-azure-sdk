package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomTrainingSettingAssignedTo string

const (
	CustomTrainingSettingAssignedToallUsers          CustomTrainingSettingAssignedTo = "AllUsers"
	CustomTrainingSettingAssignedToclickedPayload    CustomTrainingSettingAssignedTo = "ClickedPayload"
	CustomTrainingSettingAssignedTocompromised       CustomTrainingSettingAssignedTo = "Compromised"
	CustomTrainingSettingAssignedTodidNothing        CustomTrainingSettingAssignedTo = "DidNothing"
	CustomTrainingSettingAssignedTonone              CustomTrainingSettingAssignedTo = "None"
	CustomTrainingSettingAssignedToreadButNotClicked CustomTrainingSettingAssignedTo = "ReadButNotClicked"
	CustomTrainingSettingAssignedToreportedPhish     CustomTrainingSettingAssignedTo = "ReportedPhish"
)

func PossibleValuesForCustomTrainingSettingAssignedTo() []string {
	return []string{
		string(CustomTrainingSettingAssignedToallUsers),
		string(CustomTrainingSettingAssignedToclickedPayload),
		string(CustomTrainingSettingAssignedTocompromised),
		string(CustomTrainingSettingAssignedTodidNothing),
		string(CustomTrainingSettingAssignedTonone),
		string(CustomTrainingSettingAssignedToreadButNotClicked),
		string(CustomTrainingSettingAssignedToreportedPhish),
	}
}

func parseCustomTrainingSettingAssignedTo(input string) (*CustomTrainingSettingAssignedTo, error) {
	vals := map[string]CustomTrainingSettingAssignedTo{
		"allusers":          CustomTrainingSettingAssignedToallUsers,
		"clickedpayload":    CustomTrainingSettingAssignedToclickedPayload,
		"compromised":       CustomTrainingSettingAssignedTocompromised,
		"didnothing":        CustomTrainingSettingAssignedTodidNothing,
		"none":              CustomTrainingSettingAssignedTonone,
		"readbutnotclicked": CustomTrainingSettingAssignedToreadButNotClicked,
		"reportedphish":     CustomTrainingSettingAssignedToreportedPhish,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomTrainingSettingAssignedTo(input)
	return &out, nil
}
