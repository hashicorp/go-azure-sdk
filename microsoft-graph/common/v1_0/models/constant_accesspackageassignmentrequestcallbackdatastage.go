package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageAssignmentRequestCallbackDataStage string

const (
	AccessPackageAssignmentRequestCallbackDataStageassignmentFourteenDaysBeforeExpiration AccessPackageAssignmentRequestCallbackDataStage = "AssignmentFourteenDaysBeforeExpiration"
	AccessPackageAssignmentRequestCallbackDataStageassignmentOneDayBeforeExpiration       AccessPackageAssignmentRequestCallbackDataStage = "AssignmentOneDayBeforeExpiration"
	AccessPackageAssignmentRequestCallbackDataStageassignmentRequestApproved              AccessPackageAssignmentRequestCallbackDataStage = "AssignmentRequestApproved"
	AccessPackageAssignmentRequestCallbackDataStageassignmentRequestCreated               AccessPackageAssignmentRequestCallbackDataStage = "AssignmentRequestCreated"
	AccessPackageAssignmentRequestCallbackDataStageassignmentRequestGranted               AccessPackageAssignmentRequestCallbackDataStage = "AssignmentRequestGranted"
	AccessPackageAssignmentRequestCallbackDataStageassignmentRequestRemoved               AccessPackageAssignmentRequestCallbackDataStage = "AssignmentRequestRemoved"
)

func PossibleValuesForAccessPackageAssignmentRequestCallbackDataStage() []string {
	return []string{
		string(AccessPackageAssignmentRequestCallbackDataStageassignmentFourteenDaysBeforeExpiration),
		string(AccessPackageAssignmentRequestCallbackDataStageassignmentOneDayBeforeExpiration),
		string(AccessPackageAssignmentRequestCallbackDataStageassignmentRequestApproved),
		string(AccessPackageAssignmentRequestCallbackDataStageassignmentRequestCreated),
		string(AccessPackageAssignmentRequestCallbackDataStageassignmentRequestGranted),
		string(AccessPackageAssignmentRequestCallbackDataStageassignmentRequestRemoved),
	}
}

func parseAccessPackageAssignmentRequestCallbackDataStage(input string) (*AccessPackageAssignmentRequestCallbackDataStage, error) {
	vals := map[string]AccessPackageAssignmentRequestCallbackDataStage{
		"assignmentfourteendaysbeforeexpiration": AccessPackageAssignmentRequestCallbackDataStageassignmentFourteenDaysBeforeExpiration,
		"assignmentonedaybeforeexpiration":       AccessPackageAssignmentRequestCallbackDataStageassignmentOneDayBeforeExpiration,
		"assignmentrequestapproved":              AccessPackageAssignmentRequestCallbackDataStageassignmentRequestApproved,
		"assignmentrequestcreated":               AccessPackageAssignmentRequestCallbackDataStageassignmentRequestCreated,
		"assignmentrequestgranted":               AccessPackageAssignmentRequestCallbackDataStageassignmentRequestGranted,
		"assignmentrequestremoved":               AccessPackageAssignmentRequestCallbackDataStageassignmentRequestRemoved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AccessPackageAssignmentRequestCallbackDataStage(input)
	return &out, nil
}
