package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomExtensionStageSettingStage string

const (
	CustomExtensionStageSettingStageassignmentFourteenDaysBeforeExpiration CustomExtensionStageSettingStage = "AssignmentFourteenDaysBeforeExpiration"
	CustomExtensionStageSettingStageassignmentOneDayBeforeExpiration       CustomExtensionStageSettingStage = "AssignmentOneDayBeforeExpiration"
	CustomExtensionStageSettingStageassignmentRequestApproved              CustomExtensionStageSettingStage = "AssignmentRequestApproved"
	CustomExtensionStageSettingStageassignmentRequestCreated               CustomExtensionStageSettingStage = "AssignmentRequestCreated"
	CustomExtensionStageSettingStageassignmentRequestGranted               CustomExtensionStageSettingStage = "AssignmentRequestGranted"
	CustomExtensionStageSettingStageassignmentRequestRemoved               CustomExtensionStageSettingStage = "AssignmentRequestRemoved"
)

func PossibleValuesForCustomExtensionStageSettingStage() []string {
	return []string{
		string(CustomExtensionStageSettingStageassignmentFourteenDaysBeforeExpiration),
		string(CustomExtensionStageSettingStageassignmentOneDayBeforeExpiration),
		string(CustomExtensionStageSettingStageassignmentRequestApproved),
		string(CustomExtensionStageSettingStageassignmentRequestCreated),
		string(CustomExtensionStageSettingStageassignmentRequestGranted),
		string(CustomExtensionStageSettingStageassignmentRequestRemoved),
	}
}

func parseCustomExtensionStageSettingStage(input string) (*CustomExtensionStageSettingStage, error) {
	vals := map[string]CustomExtensionStageSettingStage{
		"assignmentfourteendaysbeforeexpiration": CustomExtensionStageSettingStageassignmentFourteenDaysBeforeExpiration,
		"assignmentonedaybeforeexpiration":       CustomExtensionStageSettingStageassignmentOneDayBeforeExpiration,
		"assignmentrequestapproved":              CustomExtensionStageSettingStageassignmentRequestApproved,
		"assignmentrequestcreated":               CustomExtensionStageSettingStageassignmentRequestCreated,
		"assignmentrequestgranted":               CustomExtensionStageSettingStageassignmentRequestGranted,
		"assignmentrequestremoved":               CustomExtensionStageSettingStageassignmentRequestRemoved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CustomExtensionStageSettingStage(input)
	return &out, nil
}
