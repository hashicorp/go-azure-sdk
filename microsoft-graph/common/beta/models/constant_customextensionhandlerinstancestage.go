package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomExtensionHandlerInstanceStage string

const (
	CustomExtensionHandlerInstanceStageassignmentFourteenDaysBeforeExpiration CustomExtensionHandlerInstanceStage = "AssignmentFourteenDaysBeforeExpiration"
	CustomExtensionHandlerInstanceStageassignmentOneDayBeforeExpiration       CustomExtensionHandlerInstanceStage = "AssignmentOneDayBeforeExpiration"
	CustomExtensionHandlerInstanceStageassignmentRequestApproved              CustomExtensionHandlerInstanceStage = "AssignmentRequestApproved"
	CustomExtensionHandlerInstanceStageassignmentRequestCreated               CustomExtensionHandlerInstanceStage = "AssignmentRequestCreated"
	CustomExtensionHandlerInstanceStageassignmentRequestGranted               CustomExtensionHandlerInstanceStage = "AssignmentRequestGranted"
	CustomExtensionHandlerInstanceStageassignmentRequestRemoved               CustomExtensionHandlerInstanceStage = "AssignmentRequestRemoved"
)

func PossibleValuesForCustomExtensionHandlerInstanceStage() []string {
	return []string{
		string(CustomExtensionHandlerInstanceStageassignmentFourteenDaysBeforeExpiration),
		string(CustomExtensionHandlerInstanceStageassignmentOneDayBeforeExpiration),
		string(CustomExtensionHandlerInstanceStageassignmentRequestApproved),
		string(CustomExtensionHandlerInstanceStageassignmentRequestCreated),
		string(CustomExtensionHandlerInstanceStageassignmentRequestGranted),
		string(CustomExtensionHandlerInstanceStageassignmentRequestRemoved),
	}
}

func parseCustomExtensionHandlerInstanceStage(input string) (*CustomExtensionHandlerInstanceStage, error) {
	vals := map[string]CustomExtensionHandlerInstanceStage{
		"assignmentfourteendaysbeforeexpiration": CustomExtensionHandlerInstanceStageassignmentFourteenDaysBeforeExpiration,
		"assignmentonedaybeforeexpiration":       CustomExtensionHandlerInstanceStageassignmentOneDayBeforeExpiration,
		"assignmentrequestapproved":              CustomExtensionHandlerInstanceStageassignmentRequestApproved,
		"assignmentrequestcreated":               CustomExtensionHandlerInstanceStageassignmentRequestCreated,
		"assignmentrequestgranted":               CustomExtensionHandlerInstanceStageassignmentRequestGranted,
		"assignmentrequestremoved":               CustomExtensionHandlerInstanceStageassignmentRequestRemoved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomExtensionHandlerInstanceStage(input)
	return &out, nil
}
