package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomExtensionHandlerStage string

const (
	CustomExtensionHandlerStageassignmentFourteenDaysBeforeExpiration CustomExtensionHandlerStage = "AssignmentFourteenDaysBeforeExpiration"
	CustomExtensionHandlerStageassignmentOneDayBeforeExpiration       CustomExtensionHandlerStage = "AssignmentOneDayBeforeExpiration"
	CustomExtensionHandlerStageassignmentRequestApproved              CustomExtensionHandlerStage = "AssignmentRequestApproved"
	CustomExtensionHandlerStageassignmentRequestCreated               CustomExtensionHandlerStage = "AssignmentRequestCreated"
	CustomExtensionHandlerStageassignmentRequestGranted               CustomExtensionHandlerStage = "AssignmentRequestGranted"
	CustomExtensionHandlerStageassignmentRequestRemoved               CustomExtensionHandlerStage = "AssignmentRequestRemoved"
)

func PossibleValuesForCustomExtensionHandlerStage() []string {
	return []string{
		string(CustomExtensionHandlerStageassignmentFourteenDaysBeforeExpiration),
		string(CustomExtensionHandlerStageassignmentOneDayBeforeExpiration),
		string(CustomExtensionHandlerStageassignmentRequestApproved),
		string(CustomExtensionHandlerStageassignmentRequestCreated),
		string(CustomExtensionHandlerStageassignmentRequestGranted),
		string(CustomExtensionHandlerStageassignmentRequestRemoved),
	}
}

func parseCustomExtensionHandlerStage(input string) (*CustomExtensionHandlerStage, error) {
	vals := map[string]CustomExtensionHandlerStage{
		"assignmentfourteendaysbeforeexpiration": CustomExtensionHandlerStageassignmentFourteenDaysBeforeExpiration,
		"assignmentonedaybeforeexpiration":       CustomExtensionHandlerStageassignmentOneDayBeforeExpiration,
		"assignmentrequestapproved":              CustomExtensionHandlerStageassignmentRequestApproved,
		"assignmentrequestcreated":               CustomExtensionHandlerStageassignmentRequestCreated,
		"assignmentrequestgranted":               CustomExtensionHandlerStageassignmentRequestGranted,
		"assignmentrequestremoved":               CustomExtensionHandlerStageassignmentRequestRemoved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomExtensionHandlerStage(input)
	return &out, nil
}
